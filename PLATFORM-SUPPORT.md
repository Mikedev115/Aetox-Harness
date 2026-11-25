# Platform Support — สถานะพอร์ต Linux/macOS · อัปเดต 2026-07-27

> **จุดยืน (owner, 2026-07-27):** กลับข้าง — จาก "บันทึกไว้เฉยๆ" เป็น **"ทำจริง เดสก์ท็อปก่อน"** เป้าหมาย v0.7.0
> Decision section: [ARCHITECTURE.md §48](docs/internal/ARCHITECTURE.md) · จุดยืนเดิม "Windows first, ไม่ไล่ตาม" อยู่ที่ §29 ถูก §48 แทนแล้ว
>
> **เอกสารนี้คือจุดตั้งต้นสำหรับคนที่มาต่องานนี้** อ่านจบแล้วรู้ว่าอยู่ตรงไหน ทำอะไรต่อ และห้ามทำอะไร

## สรุป

- **เฟส 0–2 เสร็จ** — `desktop/` **คอมไพล์ได้ครบทั้งสาม OS** และรันสวีททั้งชุดบนทั้งสาม
  ที่เหลือคือแท็บ browser
  > แก้ถ้อยคำ 15 ส.ค. 2026: บรรทัดนี้เคยเขียนว่า "เทสต์ผ่านครบทั้งสาม OS" ซึ่งขัดกับ
  > ย่อหน้าถัดๆ ไปในหน้าเดียวกันที่ระบุเทสต์แดงค้างอยู่ **คอมไพล์ได้** กับ **ผ่าน** ไม่ใช่
  > เรื่องเดียวกัน และหน้านี้ยืนยันอย่างแรกได้ ไม่ใช่อย่างหลัง
- ยกเว้นแท็บ browser ทุกอย่างพร้อมใช้บน Linux: chat, sessions, MCP, skills, files, tools, review, **เทอร์มินัล**
- **macOS รันจริงแล้วใน CI (อัปเดต 7 ส.ค. 2026)** — job `unix (macos-latest)` รันสวีททั้งชุด
  ไม่ใช่แค่ type-check อีกต่อไป และ**มันเจอของจริงสองอย่างที่ Windows กับ Linux ไม่เจอ**:
  `TestCloseKillsBackgroundedGrandchild` (หมดเวลาที่ 5 วิ — วิธีฆ่าลูกหลานของโพรเซสบน
  macOS ไม่เหมือนกัน) และ `TestImageOCRSkillMissingBinaryGivesActionableError` (ใช้เวลา
  13–20 วิ) ทั้งคู่แดงต่อเนื่องมาตั้งแต่อย่างน้อย 4 ส.ค. **ยังไม่มีใครแก้ เพราะวินิจฉัยจาก
  เครื่อง Windows ได้แค่เดา** — ต้องมีคนที่มี Mac จริงมารับ นี่คือรายการแรกของงานพอร์ต macOS
  ไม่ใช่ 3b
  > **แก้ 18 ก.ย. 2026:** job `unix` ถูกถอดออกจาก [ci.yml](.github/workflows/ci.yml) แล้ว
  > (owner's call) บรรทัดข้างบนจึงเป็นหลักฐานจากรอบเดือน ส.ค. 2026 ไม่ใช่สิ่งที่ CI ตรวจให้
  > ทุกวันนี้อีกแล้ว ผลที่ตามมาคือ **ไม่มีที่ไหนรัน `-race` อีกเลย** (verify.sh ข้าม stage
  > `race` เพราะเครื่องเจ้าของไม่มี C compiler) และ **ไม่มีที่ไหนคอมไพล์หรือเทสต์ไฟล์ `!windows`**
  > เหลือเพียง cross-build ของ engine บน Windows ใน ci.yml ซึ่งคอมไพล์อย่างเดียว ไม่รันเทสต์

## ตารางสถานะ

| ส่วน | Windows | Linux | macOS |
|:---|:---|:---|:---|
| `internal/` (engine, tools, providers) | ✅ ใช้จริง | ✅ **เทสต์ + race ผ่านบนเคอร์เนลจริง** | ⚠️ type-check ผ่าน ยังไม่เคยรัน |
| `cmd/aetox` (CLI) + `internal/tui` (จอเต็ม) | ✅ ใช้จริง · จอเต็มทดสอบผ่าน ConPTY และตัวจำลองเทอร์มินัล | ✅ **เทสต์ + race ผ่าน** (จอเต็มยังไม่เคยรันบนเครื่องจริง) | ⚠️ type-check ผ่าน ยังไม่เคยรัน |
| `desktop/` ทั้งแพ็กเกจ | ✅ ใช้จริง | ✅ **`go test ./...` + `-race` ผ่าน** | ⚠️ type-check ผ่าน ยังไม่เคยรัน |
| `desktop/` แท็บ browser | ✅ ใช้จริง | ✅ **ใช้จริงบน WebKitGTK (เฟส 3a, 21 พ.ย. 2026)** — ข้อจำกัดอยู่ใน §เฟส 3a | ⏳ stub (เฟส 3b) |
| binary ภายนอก (tesseract, ffmpeg) | ✅ | ✅ มีข้อความติดตั้งแยก OS อยู่แล้ว | ✅ auto-install ผ่าน brew อยู่แล้ว |

> **ระวังการอ่านตาราง (แก้ 7 ส.ค. 2026):** ✅ ของ Linux = *รันจริงบนเคอร์เนล Linux แล้วผ่าน* ·
> macOS **รันจริงแล้วบน `macos-latest` ใน CI แต่ยังไม่ผ่านครบ** — เทสต์แดงค้างสองตัว (ดูสรุปข้างบน)
> **และตั้งแต่ 18 ก.ย. 2026 CI ไม่มี job นี้แล้ว** — ช่องที่บอก "รันแล้ว" ในตารางนี้จึงเป็นสถานะ ณ เดือน ส.ค. 2026
> ช่อง ⚠️ ในตารางล่างจึงหมายถึง "รันแล้ว ยังไม่เขียว" ไม่ใช่ "ยังไม่เคยรัน" อีกต่อไป ·
> เครื่องเจ้าของยังไม่มี Mac และรัน macOS ใน VM บนเครื่องที่ไม่ใช่ Apple ผิดสัญญาอนุญาต
> จึงแก้ได้เฉพาะผ่าน CI หรือผ่านคนที่มีเครื่องจริง

## เฟส

| # | งาน | สถานะ |
|:--|:---|:---|
| **0** | CI matrix — job `unix` (ubuntu + macos), Linux เพิ่ม `-race` | ✅ **เสร็จ** |
| **1** | `terminal.go` → `ptySession` + `terminal_windows.go` / `terminal_unix.go` (`creack/pty` v1.1.24) | ✅ **เสร็จ** |
| **2** | `browser.go` → `hostBackend`/`tabView` + `browser_windows.go` + `browser_other.go` (stub) | ✅ **เสร็จ** |
| 3a | `browser_linux.go` — WebKitGTK widget ใน `GtkOverlay`/`GtkFixed` | ✅ **เสร็จ 21 พ.ย. 2026** (ดู §เฟส 3a ด้านล่าง) |
| 3b | `browser_darwin.go` — `WKWebView` เป็น subview ของ Wails `NSView` | ยังไม่เริ่ม |
| 4 | packaging (`.deb`/`.rpm`/tar.gz + `.app`/`.dmg`) + `bench.sh` | 🟡 **Linux `.deb` + tar.gz เสร็จ 26 พ.ย. 2026** (ดู §เฟส 4) · `.rpm`, macOS, `bench.sh` ยังไม่เริ่ม |

### เฟส 4 (Linux) — แพ็ก, อัปเดตตัวเอง, ความลับ, process ลูก · 26 พ.ย. 2026

`bash scripts/linux/build.sh` บิลด์แอป + engine ใน Docker แล้วแพ็กสองแบบไว้ที่ `dist/linux/`
และ `scripts/release-local.ps1` เรียกมันผ่าน WSL ทุกรอบ release (เช็ค Docker ก่อนเริ่มบิลด์อะไร)
ทั้งสองไฟล์อยู่ใน `checksums.txt` ที่เซ็นแล้ว และ `cmd/relcheck` จะไม่ปล่อย release ที่ไม่มี tarball

| ไฟล์ | อยู่ที่ไหนบนเครื่อง | อัปเดต |
|:--|:--|:--|
| `aetox_<version>_amd64.deb` | `/opt/aetox/` + เมนู (`/usr/share/applications/aetox.desktop`) · `Depends: libwebkit2gtk-4.1-0, libgtk-3-0t64 \| libgtk-3-0` | การ์ดพาไปหน้า Releases (root เป็นเจ้าของ ไม่มี apt repo) |
| `aetox-linux-amd64-portable.tar.gz` | โฟลเดอร์ `aetox/` ที่ผู้ใช้แตกไว้ มีไฟล์ `.aetox-portable` | กดปุ่มเดียว: ดาวน์โหลด → ตรวจลายเซ็น + hash → rename สลับ app กับ engine → relaunch ผ่าน `/bin/sh` ที่รอ pid เดิมตาย |

ไฟล์ประกอบอยู่ที่ [desktop/build/linux/](desktop/build/linux/README.md) · ชื่อ tarball ถูกจับทั้งชื่อ ไม่ใช่ suffix
เพราะแอป Windows ที่อยู่ในมือคนแล้วหาไฟล์ด้วย `-portable.zip` (`internal/update` `LinuxPortableName`)

**ความลับที่พัก:** `internal/atrest/secret_linux.go` เก็บ key สุ่ม 256 บิตหนึ่งตัวไว้ใน Secret Service
(GNOME Keyring / KWallet / KeePassXC) แล้วซีลไฟล์ด้วย AES-256-GCM (`keyring:v1:` + base64)
เครื่องที่ไม่มี keyring (server, SSH) เขียน plaintext 0600 เหมือนเดิม · `AETOX_KEYRING=off` ปิดได้ ·
ไม่เคยเรียก `dbus-launch` เอง · **ข้อแลกที่ต้องรู้:** ไฟล์ที่ซีลใน desktop session เปิดจาก session
ที่ไม่มี keyring ไม่ได้ จะอ่านเป็น "ยังไม่มี credentials" เหมือน DPAPI blob ที่ย้ายไปอีก account

**process ลูกตายตามแอป:** `proc.HideConsole` (ที่ทุก spawn site เรียกอยู่แล้ว) ใส่ `Pdeathsig=SIGKILL`
บน Linux · `proc.Breakaway` ถอดออกแล้วให้ `Setsid` (waiter ของ update, model server ของผู้ใช้) ·
ครอบคลุมลูกตรงเท่านั้น หลานถูกส่งต่อให้ init — process group ของ `KillOnCancel` ยังเก็บตอน cancel/shutdown

**bug ของ Linux ที่เจอระหว่างทาง:** (1) ทุกที่ที่ตัด `file:///` ทิ้งทั้งก้อนทำให้ path Linux กลายเป็น relative
(`/tmp/x` → `tmp/x`) — deck note, agent pages, การ์ดไฟล์ที่ส่งมอบ ตอนนี้ผ่าน `engine.FileURLPath` ตัวเดียว ·
(2) `SetViewZoom` มีแค่ใน `_windows.go` บิลด์ Linux จึงพังที่ `npm run build` ตั้งแต่ 23 พ.ย. — เพิ่ม stub ที่ตอบ 1
(ยังไม่ zoom จริงบน Linux)

**ยังไม่ได้พิสูจน์:** ติดตั้ง `.deb` บน GNOME/KDE จริง, เมนู/ไอคอน/การจับกลุ่มหน้าต่าง (`StartupWMClass=aetox`
ตรงกับ `ProgramName`), และ self-update ครบวงจรกับ release จริง (เทสต์ครอบ swap + verify แต่ยังไม่มี release ที่มี tarball)

ตรวจเอง: `bash scripts/linux/test-in-docker.sh` — รันเทสต์บนเคอร์เนล Linux พร้อม gnome-keyring จริงใน D-Bus session ส่วนตัว

### เฟส 3a — สิ่งที่ทำเสร็จ และสิ่งที่มันยังไม่เท่า Windows

`desktop/browser_linux.go` + `desktop/browser_gtk_linux.c/.h` ทำ `hostBackend`/`tabView`
ครบ 12 เมธอดบน WebKitGTK: แทรก `GtkOverlay`+`GtkFixed` ครอบ `webview-box` ของ Wails
(หน้าต่างลอยใช้ไม่ได้ — Wayland ให้แอปเลือกตำแหน่ง toplevel ไม่ได้), bridge ชุดเดียวกับ
WebView2 (`window.webkit.messageHandlers.aetox`), snapshot ผ่าน
`webkit_web_view_get_snapshot`, PDF ผ่าน GTK print backend ("Print to File"), zoom,
Inspector, detach เป็นหน้าต่างของตัวเอง, และ callbacks ครบตาม `tabCallbacks`

`callEngine` ตอบทุกเมธอดที่ระบบเรียกจริง โดยแปลงเป็นของ WebKit เอง (ไม่ใช่ CDP):

| เมธอด | ของ WebKit ที่รองรับ |
|:--|:--|
| `Runtime.evaluate` | `webkit_web_view_evaluate_javascript` + envelope ที่ประกอบเป็นรูป CDP |
| `Browser.getVersion` | UA จริงของแท็บ + ชื่อผลิตภัณฑ์ของโฮสต์ |
| `Emulation.setUserAgentOverride` | `WebKitSettings` ของแท็บนั้น (`webkit_web_view_set_settings`) |
| `Emulation.setDeviceMetricsOverride` / `setTouchEmulationEnabled` | user script ต่อแท็บ: `devicePixelRatio`, `maxTouchPoints`, `ontouchstart`, meta viewport |
| `DOM.enable` / `DOM.setFileInputFiles` | handle ที่ envelope แจกให้ DOM node + `DataTransfer` ที่ยัดไบต์จริง (เพดาน 32 MiB) |
| `Input.insertText` / `dispatchKeyEvent` / `dispatchMouseEvent` | event ในหน้า (`setRangeText`, `execCommand`, `KeyboardEvent`, `MouseEvent`) |
| `Page.getLayoutMetrics` / `Page.captureScreenshot` | วัดจาก layout ของหน้า + snapshot `VISIBLE`/`FULL_DOCUMENT` แล้ว crop ตาม clip |
| `Page.printToPDF` | GTK print backend ("Print to File") |

permission ได้คำตอบเท่า Windows ด้วยกลไกของโฮสต์เอง: WebKitGTK ไม่ส่ง flag user-initiated
มาให้ จึงฟัง `button-press-event`/`key-press-event` บน webview (คลิก/คีย์ที่ compositor
ส่งถึงแท็บจริง) และนับว่าภายใน 5 วินาทีคือ "ผู้ใช้กด" — สคริปต์แตะคิวอีเวนต์ของ GDK ไม่ได้
จึงปลอมไม่ได้

สิ่งที่ยังไม่เท่า: event จาก `Input.*` เป็น event ของหน้า ไม่ใช่ของเอนจิน (`isTrusted=false`)
หน้าที่ยึด `isTrusted` จึงยังต้องใช้ทาง JS ของ `browser_keys.go` — ตรงนี้ WebView2 ดีกว่า ·
`setShape` วาดขอบโค้ง/ติ่งด้วย cairo แทนการตัดหน้าต่าง (สีเป็นพื้นของแอป ไม่ใช่การเจาะจริง) ·
`sendBehind` ไม่มีความหมายใน overlay (แท็บถูกจอดนอกหน้าต่างแทน) · และ z-order บน GNOME/KDE
จริงยังต้องดูด้วยตาเหมือนเดิม

ที่ **ดีกว่า Windows**: การถ่ายภาพ/พิมพ์ PDF ของแท็บที่ไม่ถูกเห็นทำงานได้โดยไม่ต้องมีหน้าต่างลอย
นอกจอ — Windows ต้องสร้าง popup ที่ -32000 เพราะ child window เป็นได้อย่างเดียวระหว่าง
"วาดอยู่" กับ "ไม่มีใครเห็น" (browser_windows.go openTab) ส่วนที่นี่ web process วาดภาพเอง
ไม่สนใจว่าวิดเจ็ตอยู่ตรงไหน · และแคปทั้งหน้าได้จาก `SNAPSHOT_REGION_FULL_DOCUMENT` ครั้งเดียว

หลักฐาน: `desktop/browser_linux_live_test.go` (แท็ก `aetoxlivetest` +
`AETOX_LINUX_BROWSER_LIVE=1`) รันบนจอ WSLg ผ่าน Docker — โหลดหน้า, อ่านข้อความด้วย
`Runtime.evaluate`, snapshot 10,311 ไบต์, PDF 738,801 ไบต์, และเมธอดที่ไม่รองรับถูกปฏิเสธ
พร้อมชื่อ; `go vet -tags webkit2_41 ./desktop/` ผ่าน และ `scripts/linux/build.sh` บิลด์แอปได้
(80.7 MB, 3m46s)

## แผนที่ไฟล์ — ไฟล์ไหนของ OS ไหน

กฎ: **แยกไฟล์ก็ต่อเมื่อ type หรือ import ต่างกัน** ถ้าต่างแค่ชื่อคำสั่งที่จะรัน ใช้ `switch runtime.GOOS` ในไฟล์เดิม (เช่น `presets.go`/`subagents.go` ที่เลือก `explorer`/`open`/`xdg-open`) และ**ไม่แยกโฟลเดอร์** เพราะ Go นับ 1 โฟลเดอร์ = 1 package ส่วน `App` ต้องอยู่ใน `package main` เดียว ไม่งั้น Wails generate `App.d.ts` ออกมาคนละชุด แล้ว `BrowserPane.svelte` พังตอน `vite build` (§48 Decision 2)

| ไฟล์ | build tag | Win | Linux | mac | มีอะไร |
|:---|:---|:-:|:-:|:-:|:---|
| `terminal.go` | — | ✅ | ✅ | ✅ | `ptySession`, `App.Terminal*`, ลูปอ่าน |
| `terminal_windows.go` | ชื่อไฟล์ | ✅ | | | ConPTY + รายชื่อเชลล์ของ Windows |
| `terminal_unix.go` | `//go:build unix` | | ✅ | ✅ | `creack/pty` + `$SHELL` + กวาด session |
| `browser.go` | — | ✅ | ✅ | ✅ | `tabView`/`hostBackend`/`tabCallbacks`, JS builders, `sameOrigin`, `onMessage`, `navCompleted`, `App.Browser*` |
| `browser_windows.go` | ชื่อไฟล์ | ✅ | | | Win32 + WebView2 (`win32Host`/`win32Tab`) |
| `browser_other.go` | `//go:build !windows` | | ⏳ | ⏳ | stub ชั่วคราว — ลบทิ้งเมื่อ 3a และ 3b เสร็จทั้งคู่ |
| `browser_linux.go` | ชื่อไฟล์ | | ✅ | | WebKitGTK (เฟส 3a) |
| `browser_darwin.go` | ชื่อไฟล์ | | | ✅ | WKWebView (เฟส 3b) |
| `computer_tool.go`, `computer_guard.go`, `computer_refs.go`, `computer_errors.go` | — | ✅ | ✅ | ✅ | pack, ด่าน, ตาราง ref, คำศัพท์ของความล้มเหลว — ไม่มี Win32 เลย |
| `computer_reach_windows.go`, `computer_act_windows.go`, `uia_windows.go`, `computer_overlay_windows.go` | ชื่อไฟล์ | ✅ | | | UI Automation ผ่าน COM ที่เขียน binding เอง + SendInput + PrintWindow + หน้าต่างแสงครอบจอ |
| `computer_reach_other.go`, `computer_act_other.go`, `computer_overlay_other.go` | `//go:build !windows` | | ⏳ | ⏳ | ปฏิเสธอย่างชัดเจน ไม่ใช่ no-op เงียบ — no-op ที่คืนลิสต์ว่างคือการบอกโมเดลว่าเครื่องนี้ไม่มีหน้าต่าง ซึ่งเป็นคำโกหกที่มันจะเอาไปทำต่อ |

`_unix.go` ต้องเขียน `//go:build unix` เอง เพราะ Go ไม่ถือว่า `_unix` เป็นคำลงท้ายพิเศษเหมือน `_windows`/`_linux`/`_darwin` — ใช้ `unix` แทน `!windows` (แบบที่ `internal/proc` ใช้) เพราะตรงกับชื่อไฟล์และไม่ลากไป js/wasm กับ plan9 ที่ `creack/pty` ไม่รองรับ

**การพอร์ต `computer` ไม่ใช่การพอร์ตไฟล์** UI Automation เป็น API ของ Windows และของที่เทียบเท่าบน macOS กับ Linux เป็นคนละตัวที่มี permission model ของตัวเอง (macOS ต้องได้ Accessibility กับ Screen Recording ในการตั้งค่าระบบก่อนถึงจะทำอะไรได้เลย) ครึ่งที่ไม่มี Win32 — pack, ด่าน, ตาราง ref, การจำแนกความล้มเหลว — คอมไพล์และทดสอบได้ทุกแพลตฟอร์มอยู่แล้ว งานที่เหลือจึงเป็นการเขียน `reach*` กับ `overlay` ของ OS นั้น ไม่ใช่การรื้อ ดู [docs/internal/architecture/computer-use-2026-09-07.md](docs/internal/architecture/computer-use-2026-09-07.md) §5

## ทำเฟส 3a/3b ยังไง — ต้อง implement แค่ 2 interface

อยู่ใน `desktop/browser.go` ทั้งคู่ ที่เหลือทั้งไฟล์ใช้ร่วมกันหมด **ไม่ต้องแตะ**

```go
type hostBackend interface {
    start() error                  // ตั้ง UI thread / message pump ของแพลตฟอร์ม, idempotent
    do(fn func())                  // รัน fn บน thread ที่เป็นเจ้าของ webview
    openTab(id, url string, x, y, w, h int, cb tabCallbacks) tabView
}

type tabView interface {
    navigate(url string)
    eval(js string)
    setBounds(x, y, w, h int)
    setVisible(visible bool)       // true = แสดง + ยกขึ้นบนสุด
    setZoom(factor float64)
    openDevTools()
    destroy()
}
```

แล้วให้ `newHostBackend() hostBackend` ในไฟล์ของ OS นั้นคืนตัวจริง (ตัด GOOS นั้นออกจาก build tag ของ `browser_other.go` ด้วย)

**ห้ามพลาดสามข้อนี้:**

1. **`do()` ต้อง async เสมอ** — `g_idle_add` / `dispatch_async` **ห้าม `dispatch_sync`** เพราะ `browserSnapshot` เรียก `do()` แล้วบล็อกรอ channel 5 วินาที ถ้า sync และถูกเรียกจาก main thread = **deadlock ที่ path ซึ่ง AI ใช้อ่านหน้าเว็บ** unit test จับไม่ได้ (§48 Decision 3)
2. **`bridgePost` มีอยู่แล้วใน `browser_other.go`** = `window.webkit.messageHandlers.aetox.postMessage` ถูกต้องสำหรับ WebKit ทั้งสองตัว ชื่อ handler ต้องตรงกับที่ลงทะเบียนกับ `WebKitUserContentManager` / `WKUserContentController`
3. **ยึด 5 กฎท้าย [docs/internal/architecture/native-browser-embedding-2026-07-24.md](docs/internal/architecture/native-browser-embedding-2026-07-24.md)** — เขียนจากรอบดีบักจริงที่ได้ failure catalog 7 ข้อ ไม่ใช่ทฤษฎี โดยเฉพาะกฎข้อ 1 (ห้ามหาหน้าต่างตัวเองจาก global)

Linux ต้องใช้ `-tags webkit2_41` — Ubuntu 24.04 ไม่มี `libwebkit2gtk-4.0-dev` แล้ว (วัดแล้ว)

## วิธีตรวจซ้ำเอง

เครื่อง dev เป็น Windows และไม่มี C compiler → **พิสูจน์อะไรของ Linux/mac บนเครื่องตรงๆ ไม่ได้** ที่ใช้อยู่จริงคือ Docker (ต้องเปิด Docker Desktop ก่อน):

```bash
# ทั้งชุดบนเคอร์เนล Linux จริง — copy ก่อนเพื่อไม่ให้เขียนทับรีโป
docker run --rm -v "E:/Aetox/Aetox:/repo:ro" -v aetox-gomod:/go/pkg/mod golang:1.25 \
  sh -c 'cp -r /repo /tmp/a && cd /tmp/a && go test -count=1 ./...'

# race detector — เช็คที่ verify.sh skip มาตลอดเพราะเครื่อง dev ไม่มี C compiler
  ... && go test -count=1 -race -timeout 15m ./...'

# macOS ตรวจได้แค่ type-check เท่านั้น
  ... && GOOS=darwin GOARCH=arm64 go vet ./...'
```

จาก Git Bash ต้องใส่ `MSYS_NO_PATHCONV=1` นำหน้า ไม่งั้น path ฝั่งคอนเทนเนอร์โดนแปลงเป็น path Windows

## ผลวัดจริง (2026-07-27)

| เช็ค | ผล |
|:---|:---|
| Windows `go test ./...` | ✅ 24 แพ็กเกจ |
| **Linux `go test ./...`** | ✅ **24 แพ็กเกจ รวม `desktop`** |
| **Linux `go test -race ./...`** | ✅ **24 แพ็กเกจ ไม่มี data race รวม `desktop`** |
| `GOOS=darwin go vet ./...` | ✅ สะอาด |
| `TestShellSkillCancelKillsGrandchild` | ✅ `tree_other.go` (`Setpgid` + `kill(-pgid)`) พิสูจน์กับ process หลานจริง |
| `TestCloseKillsBackgroundedGrandchild` | ✅ การกวาด session ของเทอร์มินัลพิสูจน์กับ process ที่ถูก detach จริง |
| `libwebkit2gtk-4.0-dev` บน Ubuntu 24.04 | ❌ ไม่มีแล้ว เหลือ 4.1 |

## สิ่งที่พอร์ตนี้ขุดเจอ — บั๊กที่ Windows ปิดบังไว้

1. **`kill(-pgid)` ไม่พอสำหรับเชลล์ interactive** — เชลล์โยนงาน background ไป process group ใหม่เสมอ (bash และ dash, เปิด/ปิด job control) แต่ยังอยู่ session เดิม `npm run dev &` เลยรอดจากการปิดแท็บและยึดพอร์ตไว้ แก้ด้วยการกวาดทั้ง session **ตรวจแล้วว่าไม่ลามไป `proc.KillOnCancel`** เพราะ `sh -c` ไม่มี tty จึงไม่มี job control (§48 Decision 5)
2. **เทสต์ที่แยก config dir ด้วย `APPDATA`/`LOCALAPPDATA` ไม่ได้แยกอะไรเลยนอก Windows** — 9 เทสต์ไปอ่านเขียน `~/.config/aetox` ตัวจริง แก้ด้วย `isolateUserDirs(t)` ต่อแพ็กเกจ (§48 Decision 6)
3. **`TestProjectKeyStableAndDistinct` ผ่านมาตลอดด้วยเหตุผลที่ผิด** — เช็ค `filepath.Base` กับ `C:\projects\app` ที่ hardcode ซึ่งนอก Windows เป็นชื่อเดียวไม่มีตัวคั่น
4. **`win32Host.start()` ที่ล้มเหลวทำให้ทั้งระบบค้าง** — ทิ้ง `started=true` ไว้โดยไม่ปิด `ready` และ `browserHostLazy` เรียก `start()` ทุก binding แก้แล้ว

## สิ่งที่ยังไม่รู้ และรู้ได้ทางเดียว

1. **macOS ทุกอย่าง** — เฟส 3b จะ iterate ช้าที่สุด (แก้ 1 บรรทัด = รอ CI 1 รอบ ไม่เห็นหน้าจอ) นี่คือเหตุผลที่ทำ Linux ให้จบก่อน
2. **z-order / compositing บน WM จริง** — Docker/WSLg/xvfb บอกได้แค่ "รันได้ วาดออก" บอกไม่ได้ว่าบน GNOME/KDE/tiling WM หน้าต่างลูกอยู่ถูกชั้น ซึ่งเป็นหัวใจทั้งหมดของงาน **ต้องมี Linux desktop จริงหรือ VM**
3. **`proc.KillTreeOnExit` เป็น no-op นอก Windows** — ไม่มี Job Object · บน Linux ลูกตรงทุกตัวได้ `Pdeathsig` แทนแล้ว (26 พ.ย. 2026, §เฟส 4) ที่เหลือคือหลานตอนแอปถูก force-kill
4. **`WebviewUserDataPath` มีเฉพาะ `options.Windows`** — กฎ §14 บังคับกับ webview ของตัวแอปไม่ได้นอก Windows แต่**แท็บ browser ยังบังคับได้ทุก OS** เพราะเป็นโค้ดที่เราคุมเอง (`WebKitWebsiteDataManager`/`WKWebsiteDataStore` รับ path ตรงๆ)
5. **Wails v2.13 ไม่เปิด native window handle** — ไม่มี `NativeWindowHandle`/`GetNativeHandle` ในโมดูลเลย กฎข้อ 1 ของ blueprint จึงทำตามไม่ได้ถ้าไม่ patch Wails เป็น fork ตัวที่ 3 ใน `third_party/` (§48 ท้ายสุด)

## การจัดจำหน่าย — ตัวเลขที่ต้องปกป้อง

[BENCHMARK.md](docs/reports/BENCHMARK.md) §4 วัด Aetox v1.10.22 เมื่อ 28 ต.ค. 2026 ได้ **123.9 MiB** (สองไฟล์ตั้งแต่ 1.6.0) ตัวคูณพื้นที่ของเครื่องมือในตารางเป็น 1–23 เท่าเมื่อปัดเป็นจำนวนเต็ม คู่แข่งเป็นข้อมูลเดิม 27 ก.ค. 2026 ด้วยเหตุผลว่า *ใช้ webview ที่ OS มีอยู่แล้ว ไม่ได้แบก Chromium มาเอง* — เหตุผลนั้นยังจริงบน macOS เสมอ แต่**บน Linux ขึ้นกับวิธีแพ็ก**:

| แบบ | ขนาด | สรุป |
|:---|:---|:---|
| `.deb`/`.rpm`/tar.gz ประกาศ `Depends: libwebkit2gtk-4.1-0` | ~50 MB | ✅ **เลือกอันนี้** |
| AppImage บรรจุ WebKitGTK มาด้วย | ~150 MB+ | ❌ ลบจุดขายหลักทิ้ง |
| Flatpak | ~40 MB | ❌ sandbox ตีกับตัวตนของสินค้า (รัน shell/git/MCP ทั้งเครื่อง) |

macOS มีค่าใช้จ่ายที่ไม่ใช่โค้ด: แอปไม่เซ็นชื่อโดน Gatekeeper บล็อก ถ้าจะให้เนียนต้องมี Apple Developer Program **$99/ปี** — ตัดสินใจได้ตั้งแต่ตอนนี้

`bench.ps1` เป็น PowerShell ล้วน → เฟส 4 ต้องมี `bench.sh` ไม่งั้นตัวเลข Linux/mac ขึ้น README ไม่ได้ตามกติกาข้อเดียวของ BENCHMARK.md
