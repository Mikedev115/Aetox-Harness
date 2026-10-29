// Package devin connects to the official Devin CLI using ACP over stdio.
// The credential stays the CLI's own: Aetox reads the file `devin auth login`
// writes only to hand the key back to the CLI's ACP server, and stores
// nothing itself.
package devin

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/Mikedev115/Aetox/internal/cliagent"
	"github.com/Mikedev115/Aetox/internal/version"
)

const Note = "Devin runs its own agent and tools in your project. Aetox displays the conversation and approval requests; Aetox-native tools, agent teams and automatic memory injection are not available to this agent. Models and billing follow your Devin plan and organization permissions. Remaining quota is not reported here."

type Engine struct{}

func init()                     { cliagent.Register(Engine{}) }
func (Engine) Provider() string { return "devin" }
func (Engine) Label() string    { return "Devin" }

// Path also checks the official Windows installation location, so installing
// the CLI does not require restarting Aetox to refresh PATH.
func Path() (string, error) {
	if p, err := exec.LookPath("devin"); err == nil {
		return p, nil
	}
	if runtime.GOOS == "windows" {
		p := filepath.Join(os.Getenv("LOCALAPPDATA"), "devin", "cli", "bin", "devin.exe")
		if s, err := os.Stat(p); err == nil && !s.IsDir() {
			return p, nil
		}
	}
	return "", errors.New("Install the official Devin CLI from https://docs.devin.ai/cli/index, then connect your account")
}

// credentials reads the sign-in the official CLI already holds — the file
// `devin auth login` writes — and returns the two fields its ACP server
// accepts as an `authenticate` meta credential. Read fresh for every spawned
// process; nothing is stored here. Empty when the user has never signed in.
func credentials() (apiKey, apiServerURL string) {
	var p string
	if runtime.GOOS == "windows" {
		p = filepath.Join(os.Getenv("APPDATA"), "devin", "credentials.toml")
	} else if home, err := os.UserHomeDir(); err == nil {
		p = filepath.Join(home, ".config", "devin", "credentials.toml")
	}
	data, err := os.ReadFile(p)
	if err != nil {
		return "", ""
	}
	for _, line := range strings.Split(string(data), "\n") {
		k, v, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		v = strings.Trim(strings.TrimSpace(v), `"`)
		switch strings.TrimSpace(k) {
		case "windsurf_api_key":
			apiKey = v
		case "api_server_url":
			apiServerURL = v
		}
	}
	return apiKey, apiServerURL
}

func (Engine) Probe(ctx context.Context) cliagent.Status {
	p, err := Path()
	if err != nil {
		return cliagent.Status{Detail: err.Error()}
	}
	// `auth status` cannot see the credential format every account lands on,
	// so the file itself is the check: no key, no turn could authenticate.
	key, _ := credentials()
	st := cliagent.Status{Path: p, Ready: key != ""}
	if !st.Ready {
		st.Detail = "Connect your Devin account using Sign in. If access is denied, check your Devin plan and organization permissions."
	}
	return st
}

type packet struct {
	ID     json.RawMessage `json:"id"`
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
	Result json.RawMessage `json:"result"`
	Error  *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}
type client struct {
	ctx        context.Context
	in         io.Writer
	scan       *bufio.Scanner
	next       int
	update     func(json.RawMessage)
	permission func(string, string) (bool, error)
	init       struct {
		ProtocolVersion   int `json:"protocolVersion"`
		AgentCapabilities struct {
			LoadSession bool `json:"loadSession"`
		} `json:"agentCapabilities"`
		AuthMethods []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"authMethods"`
	}
	close func()
}

func open(ctx context.Context, model string, interactive bool) (*client, error) {
	p, err := Path()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(ctx)
	args := []string{"acp"}
	if model != "" && model != "default" {
		args = append(args, "--model", model)
	}
	cmd := cliagent.Command(ctx, p, args...)
	in, err := cmd.StdinPipe()
	if err != nil {
		cancel()
		return nil, err
	}
	out, err := cmd.StdoutPipe()
	if err != nil {
		cancel()
		in.Close()
		return nil, err
	}
	// stderr can contain account details; do not forward it into the chat or logs.
	cmd.Stderr = io.Discard
	if err = cmd.Start(); err != nil {
		cancel()
		in.Close()
		out.Close()
		return nil, err
	}
	c := &client{ctx: ctx, in: in, scan: bufio.NewScanner(out)}
	c.scan.Buffer(make([]byte, 4096), 16*1024*1024)
	c.close = func() { cancel(); in.Close(); _ = cmd.Wait() }
	err = c.call("initialize", map[string]any{"protocolVersion": 1, "clientCapabilities": map[string]any{}, "clientInfo": map[string]string{"name": "aetox", "version": version.Current}}, &c.init)
	if err != nil {
		c.close()
		return nil, err
	}
	if c.init.ProtocolVersion != 1 {
		c.close()
		return nil, errors.New("Devin requires an unsupported ACP protocol version")
	}
	if err = c.authenticate(interactive); err != nil {
		c.close()
		return nil, err
	}
	return c, nil
}

// ACP names its metadata field _meta on the wire (the Rust field is meta).
// Sending plain meta silently discards the key and starts the browser PKCE flow.
// Background clients never fall back to browser login; Login owns that action.
func (c *client) authenticate(_ bool) error {
	key, server := credentials()
	if key == "" {
		return errors.New("connect your Devin account using Sign in under Settings")
	}
	meta := map[string]any{"api_key": key}
	if server != "" {
		meta["api_server_url"] = server
	}
	if err := c.call("authenticate", map[string]any{"methodId": "devin-browser", "_meta": meta}, nil); err != nil {
		return fmt.Errorf("Devin sign-in needs renewing — use Sign in under Settings: %w", err)
	}
	return nil
}
func (c *client) send(v any) error { return json.NewEncoder(c.in).Encode(v) }
func (c *client) call(method string, params any, result any) error {
	c.next++
	id := c.next
	if err := c.send(map[string]any{"jsonrpc": "2.0", "id": id, "method": method, "params": params}); err != nil {
		return err
	}
	for c.scan.Scan() {
		var p packet
		if err := json.Unmarshal(c.scan.Bytes(), &p); err != nil {
			return fmt.Errorf("invalid Devin ACP message: %w", err)
		}
		if p.Method != "" {
			if len(p.ID) > 0 {
				var err error
				if p.Method == "session/request_permission" {
					err = c.answerPermission(p)
				} else {
					err = c.send(map[string]any{"jsonrpc": "2.0", "id": p.ID, "error": map[string]any{"code": -32601, "message": "Client capability not supported"}})
				}
				if err != nil {
					return err
				}
			} else if p.Method == "session/update" && c.update != nil {
				c.update(p.Params)
			}
			continue
		}
		if string(p.ID) != fmt.Sprint(id) {
			return errors.New("unexpected Devin ACP response id")
		}
		if p.Error != nil {
			return fmt.Errorf("Devin: %s (ACP %d)", p.Error.Message, p.Error.Code)
		}
		if result != nil {
			return json.Unmarshal(p.Result, result)
		}
		return nil
	}
	if c.ctx.Err() != nil {
		return c.ctx.Err()
	}
	if err := c.scan.Err(); err != nil {
		return err
	}
	return errors.New("Devin disconnected before completing the request")
}
func (c *client) answerPermission(p packet) error {
	var req struct {
		ToolCall struct {
			Title    string          `json:"title"`
			RawInput json.RawMessage `json:"rawInput"`
		} `json:"toolCall"`
		Options []struct {
			ID   string `json:"optionId"`
			Kind string `json:"kind"`
		} `json:"options"`
	}
	if err := json.Unmarshal(p.Params, &req); err != nil {
		return err
	}
	outcome := map[string]string{"outcome": "cancelled"}
	allowID, rejectID := "", ""
	for _, o := range req.Options {
		if o.Kind == "allow_once" {
			allowID = o.ID
		}
		if o.Kind == "reject_once" {
			rejectID = o.ID
		}
	}
	if allowID != "" && c.permission != nil {
		ok, err := c.permission(req.ToolCall.Title, string(req.ToolCall.RawInput))
		if err == nil {
			chosen := rejectID
			if ok {
				chosen = allowID
			}
			if chosen != "" {
				outcome = map[string]string{"outcome": "selected", "optionId": chosen}
			}
		}
	}
	return c.send(map[string]any{"jsonrpc": "2.0", "id": p.ID, "result": map[string]any{"outcome": outcome}})
}

// Login connects the account. Only the CLI's own `auth login` writes the
// credential every later process reads — its picker is a real-terminal UI,
// so a console is opened for it and this waits for the credential to land.
func Login(ctx context.Context) error {
	p, err := Path()
	if err != nil {
		return err
	}
	cmd := loginConsole(ctx, p)
	if cmd == nil {
		return errors.New("no terminal available — run `devin auth login` in one to connect Devin")
	}
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("could not open a terminal — run `devin auth login` in one to connect Devin: %w", err)
	}
	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()
	select {
	case <-ctx.Done():
		_ = cmd.Process.Kill()
		return ctx.Err()
	case err := <-done:
		if err != nil {
			return fmt.Errorf("Devin sign-in did not finish: %w", err)
		}
	}
	if runtime.GOOS != "windows" {
		// Terminal launchers may exit while their interactive login is still running.
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			// A successful login need not rotate the API key. Verify readiness
			// rather than treating a launcher exit or changed key as success.
			if key, _ := credentials(); key != "" {
				probeCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
				client, err := open(probeCtx, "", false)
				if err == nil {
					client.close()
				}
				cancel()
				if err == nil {
					break
				}
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-ticker.C:
			}
		}
	}
	if key, _ := credentials(); key == "" {
		return errors.New("Devin sign-in did not finish — run `devin auth login` in a terminal")
	}
	return nil
}

// loginConsole builds the one sign-in command: `devin auth login` in a console
// the user can see and answer. devin.exe is a console program, so on Windows
// starting it plainly from the gui build gives it a window of its own;
// elsewhere a terminal emulator is asked to host it instead.
func loginConsole(ctx context.Context, devinPath string) *exec.Cmd {
	switch runtime.GOOS {
	case "windows":
		return exec.CommandContext(ctx, devinPath, "auth", "login")
	case "darwin":
		return exec.CommandContext(ctx, "osascript", "-e",
			`tell application "Terminal" to do script `+strconv.Quote("'"+strings.ReplaceAll(devinPath, "'", "'\\''")+"' auth login"))
	default:
		for _, t := range []string{"x-terminal-emulator", "gnome-terminal", "konsole", "xterm"} {
			if bin, err := exec.LookPath(t); err == nil {
				return exec.CommandContext(ctx, bin, "-e", devinPath, "auth", "login")
			}
		}
		return nil
	}
}
func Logout(ctx context.Context) error {
	p, err := Path()
	if err != nil {
		return err
	}
	return cliagent.Command(ctx, p, "auth", "logout").Run()
}

type session struct {
	ConfigOptions []struct {
		ID           string `json:"id"`
		Category     string `json:"category"`
		CurrentValue string `json:"currentValue"`
		Options      []struct {
			Value   string `json:"value"`
			Options []struct {
				Value string `json:"value"`
			} `json:"options"`
		} `json:"options"`
	} `json:"configOptions"`
	ID     string `json:"sessionId"`
	Models struct {
		Available []struct {
			ID string `json:"modelId"`
		} `json:"availableModels"`
	} `json:"models"`
	// Modes is ACP's session-mode list (session/set_mode). Devin may instead
	// expose its modes as a config option with category "mode"; both are read.
	Modes struct {
		Current   string `json:"currentModeId"`
		Available []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"availableModes"`
	} `json:"modes"`
}

// Models asks the official CLI for account-specific availability, without inference.
// A temporary ACP session has no project context or prompt.
func discover(ctx context.Context, model string) (*session, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	c, err := open(ctx, model, false)
	if err != nil {
		return nil, err
	}
	defer c.close()
	dir, err := os.MkdirTemp("", "aetox-devin-models-")
	if err != nil {
		return nil, err
	}
	defer os.Remove(dir)
	var s session
	if err = c.call("session/new", map[string]any{"cwd": dir, "mcpServers": []any{}}, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

func Models(ctx context.Context) ([]string, error) {
	s, err := discover(ctx, "")
	if err != nil {
		return nil, err
	}
	out := []string{"default"}
	seen := map[string]bool{"default": true}
	add := func(id string) {
		if id != "" && !seen[id] {
			out = append(out, id)
			seen[id] = true
		}
	}
	for _, opt := range s.ConfigOptions {
		if opt.Category == "model" {
			for _, m := range opt.Options {
				add(m.Value)
				for _, nested := range m.Options {
					add(nested.Value)
				}
			}
		}
	}
	for _, m := range s.Models.Available {
		if m.ID != "" {
			add(m.ID)
		}
	}
	return out, nil
}

func (Engine) Run(ctx context.Context, t cliagent.Turn, emit func(cliagent.Event)) (cliagent.Result, error) {
	if t.Resume && t.Model == "default" {
		defaults, err := discover(ctx, "")
		if err != nil {
			return cliagent.Result{}, err
		}
		for _, opt := range defaults.ConfigOptions {
			if opt.Category == "model" {
				t.Model = opt.CurrentValue
				break
			}
		}
		if t.Model == "default" || t.Model == "" {
			return cliagent.Result{}, errors.New("Devin did not report its default model")
		}
	}
	c, err := open(ctx, t.Model, false)
	if err != nil {
		return cliagent.Result{}, err
	}
	defer c.close()
	return runTurn(c, t, emit)
}

// readOnly reports whether Aetox's stance forbids changing anything this turn.
func readOnly(stance string) bool {
	return stance == "plan" || stance == "consult" || stance == "think"
}

// sessionMode keeps approval and stance independent. Unknown approval is conservative.
func sessionMode(stance, approval string) string {
	switch stance {
	case "plan":
		return "plan"
	case "consult", "think":
		return "ask"
	}
	switch approval {
	case "full-access":
		return "bypass"
	case "unsafe-only":
		return "accept-edits"
	}
	return "smart"
}

// stanceDirective is re-sent on every turn, because the stance can move
// between two messages of one Devin session.
func stanceDirective(stance string) string {
	switch stance {
	case "plan":
		return "[Aetox plan mode] Investigate only and reply with a concrete, step-by-step plan. Do not create or edit files, do not run commands that change anything, and do not commit; such actions will be denied."
	case "consult", "think":
		return "[Aetox consult mode] Answer as a discussion partner from what you already know. Do not call tools; any action will be denied."
	}
	return ""
}

// syncMode reasserts the selected policy on every new or loaded turn.
// A refused switch stops the turn: retaining a previous bypass is unsafe.
func syncMode(c *client, s *session, stance, approval string) error {
	want := sessionMode(stance, approval)
	for _, opt := range s.ConfigOptions {
		if opt.Category != "mode" {
			continue
		}
		var ids []string
		for _, o := range opt.Options {
			ids = append(ids, o.Value)
			for _, n := range o.Options {
				ids = append(ids, n.Value)
			}
		}
		id := chooseMode(ids, want)
		if id == "" {
			return fmt.Errorf("Devin does not advertise mode %q", want)
		}
		return setConfig(c, s, opt.ID, id)
	}
	if len(s.Modes.Available) == 0 {
		return errors.New("Devin did not advertise session modes; refusing to prompt without an approval policy")
	}
	var ids []string
	for _, m := range s.Modes.Available {
		ids = append(ids, m.ID)
	}
	id := chooseMode(ids, want)
	if id == "" {
		return fmt.Errorf("Devin does not advertise mode %q", want)
	}
	if err := c.call("session/set_mode", map[string]any{"sessionId": s.ID, "modeId": id}, nil); err != nil {
		return err
	}
	s.Modes.Current = id
	return nil
}

func chooseMode(ids []string, want string) string {
	candidates := []string{want}
	if want == "smart" {
		candidates = append(candidates, "ask")
	}
	for _, candidate := range candidates {
		for _, id := range ids {
			if id == candidate {
				return id
			}
		}
	}
	return ""
}

func setConfig(c *client, s *session, configID, value string) error {
	var updated session
	if err := c.call("session/set_config_option", map[string]any{"sessionId": s.ID, "configId": configID, "value": value}, &updated); err != nil {
		return err
	}
	for _, opt := range updated.ConfigOptions {
		if opt.ID == configID && opt.CurrentValue != value {
			return fmt.Errorf("Devin retained %s=%q after requesting %q", configID, opt.CurrentValue, value)
		}
	}
	if len(updated.ConfigOptions) > 0 {
		s.ConfigOptions = updated.ConfigOptions
	}
	return nil
}

func effortCategory(category string) bool {
	switch strings.ToLower(category) {
	case "effort", "thinking", "thought_level", "reasoning", "reasoning_effort":
		return true
	}
	return false
}

func syncEffort(c *client, s *session, effort string) error {
	if effort == "" {
		return nil
	}
	for _, opt := range s.ConfigOptions {
		if !effortCategory(opt.Category) {
			continue
		}
		for _, o := range opt.Options {
			if o.Value == effort {
				return setConfig(c, s, opt.ID, effort)
			}
			for _, n := range o.Options {
				if n.Value == effort {
					return setConfig(c, s, opt.ID, effort)
				}
			}
		}
		return fmt.Errorf("Devin does not advertise effort %q", effort)
	}
	return errors.New("Devin does not advertise an effort option for this model")
}

func runTurn(c *client, t cliagent.Turn, emit func(cliagent.Event)) (cliagent.Result, error) {
	started := time.Now()
	res := cliagent.Result{}
	switch {
	case readOnly(t.Stance):
		// Plan and consult change nothing: Devin's default mode already runs
		// reads on its own, so every request that reaches us is a write or a
		// command, and is refused without bothering the user.
		c.permission = func(string, string) (bool, error) { return false, nil }
	case t.Approval == "full-access":
		c.permission = func(string, string) (bool, error) { return true, nil }
	default:
		c.permission = t.Permission
	}
	dir, err := filepath.Abs(t.Dir)
	if err != nil {
		return res, err
	}
	var s session
	params := map[string]any{"cwd": dir, "mcpServers": []any{}}
	if t.Resume {
		if !c.init.AgentCapabilities.LoadSession {
			return res, errors.New("Devin cannot resume this conversation; start a new chat")
		}
		params["sessionId"] = t.SessionID
		err = c.call("session/load", params, &s)
		s.ID = t.SessionID
	} else {
		err = c.call("session/new", params, &s)
	}
	if err != nil {
		return res, err
	}
	if s.ID == "" {
		return res, errors.New("Devin returned no session id")
	}
	res.SessionID = s.ID
	// The process-level default applies to new sessions. Explicitly update a
	// loaded session so changing Aetox's model picker takes effect too.
	if t.Resume && t.Model != "" && t.Model != "default" {
		method := "session/set_model"
		choice := map[string]any{"sessionId": s.ID, "modelId": t.Model}
		for _, opt := range s.ConfigOptions {
			if opt.Category == "model" {
				method = "session/set_config_option"
				choice = map[string]any{"sessionId": s.ID, "configId": opt.ID, "value": t.Model}
				break
			}
		}
		if method == "session/set_config_option" {
			err = setConfig(c, &s, choice["configId"].(string), t.Model)
		} else {
			err = c.call(method, choice, nil)
		}
		if err != nil {
			return res, err
		}
	}
	if err = syncMode(c, &s, t.Stance, t.Approval); err != nil {
		return res, err
	}
	if err = syncEffort(c, &s, t.Effort); err != nil {
		return res, err
	}
	var answer, block strings.Builder
	send := func(e cliagent.Event) {
		if emit != nil {
			emit(e)
		}
	}
	flush := func() {
		if block.Len() > 0 {
			send(cliagent.Event{Kind: cliagent.EventAnswer, Text: block.String()})
			block.Reset()
		}
	}
	c.update = func(raw json.RawMessage) {
		var msg struct {
			SessionID string `json:"sessionId"`
			Update    struct {
				Kind    string `json:"sessionUpdate"`
				Content struct {
					Type string `json:"type"`
					Text string `json:"text"`
				} `json:"content"`
				ID     string `json:"toolCallId"`
				Title  string `json:"title"`
				Status string `json:"status"`
			} `json:"update"`
		}
		if json.Unmarshal(raw, &msg) != nil || msg.SessionID != s.ID {
			return
		}
		u := msg.Update
		switch u.Kind {
		case "agent_message_chunk":
			if u.Content.Type == "text" {
				answer.WriteString(u.Content.Text)
				block.WriteString(u.Content.Text)
				send(cliagent.Event{Kind: cliagent.EventText, Text: u.Content.Text})
			}
		case "agent_thought_chunk":
			if u.Content.Type == "text" {
				send(cliagent.Event{Kind: cliagent.EventThinking, Text: u.Content.Text})
			}
		case "tool_call":
			flush()
			send(cliagent.Event{Kind: cliagent.EventToolCall, Tool: cliagent.Tool{Ref: u.ID, Name: u.Title}})
		case "tool_call_update":
			if u.Status == "completed" || u.Status == "failed" {
				send(cliagent.Event{Kind: cliagent.EventToolResult, Tool: cliagent.Tool{Ref: u.ID, Name: u.Title, OK: u.Status == "completed"}})
			}
		}
	}
	// Composer files remain workspace files; name the external agent's own tools.
	prompt := strings.NewReplacer("[attachment: user-attached file — read it with pdf_read]", "[attachment: user-attached file — inspect with your own file tools]", "[attachment: user-attached image — read it with image_ocr]", "[attachment: user-attached image — inspect with your own image tools]").Replace(t.Text)
	if d := stanceDirective(t.Stance); d != "" {
		prompt = d + "\n\n" + prompt
	}
	if !t.Resume && t.SystemNote != "" {
		prompt = t.SystemNote + "\n\nUser request:\n" + prompt
	}
	var reply struct {
		StopReason string `json:"stopReason"`
	}
	err = c.call("session/prompt", map[string]any{"sessionId": s.ID, "prompt": []any{map[string]string{"type": "text", "text": prompt}}}, &reply)
	flush()
	res.Text = answer.String()
	res.Duration = time.Since(started)
	if err == nil && reply.StopReason == "cancelled" {
		err = context.Canceled
	}
	if err == nil && reply.StopReason != "end_turn" {
		err = fmt.Errorf("Devin stopped: %s", reply.StopReason)
	}
	return res, err
}

// EffortLevels reads only the options advertised for the selected model.
func EffortLevels(ctx context.Context, model string) ([]string, error) {
	s, err := discover(ctx, model)
	if err != nil {
		return nil, err
	}
	levels := []string{}
	for _, opt := range s.ConfigOptions {
		if !effortCategory(opt.Category) {
			continue
		}
		for _, o := range opt.Options {
			if o.Value != "" {
				levels = append(levels, o.Value)
			}
			for _, n := range o.Options {
				if n.Value != "" {
					levels = append(levels, n.Value)
				}
			}
		}
	}
	return levels, nil
}
