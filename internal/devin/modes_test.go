package devin

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Mikedev115/Aetox/internal/cliagent"
)

func TestSessionModeFollowsAetoxDials(t *testing.T) {
	cases := []struct{ stance, approval, want string }{
		{"", "ask", "smart"},
		{"", "unsafe-only", "accept-edits"},
		{"", "full-access", "bypass"},
		{"plan", "full-access", "plan"},
		{"consult", "full-access", "ask"},
		{"", "", "smart"},
	}
	for _, c := range cases {
		if got := sessionMode(c.stance, c.approval); got != c.want {
			t.Errorf("sessionMode(%q,%q) = %q, want %q", c.stance, c.approval, got, c.want)
		}
	}
}

const permissionThenEnd = `{"id":1,"result":{"sessionId":"s","modes":{"currentModeId":"smart","availableModes":[{"id":"smart"},{"id":"ask"},{"id":"plan"},{"id":"accept-edits"},{"id":"bypass"}]}}}
{"id":2,"result":{}}
{"jsonrpc":"2.0","id":"p1","method":"session/request_permission","params":{"toolCall":{"title":"edit file"},"options":[{"optionId":"yes","kind":"allow_once"},{"optionId":"no","kind":"reject_once"}]}}
{"id":3,"result":{"stopReason":"end_turn"}}
`

func TestPlanStanceDeniesWritesAndSaysSo(t *testing.T) {
	c, out := fixture(permissionThenEnd)
	asked := false
	_, err := runTurn(c, cliagent.Turn{Dir: t.TempDir(), Text: "fix it", Stance: "plan", Approval: "full-access",
		Permission: func(string, string) (bool, error) { asked = true; return true, nil }}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if asked {
		t.Fatal("plan mode asked the user instead of refusing")
	}
	if !strings.Contains(out.String(), `"optionId":"no"`) || !strings.Contains(out.String(), "plan mode") {
		t.Fatal(out.String())
	}
}

func TestFullAccessApprovesWithoutAsking(t *testing.T) {
	c, out := fixture(permissionThenEnd)
	_, err := runTurn(c, cliagent.Turn{Dir: t.TempDir(), Text: "go", Approval: "full-access",
		Permission: func(string, string) (bool, error) { t.Fatal("asked in full access"); return false, nil }}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), `"optionId":"yes"`) {
		t.Fatal(out.String())
	}
}

func TestStanceSwitchesDevinSessionMode(t *testing.T) {
	c, out := fixture(`{"id":1,"result":{"sessionId":"s","modes":{"currentModeId":"code","availableModes":[{"id":"code"},{"id":"plan"},{"id":"ask"}]}}}
{"id":2,"result":{}}
{"id":3,"result":{"stopReason":"end_turn"}}
`)
	if _, err := runTurn(c, cliagent.Turn{Dir: t.TempDir(), Text: "think", Stance: "plan"}, nil); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), `"method":"session/set_mode"`) || !strings.Contains(out.String(), `"modeId":"plan"`) {
		t.Fatal(out.String())
	}
}

func TestFullAccessSetsBypassSessionMode(t *testing.T) {
	c, out := fixture(`{"id":1,"result":{"sessionId":"s","modes":{"currentModeId":"accept-edits","availableModes":[{"id":"accept-edits"},{"id":"bypass","name":"Bypass Permissions"}]}}}
{"id":2,"result":{}}
{"id":3,"result":{"stopReason":"end_turn"}}
`)
	if _, err := runTurn(c, cliagent.Turn{Dir: t.TempDir(), Text: "go", Approval: "full-access"}, nil); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), `"modeId":"bypass"`) {
		t.Fatal(out.String())
	}
}

func TestActStanceLeavesPlanModeOnResume(t *testing.T) {
	c, out := fixture(`{"id":1,"result":{"configOptions":[{"id":"mode","category":"mode","currentValue":"plan","options":[{"value":"plan"},{"value":"smart"}]}]}}
{"id":2,"result":{}}
{"id":3,"result":{"stopReason":"end_turn"}}
`)
	c.init.AgentCapabilities.LoadSession = true
	if _, err := runTurn(c, cliagent.Turn{Dir: t.TempDir(), SessionID: "saved", Resume: true, Text: "do it"}, nil); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), `"configId":"mode"`) || !strings.Contains(out.String(), `"value":"smart"`) {
		t.Fatal(out.String())
	}
}

func TestApprovalReassertedOnNewAndLoad(t *testing.T) {
	for _, resume := range []bool{false, true} {
		for _, approval := range []string{"ask", "unsafe-only", "full-access"} {
			want := sessionMode("act", approval)
			c, out := fixture(fmt.Sprintf(`{"id":1,"result":{"sessionId":"s","configOptions":[{"id":"mode","category":"mode","currentValue":%q,"options":[{"value":"smart"},{"value":"accept-edits"},{"value":"bypass"}]}]}}
{"id":2,"result":{"configOptions":[{"id":"mode","category":"mode","currentValue":%q}]}}
{"id":3,"result":{"stopReason":"end_turn"}}
`, want, want))
			c.init.AgentCapabilities.LoadSession = true
			_, err := runTurn(c, cliagent.Turn{Dir: t.TempDir(), Resume: resume, SessionID: "s", Approval: approval}, nil)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(out.String(), `"value":"`+want+`"`) {
				t.Fatal(out.String())
			}
		}
	}
}
func TestAskStillGatesPermission(t *testing.T) {
	c, out := fixture(permissionThenEnd)
	asked := false
	_, err := runTurn(c, cliagent.Turn{Dir: t.TempDir(), Approval: "ask", Permission: func(string, string) (bool, error) { asked = true; return false, nil }}, nil)
	if err != nil || !asked || !strings.Contains(out.String(), `"optionId":"no"`) {
		t.Fatalf("asked=%v err=%v output=%s", asked, err, out)
	}
}
func TestRefusedModeStopsBeforePrompt(t *testing.T) {
	c, out := fixture(`{"id":1,"result":{"sessionId":"s","configOptions":[{"id":"mode","category":"mode","currentValue":"bypass","options":[{"value":"smart"}]}]}}
{"id":2,"error":{"code":-1,"message":"refused"}}
`)
	_, err := runTurn(c, cliagent.Turn{Dir: t.TempDir(), Approval: "ask"}, nil)
	if err == nil || strings.Contains(out.String(), "session/prompt") {
		t.Fatal("continued after refused mode", out)
	}
}
func TestEffortUsesAdvertisedConfigID(t *testing.T) {
	c, out := fixture(`{"id":1,"result":{"sessionId":"s","configOptions":[{"id":"depth","category":"thought_level","options":[{"value":"extra-high"}]}],"modes":{"currentModeId":"smart","availableModes":[{"id":"smart"},{"id":"ask"},{"id":"plan"},{"id":"accept-edits"},{"id":"bypass"}]}}}
{"id":2,"result":{}}
{"id":3,"result":{"configOptions":[{"id":"depth","category":"thought_level","currentValue":"extra-high"}]}}
{"id":4,"result":{"stopReason":"end_turn"}}
`)
	_, err := runTurn(c, cliagent.Turn{Dir: t.TempDir(), Effort: "extra-high"}, nil)
	if err != nil || !strings.Contains(out.String(), `"configId":"depth"`) {
		t.Fatal(err, out)
	}
}
func TestConfigMismatchRejected(t *testing.T) {
	c, _ := fixture(`{"id":1,"result":{"configOptions":[{"id":"mode","currentValue":"bypass"}],"modes":{"currentModeId":"smart","availableModes":[{"id":"smart"},{"id":"ask"},{"id":"plan"},{"id":"accept-edits"},{"id":"bypass"}]}}}
{"id":2,"result":{}}
`)
	if err := setConfig(c, &session{ID: "s"}, "mode", "smart"); err == nil {
		t.Fatal("accepted wrong mode")
	}
}

func TestMissingModeAdvertisementStopsBeforePrompt(t *testing.T) {
	for _, resume := range []bool{false, true} {
		c, out := fixture(`{"id":1,"result":{"sessionId":"s"}}
`)
		c.init.AgentCapabilities.LoadSession = true
		_, err := runTurn(c, cliagent.Turn{Dir: t.TempDir(), SessionID: "s", Resume: resume, Approval: "ask"}, nil)
		if err == nil || strings.Contains(out.String(), "session/prompt") {
			t.Fatal("prompted without establishing policy")
		}
	}
}
func TestUnsupportedEffortIsNotSilentlyIgnored(t *testing.T) {
	c, _ := fixture("")
	if syncEffort(c, &session{}, "xhigh") == nil {
		t.Fatal("ignored unsupported effort")
	}
}
