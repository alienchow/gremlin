package main

import (
	"os"
	"os/exec"
	"path"
	"syscall"
)

const (
	jailDirEnvVar = "JAIL_DIR"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
	}
}

func run() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWNS | syscall.CLONE_NEWPID,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	must(cmd.Run())
}

func child() {
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Sethostname([]byte("lOsEr")))

	jailDir := os.Getenv(jailDirEnvVar)
	if jailDir == "" {
		jailDir = "/mnt/jail"
	}

	oldRoot := path.Join(jailDir, "old_root")
	if err := os.Mkdir(oldRoot, 0755); err != nil {
		if !os.IsExist(err) {
			panic(err)
		}
	}
	must(syscall.Mount(jailDir, jailDir, "", syscall.MS_BIND|syscall.MS_REC, ""))
	must(syscall.Mount(jailDir, oldRoot, "", syscall.MS_BIND|syscall.MS_REC, ""))
	must(syscall.PivotRoot(jailDir, oldRoot))
	// must(syscall.Chroot(jailDir))
	must(os.Chdir("/"))
	// must(syscall.Unmount(oldRoot, syscall.MNT_DETACH))
	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
