package goutil

import (
	"context"
	"os"
	"os/exec"
	"strings"
)

func RunCommandCtx(ctx context.Context, command string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, command, args...)
	out, err := cmd.Output()

	return string(out), err
}

func RunCommandCtxWithOutput(ctx context.Context, command string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, command, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	out, err := cmd.Output()

	return string(out), err
}

func RunCommandContainsCtx(ctx context.Context, contains string, command string, args ...string) bool {
	out, err := RunCommandCtx(ctx, command, args...)
	if err != nil {
		return false
	}

	return strings.Contains(out, contains)
}
