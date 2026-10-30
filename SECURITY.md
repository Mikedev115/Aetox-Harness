# Security policy

English · [ภาษาไทย](SECURITY.th.md)

Aetox runs on your computer and acts with your permissions: it reads your files, runs commands, drives a
browser and writes documents. That is the point of it, and it is also where a bug can cost you something.
This page says how to tell us about one, and what we will do about it.

## Supported versions

The latest release is the one we fix. Aetox updates itself from inside the app (the Microsoft Store copy
updates through Windows), so a fix reaches you in the next release rather than in a patch on an old one.
Older builds are not maintained.

## How to report a vulnerability

Use **GitHub's private vulnerability reporting** on this repository: open the **Security** tab and press
**Report a vulnerability**. The report stays private between you and the maintainer.

Please do not open a public issue, a discussion, or a pull request that describes how to break something.
If you cannot use private reporting for any reason, open an issue that says only that you need a private
channel — no details, no steps, no proof of concept — and we will open one with you.

## What we will do

- Acknowledge your report within **7 days**. Aetox is built by one person, so this is a best effort, not a
  service-level promise.
- Tell you whether we could reproduce it, and what we intend to change.
- Credit you in the release notes when the fix ships, unless you ask us not to.
- Agree with you on when the report becomes public. We will not publish details before a fix is out.

## In scope

- **Permissions and approvals** — anything that does work without the approval your setting asks for, or
  that widens what a setting allows.
- **Path checks** — a file tool reading or writing outside the project or workspace it was given.
- **Credential protection** — reading a key store, SSH key, cloud credential or browser profile that Aetox
  is meant to refuse.
- **MCP and skills** — installing or running one that reaches further than its declared tools, or a skill
  body that can act on the machine when the user only asked it to describe something.
- **The mobile link** — pairing, the Tailscale HTTPS check, or anything that lets a device other than the
  paired one reach the PC.
- **The update path** — making Aetox install or run something other than its own signed release.
- **The engine over SSH** — a connection that is not the one the user configured.

## Out of scope

- The security of a model provider you connect. Report that to them; tell us if Aetox is what exposes it.
- Third-party MCP servers and skills you install. Aetox gates what they reach, not what they do inside
  that gate.
- A prompt injection that only does what you already allowed. Report it if it crosses a permission
  boundary — that part is ours.
- An attacker who already has administrator rights on your Windows account, or your unlocked machine.
- Dependencies. Report those upstream; a note to us is welcome so we can bump the version.

## Choices that are not bugs

These are deliberate, documented, and reported to us regularly:

- **The agent runs shell commands.** That is the product. Your approval setting is the boundary, so the bug
  is a command that runs *without* the approval you set, not that commands can run at all.
- **The GitHub installer is not code-signed yet.** Windows SmartScreen warns once. The Microsoft Store and
  winget builds are signed by Microsoft and do not warn. This is on the download page and in the README.
- **The engine can run on another machine** over SSH or WSL when you set that up.
- **Requests go to the provider you chose**, straight from your machine. There is no server of ours in
  between, and nothing is logged to us.

## Safe harbour

Research on your own machine, with your own data, is welcome. Do not test against other people's
installations, do not run denial-of-service tests, and do not use social engineering against the
maintainer. Work done in good faith under this page will not be met with legal action from us.

## Thank you

A harness that runs your real work has to be trustworthy before it is useful. Reports are how it stays
that way.
