# MegaVerse Messenger

پیامرسان مدرن با معماری میکروسرویس

## اجرا

### سرور (Go)

```bash
cd services/messaging
go mod tidy
go run main.go
```

سرور روی پورت `8080` اجرا میشه.

### وب (Next.js)

```bash
cd apps/web
npm install
npm run dev
```

وب روی پورت `3000` اجرا میشه.

## ویژگی‌ها

- ✅ ثبت نام و ورود
- ✅ چت خصوصی
- ✅ پیام‌رسانی real-time با WebSocket
- ✅ نمایش کاربران آنلاین
- ✅ رابط کاربری مدرن و فارسی
- ✅ ذخیره‌سازی با SQLite (بدون نیاز به Docker)

## تکنولوژی‌ها

| لایه | تکنولوژی |
|------|----------|
| Backend | Go + SQLite + WebSocket |
| Frontend | Next.js + React + Tailwind |
| Database | SQLite |
| Protocol | WebSocket |
