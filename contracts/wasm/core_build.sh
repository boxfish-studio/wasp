#!/bin/bash
go install ../../tools/schema
cd ../../packages/wasmvm/wasmlib
schema -core -go -rust -ts -force
cd ../../../contracts/wasm
rm -rf ./node_modules/wasmlib/
rm -rf ./node_modules/wasmclient/
cp -R ../../packages/wasmvm/wasmlib/ts/wasmlib ./node_modules

# gascalibration
for dir in ./gascalibration/*; do
  if [ -d "$dir" ]; then
    mkdir -p "$dir"/ts/node_modules
    cp -R ../../packages/wasmvm/wasmlib/ts/wasmlib "$dir"/ts/node_modules
  fi
done

