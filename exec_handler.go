package main

import (
	"fmt"
	"path/filepath"
	"syscall"

	ll "github.com/nabla-containers/runnc/llif"
)

type firejailExecHandler struct{}

func (h *firejailExecHandler) ExecCreateFunc(i *ll.ExecCreateInput) (*ll.LLState, error) {
	ret := &ll.LLState{}
	return ret, nil
}

func (h *firejailExecHandler) ExecRunFunc(i *ll.ExecRunInput) error {
	cfg := i.Config
	bin := filepath.Join(cfg.Rootfs, cfg.Args[0])
	argv := []string{"firejail", bin}
	argv = append(argv, cfg.Args[1:]...)
	env := []string{"PATH=/usr/sbin:/usr/bin:/sbin:/bin"}
	fmt.Printf("Executing following command: /usr/bin/firejail %v\n with env: \n%v\n", argv, env)
	return syscall.Exec("/usr/bin/firejail", argv, env)
}

func (h *firejailExecHandler) ExecDestroyFunc(i *ll.ExecDestroyInput) (*ll.LLState, error) {
	ret := &ll.LLState{}
	return ret, nil
}
