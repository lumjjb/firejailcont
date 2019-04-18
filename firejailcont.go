package main

import (
	"github.com/nabla-containers/runnc/llcli"
	ll "github.com/nabla-containers/runnc/llif"
	llfs "github.com/nabla-containers/runnc/llmodules/fs"
	llnet "github.com/nabla-containers/runnc/llmodules/network"
)

func main() {
	fsH, err := llfs.NewNoopFsHandler()
	if err != nil {
		panic(err)
	}
	networkH, err := llnet.NewNoopNetworkHandler()
	if err != nil {
		panic(err)
	}

	weirdVagrantLLCHandler := ll.RunllcHandler{
		FsH:      fsH,
		NetworkH: networkH,
	}

	// We run the OCI runtime called "runnc", with root dir "/run/runnc"
	// with the low level handlers chosen above.
	llcli.Runllc("firejailcont", "/run/firejailcont", weirdVagrantLLCHandler)
}
