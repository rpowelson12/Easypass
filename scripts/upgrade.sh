
#!/bin/bash

set -e

if ! command -v curl &> /dev/null; then
  echo "❌ curl is required but not installed. Please install curl and retry."
  exit 1
fi

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

curl -fsSL --max-time 30 -o "$TMP" "$URL"

if [[ ! -s "$TMP" ]]; then
  echo "❌ Download failed or binary is empty."
  exit 1
fi

chmod +x "$TMP"

if [ -f /usr/local/bin/easypass ]; then
  echo "💾 Backing up current Easypass binary..."
  sudo cp /usr/local/bin/easypass /usr/local/bin/easypass.bak-$(date +%Y%m%d%H%M%S)
fi

echo "🔁 Replacing existing Easypass binary..."
sudo mv "$TMP" /usr/local/bin/easypass

echo "✅ Easypass upgraded!"
