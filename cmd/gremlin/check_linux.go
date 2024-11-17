package main

import (
	"os"
	"os/exec"
)

func checkLinux() {
	cmd := exec.Command("/bin/bash")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	cmd.Process.Kill()
}
