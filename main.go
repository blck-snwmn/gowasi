package main

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "embed"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	"github.com/tetratelabs/wazero/sys"
)

//go:embed wasi/main.wasm
var wasm []byte

func main() {
	ctx := context.Background()
	r := wazero.NewRuntime(ctx)
	defer func() {
		if err := r.Close(ctx); err != nil {
			log.Printf("failed to close runtime: %v", err)
		}
	}()

	cfg := wazero.NewModuleConfig().WithStdout(os.Stdout).WithStderr(os.Stderr)

	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	_, err := r.InstantiateWithConfig(ctx, wasm, cfg.WithArgs("wasi", "-r=11", "-l=33"))
	if err != nil {
		if exitErr, ok := err.(*sys.ExitError); ok && exitErr.ExitCode() != 0 {
			fmt.Fprintf(os.Stderr, "exit_code: %d\n", exitErr.ExitCode())
		} else {
			log.Panicln(err)
		}
	}
}
