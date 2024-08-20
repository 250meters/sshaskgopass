package gopass

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// GopassPath represents the absolute path to the `gopass` executable.
type GopassPath string

// New returns a new instance of GopassPath with the provided path as default. If no path is provided,
// it uses the value from the $GOPASS environment variable if set, otherwise defaults to "gopass".
func New(gopassPath string) GopassPath {
	if gopassPath == "" {
		gopassPath = os.Getenv("GOPASS")
		if gopassPath == "" {
			gopassPath = "gopass"
		}
	}
	return GopassPath(gopassPath)
}

// Otp runs the `gopass otp` command with the provided key and if successful it will return the OTP code.
func (g GopassPath) Otp(ctx context.Context, key string) (code string, err error) {
	return g.cmd(ctx, "otp", key)
}

// Passphrase runs the `gopass passphrase` command with the provided key and if successful it will return the requested passphrase
func (g GopassPath) Passphrase(ctx context.Context, key string) (code string, err error) {
	return g.cmd(ctx, "passphrase", key)
}

// Password runs the `gopass password` command with the provided key and if successfull it will return the requested passphrase
func (g GopassPath) Password(ctx context.Context, key string) (code string, err error) {
	return g.cmd(ctx, "password", key)
}

func (gp GopassPath) cmd(ctx context.Context, kind, key string) (code string, err error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	sgCtx, sgCancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer sgCancel()

	arg := "show"
	if kind == "otp" {
		arg = "otp"
	}

	if kind != "passphrase" {
		key = "/" + key
	}

	cmd := exec.CommandContext(sgCtx, string(gp), arg, "-o", "ssh/"+kind+key) // #nosec G204
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("gopass could not get %s: %w, %s", kind, err, strings.TrimSpace(errb.String()))
	}

	return outb.String(), nil
}
