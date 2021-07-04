package main

import (
	"github.com/urfave/cli/v2"

	"github.com/leonidboykov/getmoe"
	_ "github.com/leonidboykov/getmoe/provider/danbooru"
	_ "github.com/leonidboykov/getmoe/provider/gelbooru"
	_ "github.com/leonidboykov/getmoe/provider/moebooru"
	_ "github.com/leonidboykov/getmoe/provider/sankaku"
	_ "github.com/leonidboykov/getmoe/provider/sankaku/v2"
)

var rootFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    "config",
		Aliases: []string{"c"},
		Usage:   "override config file name",
		Value:   "getmoe.yaml",
	},
}

func rootAction(ctx *cli.Context) error {
	configFlag := ctx.String("config")
	config, err := getmoe.ReadConfiguraton(configFlag)
	if err != nil {
		return err
	}

	boards, err := getmoe.LoadBoards(config.Boards)
	if err != nil {
		return err
	}
	if err := boards.ExecuteCommands(config.Download); err != nil {
		return err
	}

	return nil
}
