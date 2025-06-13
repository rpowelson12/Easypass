
#!/bin/bash

set -e

OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [[ "$ARCH" == "x86_64" ]]; then
  ARCH="amd64"
elif [[ "$ARCH" == "arm64" || "$ARCH" == "aarch64" ]]; then
  ARCH="arm64"
fi

BINARY="easypass-$OS-$ARCH"
URL="https://github.com/rpowelson12/Easypass/releases/latest/download/$BINARY"

echo "Downloading latest Easypass from $URL..."
curl -L -o easypass "$URL"
chmod +x easypass
sudo mv easypass /usr/local/bin/easypass

echo "Easypass upgraded!"
easypass version
