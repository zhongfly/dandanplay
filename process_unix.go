//go:build linux || darwin

package main

import (
	"os"
	"os/exec"
	"strings"
)

// 在 macOS 上获取父进程名称
func getParentProcessName() (string, error) {
	ppid := os.Getppid()
	cmd := exec.Command("ps", "-p", string(ppid), "-o", "comm=")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}
