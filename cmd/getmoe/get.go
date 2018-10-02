package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/leonidboykov/getmoe"
	"github.com/leonidboykov/getmoe/internal/helper"
	"github.com/leonidboykov/getmoe/provider"
)

var getCommand = cli.Command{
	Name:   "get",
	Usage:  "get data from booru",
	Action: getAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "from",
			Usage: "source booru (config file is not implemented yet)",
		},
		cli.StringFlag{
			Name:  "to",
			Usage: "destination folder",
		},
		cli.StringSliceFlag{
			Name:  "tags",
			Usage: "Tags",
		},
		cli.StringFlag{
			Name:  "login, l",
			Usage: "login for booru",
		},
		cli.StringFlag{
			Name:  "password, p",
			Usage: "password for booru",
		},
	},
}

func getAction(ctx *cli.Context) error {
	srcFlag := ctx.String("from")
	dstFlag := ctx.String("to")
	tagFlag := ctx.StringSlice("tags")
	loginFlag := ctx.String("login")
	passwordFlag := ctx.String("password")
	quietFlag := ctx.GlobalBool("quiet")

	board, ok := provider.AvailableBoards[srcFlag]
	if !ok {
		fmt.Printf("There is no %s source specified\n", srcFlag)
		os.Exit(1)
	}

	board.Provider.Auth(getmoe.AuthConfiguration{
		Login:    loginFlag,
		Password: passwordFlag,
	})

	board.Provider.BuildRequest(getmoe.RequestConfiguration{
		Tags: tagFlag,
	})

	posts, err := board.RequestAll()
	if err != nil {
		fmt.Println(err)
	}

	return saveImage(posts, dstFlag, quietFlag)
}

func saveImage(posts []getmoe.Post, saveDir string, quiet bool) error {
	for _, p := range posts {
		os.MkdirAll(saveDir, os.ModePerm)
		if err := p.Save(saveDir); err != nil {
			return err
		}

		if !quiet {
			fName, _ := utils.FileURLUnescape(p.FileURL)
			fmt.Println("Getting", fName[:32], "...")
		}
	}
	return nil
}
