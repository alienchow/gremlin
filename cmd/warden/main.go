package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	//cmd.SysProcAttr = &syscall.SysProcAttr{
	//	Cloneflags:   syscall.CLONE_NEWNS,
	//	Unshareflags: syscall.CLONE_NEWNS,
	//}

	must(syscall.Chroot("/mnt/jail"))
	must(os.Chdir("/"))
	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
