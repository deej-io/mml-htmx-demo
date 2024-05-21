#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
VENDOR_DIR="$SCRIPT_DIR/../../vendor"

echo updating submodules...
git submodule update --init --remote

echo patching submodules...
echo -n > "$VENDOR_DIR/mml/check-node-version.js"

echo building vendored packages
( cd "$VENDOR_DIR/mml/packages/observable-dom" && npm install && npm run build && npm link )

echo linking vendored packages
( cd "$SCRIPT_DIR/.." && npm link @mml-io/observable-dom )
