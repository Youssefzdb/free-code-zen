# Free Claude Code — Zen API Edition

Claude Code معدّلة تشتغل مع OpenCode Zen API مجاناً — بدون API key، بدون telemetry، بدون قيود.

## المميزات

- ✅ كل الـ telemetry مشال
- ✅ كل الـ guardrails مشالة
- ✅ كل الـ experimental features مفعّلة (54 flag)
- ✅ 5 موديلات مجانية من Zen API
- ✅ Tools شغّالة (Bash, Read, Write, Edit)
- ✅ اعتراض fetch لتحويل صيغة الـ tools تلقائياً

## الموديلات المجانية

| Model | الوصف |
|-------|--------|
| `deepseek-v4-flash-free` | الأفضل للبرمجة (default) |
| `mimo-v2.5-free` | سريع وخفيف |
| `ling-3.0-flash-free` | متوازن |
| `nemotron-3-ultra-free` | reasoning قوي |
| `laguna-s-2.1-free` | خفيف جداً |

## التثبيت السريع

```bash
# 1. تنزيل المشروع
git clone https://github.com/USER/free-code-zen.git
cd free-code-zen

# 2. تثبيت Bun (إذا لم يكن مثبتاً)
npm install -g bun

# 3. تثبيت الاعتمادات وبناء
bun install
bun run build:dev:full

# 4. تشغيل
./go "your prompt"
# أو وضع تفاعلي
./go
```

## الاستخدام

```bash
# وضع تفاعلي
./go

# أمر مباشر
./go "list files in current directory"

# تغيير الموديل
ZEN_MODEL=nemotron-3-ultra-free ./go "explain this code"

# أو باستخدام الـ binary مباشرة
ANTHROPIC_API_KEY=public \
ANTHROPIC_BASE_URL=https://opencode.ai/zen \
./cli-dev -p "your prompt" --model deepseek-v4-flash-free
```

## أوامر التشغيل

```bash
# بناء سريع
bun run build:dev:full

# بناء production
bun run build

# تشغيل من المصدر (بدون بناء)
bun run dev
```

## ملاحظات

- Zen API مجاني ومفتوح — لا يحتاج API key
- الـ tools تحوّل تلقائياً من صيغة Anthropic إلى OpenAI عبر fetch interceptor
- Streaming يتحوّل إلى non-stream ويُعاد تحويله لـ SSE تلقائياً
