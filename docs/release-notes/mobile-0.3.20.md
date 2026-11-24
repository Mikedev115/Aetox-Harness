# Aetox Mobile Android 0.3.20

รุ่นนี้แก้เรื่องที่เจอทันทีหลังปล่อย 0.3.19: กด «ตรวจสอบอัปเดต» แล้วแอปตอบว่าล่าสุดแล้ว ทั้งที่มีรุ่นใหม่กว่าออกไปแล้ว

## มีอะไรใหม่

- **ปุ่มตรวจสอบอัปเดตถามรุ่นที่เผยแพร่จริงก่อน** — เดิมแถว Settings › เวอร์ชันแอปถามแค่ `/install/meta` ของ PC ซึ่ง APK ที่มากับโปรแกรมบน PC เป็นรุ่นเก่ากว่าที่ปล่อยแล้ว (1.11.36 ถือ 0.3.18) จึงตอบว่า «ล่าสุดแล้ว» (เจ้าของ 24 พ.ย. 2026: "ไม่เห็นขึ้นอัปเดต")
- ตอนนี้ถาม `mobile-latest.json` ทันที ไม่ติดเพดานวันละครั้งแบบตัวเช็กเบื้องหลัง — `NativeBridge.latest()` ตอบกลับทาง `window.__aetoxLatest` (`newer:<code>:<name>` หรือ `none`)
- กดติดตั้งแล้วโหลด APK จาก release ตาม URL และ SHA-256 ที่ manifest ระบุ ฝั่ง native เป็นคนโหลด (`NativeBridge.installLatest()`) หน้าไม่ต้องส่งอะไรไป และรายงานทุกขั้นทาง `__aetoxUpdate` เหมือนทางของ PC
- ถ้าถาม release ไม่ได้ จะถอยไปถาม PC ตามเดิม ไม่ใช่ขึ้นว่าล่าสุดแล้ว

## ต้องใช้

Android 10 ขึ้นไป · ติดตั้งทับ 0.3.19 ได้ (คีย์เซ็นเดิม `c6330d76…`)

การเช็กและติดตั้งอัปเดตทำงานในตัว APK เอง ไม่ต้องรอ Aetox บน PC รุ่นใหม่

## สถานะการทดสอบ

- `npm --prefix mobile run test` — 201 tests ผ่าน 31 ไฟล์ รวม `appUpdate.test.ts` ที่เพิ่มเคส «release ใหม่กว่าที่ PC ถือ» · 3 uncaught `element.animate is not a function` จาก `tests/message.test.ts` เป็นอาการเดิมของ jsdom ที่มีอยู่ก่อนรุ่นนี้ (0.3.19 ผ่าน 199)
- `npm --prefix mobile run build:android` ผ่าน — gradle unit test + lint + ตรวจลายเซ็น
- `apksigner verify` ผ่าน cert SHA-256 `c6330d76ba48be84e4b7bbbc3928a6ee486fb234ae46012c78c6f8423ee7fb5b` ตรงกับ 0.3.19
- APK 8,059,345 ไบต์ · SHA-256 `142ca65ec314d6a1c3e33af97e4e623f7291cb20a52b31ffefcd1d236dff4d35`
- **ยังไม่ได้กดบนจอ Android จริง** — ทางเช็กอัปเดตยังไม่ได้เห็นด้วยตา
