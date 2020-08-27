package main

import (
	"fmt"
	"os"

	"github.com/bamboovir/terradeps/cmd"
	"github.com/bamboovir/terradeps/cmd/store"
	"github.com/bamboovir/terradeps/cmd/types"
	log "github.com/sirupsen/logrus"
)

func main() {
	state := &types.State{
		Logger: log.StandardLogger(),
	}

	rootCMD := cmd.NewRootCMD(os.Args[1:], state, store.NewBasicStore())

	if err := rootCMD.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
