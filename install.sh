#!/usr/bin/env bash
set -e

echo "=== Free Claude Code — Zen API Edition ==="
echo ""

# Check for bun
if ! command -v bun &>/dev/null; then
  echo "[1/4] Installing Bun..."
  curl -fsSL https://bun.sh/install | bash
  export BUN_INSTALL="$HOME/.bun"
  export PATH="$BUN_INSTALL/bin:$PATH"
fi

echo "[1/4] Bun ready: $(bun --version)"

echo "[2/4] Installing dependencies..."
bun install --frozen-lockfile 2>/dev/null || bun install

echo "[3/4] Building..."
bun run build:dev:full

echo "[4/4] Done!"
echo ""
echo "Run: ./go \"your prompt\""
echo "Or:  ./go  (interactive mode)"
