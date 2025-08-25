# gowasi

A simple WebAssembly (WASI) example using Go and Wazero runtime.

## Requirements

- Go 1.25.0 or later

## Build

Build the WebAssembly module:

```bash
GOOS=wasip1 GOARCH=wasm go build -o wasi/main.wasm wasi/main.go
```

## Run

Execute the WASI module with Wazero:

```bash
go run main.go
```

