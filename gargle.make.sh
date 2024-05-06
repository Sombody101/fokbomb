#!/bin/bash

echo "go check..."
if ! hash "$(which go)"; then
    echo "Go is not installed!"
fi

echo "garble check..."
if ! hash ~/go/bin/garble; then
    echo "Garble is not installed!"
fi

cd "./src" || {
    echo "Failed to cd into ./src"
    exit 1
}

echo "Building for Windowsx64"
GOOS=windows GOARCH=amd64 ~/go/bin/garble build \
    -ldflags "-w -s -X main.__DEBUG_str=false" \
    -o ../build/fokbomb_garbled.exe || {
        ret="$?"
        echo "Build failed"
        exit "$ret"
    }

echo "Built to ./build/fokbomb_garbled.exe"
cd ..
