package main

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"syscall"
)

const (
	jailDirEnvVar = "JAIL_DIR"
	oldRootDir    = ".old_root"
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
	setChangeGroup()

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Sethostname([]byte("lOsEr")))

	jailDir := os.Getenv(jailDirEnvVar)
	if jailDir == "" {
		jailDir = "/mnt/jail"
	}

	oldRoot := path.Join(jailDir, oldRootDir)
	if err := os.Mkdir(oldRoot, 0700); err != nil {
		if !os.IsExist(err) {
			panic(err)
		}
	}

	must(syscall.PivotRoot(jailDir, oldRoot))
	// must(syscall.Chroot(jailDir))
	must(os.Chdir("/"))
	must(syscall.Unmount("/"+oldRootDir, syscall.MNT_DETACH))
	must(os.Remove("/" + oldRootDir))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	must(cmd.Run())
}

const (
	cgroupRoot   = "/sys/fs/cgroup"
	memLimit     = "500M"
	processLimit = "50"
)

func setChangeGroup() {
	cgroupRoot := "/sys/fs/cgroup"
	gremlin := filepath.Join(cgroupRoot, "gremlin")
	os.Mkdir(gremlin, 0755)
	must(os.WriteFile(filepath.Join(gremlin, "cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
	must(os.WriteFile(filepath.Join(gremlin, "pids.max"), []byte(processLimit), 0700))
	must(os.WriteFile(filepath.Join(gremlin, "memory.max"), []byte(memLimit), 0700))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
