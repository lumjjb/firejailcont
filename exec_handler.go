package main

import (
	ll "github.com/nabla-containers/runnc/llif"
	"os"
	"path/filepath"
	"syscall"
)

type firejailExecHandler struct{}

func (h *firejailExecHandler) ExecCreateFunc(i *ll.ExecCreateInput) (*ll.LLState, error) {
	ret := &ll.LLState{}
	return ret, nil
}

func (h *firejailExecHandler) ExecRunFunc(i *ll.ExecRunInput) error {
	cfg := i.Config
	bin := filepath.Join(cfg.Rootfs, cfg.Args[0])
	argv := []string{bin}
	argv = append(argv, cfg.Args[1:]...)
	syscall.Exec("firejail", argv, os.Environ())
	return nil
}

func (h *firejailExecHandler) ExecDestroyFunc(i *ll.ExecDestroyInput) (*ll.LLState, error) {
	ret := &ll.LLState{}
	return ret, nil
}
