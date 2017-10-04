package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/leonidboykov/getmoe"
	"github.com/leonidboykov/getmoe/board"
)

func main() {
	fromFlag := flag.String("from", "", "Source")
	tagsFlag := flag.String("tags", "", "Tags")
	formatFlag := flag.String("format", "image", "Save for {image|json}")
	saveFlag := flag.String("save", "", "Save directory")
	loginFlag := flag.String("login", "", "Login")
	passwordFlag := flag.String("password", "", "Password")
	flag.Parse()

	if *fromFlag == "" || *formatFlag == "" || *saveFlag == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	board, ok := board.AvailableBoards[*fromFlag]
	if !ok {
		fmt.Printf("There are no %s source specified\n", *fromFlag)
		os.Exit(1)
	}

	if *loginFlag != "" && *passwordFlag != "" {
		board.BuildAuth(*loginFlag, *passwordFlag)
	}

	board.Query = getmoe.Query{
		Tags: strings.Split(*tagsFlag, " "),
		Page: 1,
	}

	posts, err := board.RequestAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, p := range posts {
		fmt.Println(p.FileURL)
	}
}
