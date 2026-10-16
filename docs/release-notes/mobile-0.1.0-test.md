# Aetox Mobile Android 0.1.0-test — Preview

**ไฟล์ APK สำหรับทดลอง Aetox Mobile บน Android โดยให้ Windows เป็นผู้ทำงานและเก็บข้อมูล**

## ก่อนดาวน์โหลด

- เป็น **test-signed pre-release** ไม่ใช่รุ่น Play Store หรือรุ่น production
- ต้องใช้ Android 10 ขึ้นไป และ Tailscale บนมือถือกับ Windows
- ฝั่ง Windows ต้องเป็น development build ที่มี **Aetox Link** และตั้ง trusted HTTPS แล้ว รุ่น desktop stable เดิมอาจยังไม่มีฟีเจอร์นี้
- Windows และ Aetox ต้องเปิดอยู่ ไม่รับรอง headless, background execution หรือ push

## การเริ่มใช้

1. ดาวน์โหลด `Aetox-Mobile-Android-Test.apk` จาก Assets แล้วเปิดติดตั้งบน Android
2. เปิดแอป แล้วใส่ HTTPS URL หรือ invitation URL ของ Aetox Link บน Windows
3. ขอจับคู่ และให้เจ้าของเครื่อง Windows ตรวจรหัส/เลือกบทสนทนาที่แชร์/ยืนยันอุปกรณ์

การติดตั้งอย่างเดียว **ไม่ให้สิทธิ์อ่านแชตหรือสั่งเครื่อง** และ APK ไม่ข้าม certificate errors

## สถานะการทดสอบ

ผ่าน Android build, JVM policy tests, lint, APK signature verification และ manifest checks
ยังไม่ผ่าน end-to-end installation/pairing บนมือถือจริง จึงเผยแพร่แยกจาก Latest ของ Windows

Package: `io.github.mikedev115.aetox.mobile.dev` · Version: `0.1.0-test`

APK SHA-256:
`095de599868f20fa4c2e7509d52a29ccb12522fdcab75a1e8b17026a0f7717fb`

ไฟล์ iPhone Web Clip ต้องกำหนด URL ของ Windows แต่ละเครื่อง จึงไม่ได้รวมโปรไฟล์ส่วนตัวไว้ใน release สาธารณะนี้
