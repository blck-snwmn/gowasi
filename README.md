# gowasi

A simple WebAssembly (WASI) example using Go and Wazero runtime.

## Requirements

- Go 1.25.0 or later
- aqua

## Development

Install development tools with aqua:

```bash
aqua i
```

Install lefthook:

```bash
lefthook install
```

lefthook runs `golangci-lint run --enable gosec`, a WASI build check, and `go test -v ./...` before commit.

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
