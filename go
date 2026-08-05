#!/usr/bin/env bash
# === Free Claude Code (Zen API) Launcher ===
# Usage: ./go "your prompt here"
# Or:    ./go  (interactive mode)

export ANTHROPIC_API_KEY="public"
export ANTHROPIC_BASE_URL="https://opencode.ai/zen"
export ANTHROPIC_MODEL="deepseek-v4-flash-free"
export ANTHROPIC_SMALL_FAST_MODEL="deepseek-v4-flash-free"
export DISABLE_TELEMETRY=1
export DISABLE_ERROR_REPORTING=1
export CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC=1

# Allow running as root
export IS_SANDBOX=1

# Available free models:
# deepseek-v4-flash-free  (best for coding)
# mimo-v2.5-free          (fast)
# ling-3.0-flash-free     (balanced)
# nemotron-3-ultra-free   (reasoning)
# laguna-s-2.1-free       (lightweight)

MODEL="${ZEN_MODEL:-deepseek-v4-flash-free}"

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"

if [ -f "$SCRIPT_DIR/cli-dev" ]; then
  BIN="$SCRIPT_DIR/cli-dev"
elif [ -f "$SCRIPT_DIR/cli" ]; then
  BIN="$SCRIPT_DIR/cli"
else
  echo "Building..."
  cd "$SCRIPT_DIR" && bun run build:dev:full 2>&1 | tail -1
  BIN="$SCRIPT_DIR/cli-dev"
fi

if [ -n "$1" ]; then
  "$BIN" -p "$*" --model "$MODEL" 2>&1
else
  "$BIN" --model "$MODEL" 2>&1
fi
