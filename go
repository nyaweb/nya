#!/bin/bash

murl="https://go.dev/dl/"

version=$(curl -s https://go.dev/dl/ | grep -oP '(go[0-9]+\.[0-9]+(\.[0-9]+)?)\.src\.tar\.gz' | sed 's/\.src\.tar\.gz//' | uniq | head -n 1)

if [ "$(dpkg --print-architecture)" = "arm64" ]; then
  ark="linux-arm64"
elif [ "$(dpkg --print-architecture)" = "amd64" ]; then
  ark="linux-amd64"
elif [ "$(dpkg --print-architecture)" = "i386" ]; then
  ark="linux-386"
elif [ "$(dpkg --print-architecture)" = "armhf" ]; then
  ark="linux-armv6l"
else
  ark="arquitectura-no-soportada"
fi

extension=".tar.gz"

URL="$murl$version.$ark$extension"

curl -OL "$URL"

tar -C /usr/local -xvzf go*.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
go version