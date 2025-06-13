

#!/usr/bin/env bash

set -e

REPO="rpowelson12/Easypass"
BIN_NAME="easypass"

# Get latest release tag from GitHub API
LATEST_TAG=$(curl -s https://api.github.com/repos/$REPO/releases/latest \
  | grep '"tag_name":' \
  | sed -E 's/.*"([^"]+)".*/\1/')

if [[ -z "$LATEST_TAG" ]]; then
  echo "❌ Failed to fetch latest release version"
  exit 1
fi

echo "📦 Installing Easypass version: $LATEST_TAG"

OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Normalize OS for Windows under Git Bash/MSYS/Cygwin
if [[ "$OS" == "mingw"* ]] || [[ "$OS" == "msys"* ]] || [[ "$OS" == "cygwin"* ]]; then
  OS="windows"
fi

# Normalize architecture names
if [[ "$ARCH" == "x86_64" ]]; then
  ARCH="amd64"
elif [[ "$ARCH" == "arm64" ]] || [[ "$ARCH" == "aarch64" ]]; then
  ARCH="arm64"
else
  echo "❌ Unsupported architecture: $ARCH"
  exit 1
fi

echo "🔍 Detected OS: $OS"
echo "🔍 Detected Arch: $ARCH"

if [[ "$OS" == "windows" ]]; then
  ASSET="${BIN_NAME}_${OS}_${ARCH}.exe.zip"
else
  ASSET="${BIN_NAME}_${OS}_${ARCH}.tar.gz"
fi

URL="https://github.com/${REPO}/releases/download/${LATEST_TAG}/${ASSET}"

echo "⬇️ Downloading: $URL"

TMPDIR=$(mktemp -d)
cd "$TMPDIR"

curl -L -o "$ASSET" "$URL"

if [[ "$OS" == "windows" ]]; then
  unzip "$ASSET"
  # Ensure ~/bin exists
  INSTALL_DIR="$HOME/bin"
  mkdir -p "$INSTALL_DIR"
  cp "${BIN_NAME}_${OS}_${ARCH}.exe" "$INSTALL_DIR/$BIN_NAME.exe"
  chmod +x "$INSTALL_DIR/$BIN_NAME.exe"
  echo "📂 Installed $BIN_NAME.exe to $INSTALL_DIR"
  echo "❗ Make sure $INSTALL_DIR is in your PATH environment variable."
else
  echo "📦 Extracting $ASSET..."
  tar -xzf "$ASSET"
  if [[ ! -f "$BIN_NAME" ]]; then
    echo "❌ Binary '$BIN_NAME' not found in archive"
    exit 1
  fi
  chmod +x "$BIN_NAME"
  echo "📂 Installing to /usr/local/bin (you may need to enter your password)..."
  sudo mv "$BIN_NAME" /usr/local/bin/
fi

cd -
rm -rf "$TMPDIR"

echo "✅ Easypass $LATEST_TAG installed!"
echo "ℹ️ Run 'easypass version' to verify."
