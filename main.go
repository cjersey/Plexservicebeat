package main

import (
	"os"

	"github.com/cjersey/plexservicebeat/cmd"

	_ "github.com/cjersey/plexservicebeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
