package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/leonidboykov/getmoe/board"
	"github.com/leonidboykov/getmoe/conf"
)

var (
	version = "master"
	commit  = "none"
	date    = "unknown"
)

func main() {
	config, err := conf.Load("getmoe.yaml")
	if err != nil {
		fmt.Println(err)
	}

	for k, p := range config.Boards {
		fmt.Println("Key      :", k)
		fmt.Println("Name     :", p.Name)
		fmt.Println("Provider :", p.Provider)
		fmt.Println("Login    :", p.Auth.Login)
		fmt.Println("Password :", p.Auth.Password)
		fmt.Println("URL     :", p.URL)
		fmt.Println()

		b, err := board.New(p)
		if err != nil {
			fmt.Println(err)
		}

		b.Provider.Auth(p.Auth, b.URL)
		fmt.Println(b.URL.String())
	}
}

func tempMain() {
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
