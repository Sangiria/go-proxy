#!/usr/bin/env bash
set -euo pipefail

PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

BIN_DIR="${PROJECT_ROOT}/bin"
XRAY_PATH="${BIN_DIR}/xray"

command -v curl >/dev/null 2>&1 || { echo "curl is required"; exit 1; }
command -v unzip >/dev/null 2>&1 || { echo "unzip is required"; exit 1; }

mkdir -p "$BIN_DIR"

ARCH_RAW="$(uname -m)"

case "$ARCH_RAW" in
  x86_64|amd64)
    XRAY_ASSET="Xray-linux-64.zip"
    ;;
  *)
    echo "Unsupported architecture: $ARCH_RAW"
    exit 1
    ;;
esac

echo "Detected architecture: $ARCH_RAW"
echo "Using asset: $XRAY_ASSET"

TMP_DIR="$(mktemp -d)"
cleanup() { rm -rf "$TMP_DIR"; }
trap cleanup EXIT

REL_JSON="$TMP_DIR/release.json"
curl -sSL \
  "https://api.github.com/repos/XTLS/Xray-core/releases/latest" \
  -o "$REL_JSON"

ASSET_URL="$(
  grep -oE '"browser_download_url":\s*"[^"]+"' "$REL_JSON" \
  | sed -E 's/.*"([^"]+)"/\1/' \
  | grep -E "${XRAY_ASSET}$" \
  | head -n 1 || true
)"

if [[ -z "$ASSET_URL" ]]; then
  echo "Failed to find asset ${XRAY_ASSET} in Xray v${XRAY_VERSION}"
  echo "Check manually:"
  echo "https://github.com/XTLS/Xray-core/releases/tag/v${XRAY_VERSION}"
  exit 1
fi

echo "Downloading: $ASSET_URL"

ZIP_PATH="$TMP_DIR/xray.zip"
curl -L "$ASSET_URL" -o "$ZIP_PATH"

UNPACK_DIR="$TMP_DIR/unpacked"
mkdir -p "$UNPACK_DIR"
unzip -q "$ZIP_PATH" -d "$UNPACK_DIR"

cp -f "$UNPACK_DIR/xray" "$XRAY_PATH"
chmod +x "$XR

# for debugging. will delete later
echo "Xray installed at: $XRAY_PATH"
"$XRAY_PATH" version || true