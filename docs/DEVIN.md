# Devin CLI provider

Aetox can run the official Devin CLI as an external agent using ACP over stdio.
Install the CLI using https://docs.devin.ai/cli/index, then add Devin in Settings
and choose Sign in. The CLI opens its own interactive terminal login. Models
and billing follow the signed-in Devin account and organization. The installed
CLI must advertise ACP model and session-mode configuration options.

The account credential remains in the official CLI credential file. Each ACP
process receives it through `authenticate` using the protocol's `_meta` field.
Model and effort discovery do not run inference or start browser sign-in.

The composer offers the account's live models, friendly names with billing IDs,
and one effort selector. When the CLI advertises `thought_level`, that option
controls effort; otherwise advertised SWE-2 tier IDs are separate choices.
The visible name includes an explicit effort override. Selecting another model
clears the previous effort override. Selecting Default on a resumed session
resolves the CLI's current default before applying it.

Approval and stance changes apply on the next message:

| Aetox | Devin mode |
| --- | --- |
| Ask | Smart, falling back to Ask |
| Unsafe Only | Accept Edits |
| Full Access | Bypass |
| Plan stance | Plan |
| Consult stance | Ask, with tool permission requests denied |

A missing/refused mode change stops the turn. Approval requests are displayed
in Aetox. Session IDs are persisted per Aetox chat, provider and project so
reopening a chat resumes the external agent's own history.

Devin runs its own tools and workflow. Aetox-native agent teams, automatic
memory injection and native tools are not available. Attachment paths are
forwarded as workspace files with instructions to use Devin's own tools;
no native Aetox image/PDF tool is injected. Quota is not reported by ACP.
The provider is intended for the local engine, using a local CLI installation.

## Verification

Normal Go/frontend tests use scripted protocol transcripts and need no account.
Installed CLI tests are opt-in and use the current account. The round-trip test
sends billable prompts in temporary folders:

```powershell
$env:AETOX_TEST_DEVIN = '1'
go test -v ./internal/devin -run Installed
```

Validated on Windows with Devin CLI 3000.11.228. Non-Windows terminal login
launchers are supported but were not manually tested in this contribution.
