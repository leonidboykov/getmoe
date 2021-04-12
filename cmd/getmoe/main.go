package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "getmoe",
		Usage:   `cli tool for boorus`,
		Version: version,
		Commands: []*cli.Command{
			&getCommand,
		},
		Flags:  rootFlags,
		Action: rootAction,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
