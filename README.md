# Free Claude Code — Zen API Edition

Claude Code معدّلة تشتغل مع OpenCode Zen API مجاناً — بدون API key، بدون telemetry، بدون قيود.

## التثبيت السريع (Binary جاهز — بدون bun install)

```bash
# 1. تنزيل الـ binary
wget https://github.com/Youssefzdb/free-code-zen/releases/download/v1.0.0/cli-dev.gz
gunzip cli-dev.gz
chmod +x cli-dev

# 2. تشغيل
ANTHROPIC_API_KEY=public ANTHROPIC_BASE_URL=https://opencode.ai/zen ./cli-dev -p "say hello" --model deepseek-v4-flash-free
```

## أو التثبيت من المصدر

```bash
git clone https://github.com/Youssefzdb/free-code-zen.git
cd free-code-zen
rm -rf node_modules
rm -rf ~/.bun/install/cache
bun install
bun run build:dev:full
./go "say hello"
```

## الموديلات المجانية

| Model | الوصف |
|-------|--------|
| `deepseek-v4-flash-free` | الأفضل للبرمجة (default) |
| `mimo-v2.5-free` | سريع وخفيف |
| `ling-3.0-flash-free` | متوازن |
| `nemotron-3-ultra-free` | reasoning قوي |
| `laguna-s-2.1-free` | خفيف جداً |

## الاستخدام

```bash
# أمر مباشر
ANTHROPIC_API_KEY=public ANTHROPIC_BASE_URL=https://opencode.ai/zen ./cli-dev -p "your prompt" --model deepseek-v4-flash-free

# وضع تفاعلي
ANTHROPIC_API_KEY=public ANTHROPIC_BASE_URL=https://opencode.ai/zen ./cli-dev --model deepseek-v4-flash-free

# تغيير الموديل
ANTHROPIC_API_KEY=public ANTHROPIC_BASE_URL=https://opencode.ai/zen ./cli-dev -p "analyze this" --model nemotron-3-ultra-free
```

## المميزات

- ✅ كل الـ telemetry مشال
- ✅ كل الـ guardrails مشالة
- ✅ كل الـ experimental features مفعّلة
- ✅ 5 موديلات مجانية من Zen API
- ✅ Tools شغّالة (Bash, Read, Write, Edit)
- ✅ اعتراض fetch لتحويل صيغة الـ tools تلقائياً
