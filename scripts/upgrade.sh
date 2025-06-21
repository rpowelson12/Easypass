
#!/bin/bash

set -e

OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [[ "$ARCH" == "x86_64" ]]; then
  ARCH="amd64"
elif [[ "$ARCH" == "arm64" || "$ARCH" == "aarch64" ]]; then
  ARCH="arm64"
fi

BINARY="easypass_${OS}_${ARCH}"
URL="https://github.com/rpowelson12/Easypass/releases/latest/download/$BINARY"

echo "📦 Downloading latest Easypass from $URL..."

TMP=$(mktemp)
curl -fsSL -o "$TMP" "$URL"

if [[ ! -s "$TMP" ]]; then
  echo "❌ Download failed or binary is empty."
  exit 1
fi

chmod +x "$TMP"

echo "🔁 Replacing existing Easypass binary..."
sudo mv "$TMP" /usr/local/bin/easypass

echo "✅ Easypass upgraded!"
