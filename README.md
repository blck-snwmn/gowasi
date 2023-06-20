Using Go1.21rc1

# Build
```bash
GOOS=wasip1 GOARCH=wasm go build -o wasi/main.wasm wasi/main.go
```

# Run
```bash
go run main.go
```

