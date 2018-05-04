package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var (
	version = "master"
	commit  = "none"
	date    = "unknown"
)

func main() {
	app := cli.NewApp()
	app.Name = "getmoe"
	app.Usage = "cli tool for boorus"
	app.Version = version
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Leonid Boykov",
			Email: "leonid.v.boykov@gmail.com",
		},
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "quiet, q",
			Usage: "disable progress bar",
		},
	}
	app.Commands = []cli.Command{
		getCommand,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
