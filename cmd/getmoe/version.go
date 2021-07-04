package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var (
	version = "master"
	commit  = "none"
	date    = "unknown"
)

var versionCommand = cli.Command{
	Name:   "version",
	Usage:  "Print the version number of Getmoe",
	Action: versionAction,
}

func versionAction(ctx *cli.Context) error {
	fmt.Printf("getmoe v%v Commit=%v BuildDate=%v\n",
		version,
		commit,
		date,
	)
	return nil
}
