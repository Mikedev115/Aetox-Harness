# Aetox Mobile Android 0.3.21

From this version on, the app's screen ships inside the APK itself. It no longer loads from Aetox on the PC, so a fix to the phone's screen reaches you as soon as the APK updates, without waiting for a new Aetox on the PC.

## What's new

- **The screen comes with the app.** Before, every page in the app came from the Aetox `.exe` on the PC. Even a fix to the back button had to wait for a new desktop release (owner, 24 Nov 2026: "มือถือควรจะแยกเวอร์ชั่นของตัวเอง", mobile should have its own version).
- **The PC still does everything else.** It answers chat, files, voice, and every `/api/*` request, at the same address and on the same pairing as before. You don't need to pair again.
- **No stale copy.** The service worker that the PC's page left on the phone no longer serves an old copy of the screen over the app's own. The app opens with `/?shell=app` and removes that old worker.
- **iPhone and browsers are unchanged.** They still use the page from the PC.

Fixes from earlier versions that lived in the page reach the app with this version too, even if the PC has not been updated: the back button asks the page first, and the "Check for updates" button asks the published release.

## Requirements

Android 10 or later. Installs over 0.3.20 (same signing key `c6330d76…`).

## Test status

- `npm --prefix mobile run test`: 202 tests passed in 31 files, including a new `pwa.test.ts` case ("no worker in the app, and the old one is removed"). The 3 uncaught `element.animate is not a function` errors are the same jsdom issue as before.
- `npm --prefix mobile run build:android` passed: gradle unit tests (46, including the new `PagePathTest`), lint, and the signature check.
- The APK contains `assets/page` (index.html, assets, icons, manifest), with no `sw.js` and no previous APK inside.
- `apksigner verify` passed. Cert SHA-256 `c6330d76ba48be84e4b7bbbc3928a6ee486fb234ae46012c78c6f8423ee7fb5b` matches 0.3.20.
- APK: 8,765,942 bytes. SHA-256 `1d11c7d1bb8c50446d9b02ea1ede9ffb2bddd7f79aceb8bc0059f9cb0a6d0286`.
- **Not yet tested on a real Android screen.** It hasn't been confirmed on a phone that the first open after the update shows the app's own page instead of the cached copy.
