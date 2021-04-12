package main

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/leonidboykov/getmoe"
	_ "github.com/leonidboykov/getmoe/provider/danbooru"
	_ "github.com/leonidboykov/getmoe/provider/gelbooru"
	_ "github.com/leonidboykov/getmoe/provider/moebooru"
	_ "github.com/leonidboykov/getmoe/provider/sankaku"
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

	for boardName, board := range config.Boards {
		fmt.Println(boardName, board.Provider.Name)
	}

	boards, err := getmoe.LoadBoards(config.Boards)
	_ = boards

	return nil
}
