#!/bin/bash

###
# I could not get this to work inside ./makefile because
# it would throw an error saying that the project directory had
# to be the first argument.
#
# This file should only be run either by make, or if you are
# in the same directory as the script.
###

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

echo "Building for Linuxx64"
GOOS=linux GOARCH=amd64 ~/go/bin/garble build \
    -ldflags "-w -s -X main.__DEBUG_str=false" \
    -o ../build/fokbomb_garbled || {
        ret="$?"
        echo "Build failed"
        exit "$ret"
    }
echo "Built to ./build/fokbomb_garbled"

cd ..
