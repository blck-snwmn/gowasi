package main

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

func TestDo(t *testing.T) {
	ctx := context.Background()
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)

	stdout := bytes.NewBuffer(nil)
	stderr := bytes.NewBuffer(nil)
	cfg := wazero.NewModuleConfig().WithStdout(stdout).WithStderr(stderr)

	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	_, err := r.InstantiateWithConfig(ctx, wasm, cfg.WithArgs("wasi", "-r=20", "-l=15"))
	if err != nil {
		t.Fatal(err)
	}
	stdoutstr := stdout.String()
	fmt.Println(stdoutstr)
	if stdoutstr != "35\n" {
		t.Fatalf("unexpected output: %q", stdoutstr)
	}
	stderrstr := stderr.String()
	if stderrstr != "" {
		t.Fatalf("unexpected error output: %q", stderrstr)
	}
}
