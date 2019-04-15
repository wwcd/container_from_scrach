package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		parent()
	case "child":
		child()
	default:
		panic("wat should I do")
	}
}

func parent() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET,
		// Unshareflags: syscall.CLONE_NEWNS,
	}

	must(cmd.Start())
	fmt.Println("Container PID is ", cmd.Process.Pid)
	must(cmd.Wait())
}

func child() {
	fmt.Println("Container init PID is ", os.Getpid())

	must(syscall.Sethostname([]byte("simplecontainer")))
	must(syscall.Chroot("./rootfs/mountedfs"))
	must(os.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Env = []string{"PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin"}

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
