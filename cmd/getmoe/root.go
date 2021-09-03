package main

import (
	"github.com/urfave/cli/v2"

	"github.com/leonidboykov/getmoe"
	"github.com/leonidboykov/getmoe/cmd/getmoe/downloader"
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

	d, err := downloader.NewDownloader(config.Boards)
	if err != nil {
		return err
	}

	if err := d.Execute(config.Download); err != nil {
		return err
	}

	return nil
}
