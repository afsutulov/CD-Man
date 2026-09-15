#!/bin/sh
set -eu
cd "$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)"

if [ "$(uname -s)" != Linux ]; then
    echo 'Run this script on Linux.' >&2
    exit 1
fi
if ! command -v go >/dev/null 2>&1; then
    echo 'Go 1.22 or newer is required.' >&2
    exit 1
fi

export CGO_ENABLED=1
mkdir -p build
printf '%s\n' 'Running tests...'
go test ./...
printf '%s\n' 'Building CD-Man...'
go build -buildvcs=false -trimpath -o build/cdman ./cmd/cdman
printf '%s\n' 'Build complete: build/cdman' 'SDL2 must be installed to run the game.'
