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
	// pids := filepath.Join(cgroupRoot, "pids")
	// os.Mkdir(pids, 0755)
	// os.Mkdir(filepath.Join(pids, "gremlin"), 0755)
	// must(os.WriteFile(filepath.Join(pids, "gremlin/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
	// must(os.WriteFile(filepath.Join(pids, "gremlin/pids.max"), []byte(processLimit), 0700))
	// must(os.WriteFile(filepath.Join(pids, "gremlin/notify_on_release"), []byte("1"), 0700))
	//
	// mem := filepath.Join(cgroupRoot, "memory")
	// os.Mkdir(mem, 0755)
	// os.Mkdir(filepath.Join(mem, "gremlin"), 0755)
	// must(os.WriteFile(filepath.Join(mem, "gremlin/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
	// must(os.WriteFile(filepath.Join(mem, "gremlin/memory.limit_in_bytes"), []byte(memLimit), 0700))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
