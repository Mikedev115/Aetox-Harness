package devin

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/Mikedev115/Aetox/internal/cliagent"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

func fixture(lines string) (*client, *bytes.Buffer) {
	out := new(bytes.Buffer)
	return &client{ctx: context.Background(), in: out, scan: bufio.NewScanner(strings.NewReader(lines))}, out
}

func TestPermissionRequiresExplicitApproval(t *testing.T) {
	const request = `{"jsonrpc":"2.0","id":"permission-1","method":"session/request_permission","params":{"toolCall":{"title":"delete files","rawInput":{"path":"/project"}},"options":[{"optionId":"yes","kind":"allow_once"},{"optionId":"no","kind":"reject_once"}]}}`
	for _, choice := range []string{"missing", "allow", "deny", "cancel"} {
		t.Run(choice, func(t *testing.T) {
			c, out := fixture(request + "\n" + `{"id":1,"result":{}}` + "\n")
			if choice != "missing" {
				c.permission = func(title, detail string) (bool, error) {
					if title != "delete files" || !strings.Contains(detail, "/project") {
						t.Fatal("lost permission details")
					}
					if choice == "cancel" {
						return false, context.Canceled
					}
					return choice == "allow", nil
				}
			}
			if err := c.call("session/prompt", nil, nil); err != nil {
				t.Fatal(err)
			}
			lines := strings.Split(strings.TrimSpace(out.String()), "\n")
			var reply struct {
				ID     string `json:"id"`
				Result struct {
					Outcome struct{ Outcome, OptionID string } `json:"outcome"`
				} `json:"result"`
			}
			if err := json.Unmarshal([]byte(lines[1]), &reply); err != nil {
				t.Fatal(err)
			}
			if reply.ID != "permission-1" {
				t.Fatal("response lost request id")
			}
			want := "cancelled"
			if choice == "allow" || choice == "deny" {
				want = "selected"
			}
			if reply.Result.Outcome.Outcome != want {
				t.Fatalf("unexpected reply %s", lines[1])
			}
			if choice == "allow" && !strings.Contains(lines[1], `"optionId":"yes"`) {
				t.Fatal(lines[1])
			}
			if choice == "deny" && !strings.Contains(lines[1], `"optionId":"no"`) {
				t.Fatal(lines[1])
			}
		})
	}
}

func TestTurnStreamsToolsAndUsesServerSessionID(t *testing.T) {
	c, out := fixture(`{"id":1,"result":{"sessionId":"real-session"}}
{"method":"session/update","params":{"sessionId":"real-session","update":{"sessionUpdate":"agent_message_chunk","content":{"type":"text","text":"Before "}}}}
{"method":"session/update","params":{"sessionId":"real-session","update":{"sessionUpdate":"tool_call","toolCallId":"t1","title":"read_file"}}}
{"method":"session/update","params":{"sessionId":"real-session","update":{"sessionUpdate":"tool_call_update","toolCallId":"t1","status":"completed"}}}
{"method":"session/update","params":{"sessionId":"real-session","update":{"sessionUpdate":"agent_message_chunk","content":{"type":"text","text":"After"}}}}
{"id":2,"result":{"stopReason":"end_turn"}}
`)
	var events []cliagent.Event
	res, err := runTurn(c, cliagent.Turn{Dir: t.TempDir(), Text: "hello", SystemNote: "Aetox context"}, func(e cliagent.Event) { events = append(events, e) })
	if err != nil {
		t.Fatal(err)
	}
	if res.Text != "Before After" || res.SessionID != "real-session" {
		t.Fatalf("%+v", res)
	}
	if len(events) != 6 || events[1].Kind != cliagent.EventAnswer || events[2].Kind != cliagent.EventToolCall || events[3].Tool.Ref != "t1" || !events[3].Tool.OK {
		t.Fatalf("%+v", events)
	}
	if !strings.Contains(out.String(), `"sessionId":"real-session"`) || !strings.Contains(out.String(), "Aetox context") {
		t.Fatal(out.String())
	}
}

func TestResumeUsesModelConfigAndDoesNotReplayHistory(t *testing.T) {
	c, out := fixture(`{"id":1,"result":{"configOptions":[{"id":"model-select","category":"model","options":[]}]}}
{"id":2,"result":{}}
{"id":3,"result":{"stopReason":"end_turn"}}
`)
	c.init.AgentCapabilities.LoadSession = true
	_, err := runTurn(c, cliagent.Turn{Dir: t.TempDir(), SessionID: "saved", Resume: true, Model: "chosen-model", Text: "next", SystemNote: "never replay"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"session/load", "session/set_config_option", `"value":"chosen-model"`, `"sessionId":"saved"`} {
		if !strings.Contains(out.String(), want) {
			t.Fatal(out.String())
		}
	}
	if strings.Contains(out.String(), "never replay") {
		t.Fatal("replayed context")
	}
}

func TestFailuresAreNotSuccessfulAnswers(t *testing.T) {
	for _, body := range []string{"", `{"id":1,"error":{"code":-32000,"message":"Login required"}}`, `{"id":9,"result":{}}`, "invalid"} {
		c, _ := fixture(body + "\n")
		if err := c.call("session/new", nil, nil); err == nil {
			t.Fatalf("accepted %q", body)
		}
	}
	c, _ := fixture("")
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	c.ctx = ctx
	if err := c.call("session/prompt", nil, nil); !errors.Is(err, context.Canceled) {
		t.Fatalf("%v", err)
	}
}

// Opt-in transport smoke: initialize the installed official CLI — project
// access, model inference and billable prompts stay out of it.
func TestInstalledCLIHandshake(t *testing.T) {
	if os.Getenv("AETOX_TEST_DEVIN") != "1" {
		t.Skip("set AETOX_TEST_DEVIN=1 for installed CLI handshake")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	c, err := open(ctx, os.Getenv("AETOX_TEST_DEVIN_MODEL"), false)
	if err != nil {
		t.Fatal(err)
	}
	defer c.close()
	if c.init.ProtocolVersion != 1 || len(c.init.AuthMethods) == 0 {
		t.Fatal("missing protocol/auth advertisement")
	}
	var s session
	if err := c.call("session/new", map[string]any{"cwd": t.TempDir(), "mcpServers": []any{}}, &s); err != nil {
		t.Fatal(err)
	}
	for _, opt := range s.ConfigOptions {
		t.Logf("config category=%q id=%q current=%q", opt.Category, opt.ID, opt.CurrentValue)
	}
	var models []string
	for _, opt := range s.ConfigOptions {
		if opt.Category == "model" {
			t.Logf("current model: %q", opt.CurrentValue)
			for _, m := range opt.Options {
				models = append(models, m.Value)
				for _, n := range m.Options {
					models = append(models, n.Value)
				}
			}
		}
	}
	t.Logf("account models: %v", models)
}

// Opt-in mode switch: session/set_config_option must accept a mode change on
// a real session. Free — no prompt is sent.
func TestInstalledCLIModeSwitch(t *testing.T) {
	if os.Getenv("AETOX_TEST_DEVIN") != "1" {
		t.Skip("set AETOX_TEST_DEVIN=1 for installed CLI mode switch")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	c, err := open(ctx, "", false)
	if err != nil {
		t.Fatal(err)
	}
	defer c.close()
	var s session
	if err := c.call("session/new", map[string]any{"cwd": t.TempDir(), "mcpServers": []any{}}, &s); err != nil {
		t.Fatal(err)
	}
	for _, mode := range []string{"bypass", "smart", "accept-edits", "plan", "ask"} {
		if err := setConfig(c, &s, "mode", mode); err != nil {
			t.Fatal(err)
		}
		assertConfig(t, &s, "mode", mode)
	}
	for _, opt := range s.ConfigOptions {
		if effortCategory(opt.Category) {
			t.Logf("live effort options: %+v", opt.Options)
			for _, choice := range opt.Options {
				if choice.Value == "" {
					continue
				}
				if err := syncEffort(c, &s, choice.Value); err != nil {
					t.Fatal(err)
				}
				assertConfig(t, &s, opt.ID, choice.Value)
				for _, current := range s.ConfigOptions {
					if current.Category == "model" {
						t.Logf("effort %s -> model %s", choice.Value, current.CurrentValue)
					}
				}
			}
		}
	}

}

// Opt-in end-to-end: real authenticate, real session, one billable prompt
// against the sign-in on this machine — run it deliberately, never in CI.
func TestInstalledCLIRoundTrip(t *testing.T) {
	if os.Getenv("AETOX_TEST_DEVIN") != "1" {
		t.Skip("set AETOX_TEST_DEVIN=1 for installed CLI round-trip")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()
	models, err := Models(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(models) < 2 {
		t.Fatalf("model list: %v", models)
	}
	var events []cliagent.Event
	res, err := Engine{}.Run(ctx, cliagent.Turn{
		Dir:  t.TempDir(),
		Text: "Reply with exactly the word PONG and nothing else.",
	}, func(e cliagent.Event) { events = append(events, e) })
	if err != nil {
		t.Fatal(err)
	}
	if res.SessionID == "" || !strings.Contains(res.Text, "PONG") {
		t.Fatalf("%+v", res)
	}
	// Plan stance: syncMode must land on a real session and the read-only
	// directive must hold — the turn answers without error and writes nothing.
	res, err = Engine{}.Run(ctx, cliagent.Turn{
		Dir:       t.TempDir(),
		Text:      "In one sentence, what would you inspect first in an empty project?",
		Stance:    "plan",
		SessionID: res.SessionID,
		Resume:    true,
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.SessionID == "" || res.Text == "" {
		t.Fatalf("%+v", res)
	}
}

func assertConfig(t *testing.T, s *session, id, value string) {
	t.Helper()
	for _, opt := range s.ConfigOptions {
		if opt.ID == id && opt.CurrentValue == value {
			return
		}
	}
	t.Fatalf("response did not confirm %s=%s", id, value)
}

func TestAuthenticateUsesACPMetadataWithoutBrowserFallback(t *testing.T) {
	root := t.TempDir()
	t.Setenv("APPDATA", root)
	t.Setenv("HOME", root)
	dir := filepath.Join(root, "devin")
	if runtime.GOOS != "windows" {
		dir = filepath.Join(root, ".config", "devin")
	}
	if err := os.MkdirAll(dir, 0700); err != nil {
		t.Fatal(err)
	}
	file := filepath.Join(dir, "credentials.toml")
	if err := os.WriteFile(file, []byte("windsurf_api_key = \"fixture-key\"\napi_server_url = \"https://fixture.invalid\"\n"), 0600); err != nil {
		t.Fatal(err)
	}
	for _, result := range []string{`{"id":1,"result":{}}`, `{"id":1,"error":{"code":-1,"message":"expired"}}`} {
		c, out := fixture(result + "\n")
		err := c.authenticate(false)
		if strings.Contains(result, "expired") != (err != nil) {
			t.Fatal("unexpected authentication result")
		}
		var request struct {
			Params map[string]json.RawMessage `json:"params"`
		}
		if json.Unmarshal(out.Bytes(), &request) != nil {
			t.Fatal("expected exactly one request")
		}
		if _, ok := request.Params["meta"]; ok {
			t.Fatal("nonstandard meta triggers browser login")
		}
		var meta map[string]string
		if json.Unmarshal(request.Params["_meta"], &meta) != nil || meta["api_key"] != "fixture-key" || meta["api_server_url"] != "https://fixture.invalid" {
			t.Fatal("missing credential metadata")
		}
	}
	if err := os.Remove(file); err != nil {
		t.Fatal(err)
	}
	c, out := fixture("")
	if c.authenticate(false) == nil || out.Len() != 0 {
		t.Fatal("missing credentials must not start browser authentication")
	}
}
