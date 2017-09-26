package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/leonidboykov/getmoe/board/moebooru"
)

func main() {
	board := moebooru.YandeReConfig
	board.BuildAuth("login", "password")

	board.Query = board.Query{Tags: []string{"nipples", "rating:e"}}

	posts, err := board.Request(tags)
	if err != nil {
		fmt.Println(err)
	}

	for _, p := range posts {
		fmt.Println(p.FileURL)
	}
}

type command interface {
	Name() string
	Args() string
	ShortHelp() string
	LongHelp() string
	Register(*flag.FlagSet)
	Hidden() bool
	Run(*Ctx, []string) error
}

// Ctx ...
type Ctx struct {
	Verbose bool
}

func cli() {
	// Based on go dep command interface
	commands := []command{
		&versionCommand{},
	}

	usage := func() {
		fmt.Println("getmoe is a tool for accessing to images boards (boorus)")
		fmt.Println()
		fmt.Println("Usage: getmoe <command>")
		fmt.Println()
		fmt.Println("Commands:")
		fmt.Println()
		w := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
		for _, cmd := range commands {
			if !cmd.Hidden() {
				fmt.Fprintf(w, "\t%s\t%s\n", cmd.Name(), cmd.ShortHelp())
			}
		}
		w.Flush()
		fmt.Println()
		fmt.Println("Use \"getmoe help [command]\" for more information about a command.")
	}

	cmdName, printCommandHelp, exit := parseArgs(os.Args)
	if exit {
		usage()
	}

	for _, cmd := range commands {
		if cmd.Name() == cmdName {
			// Build flag set with global flags in there.
			fs := flag.NewFlagSet(cmdName, flag.ContinueOnError)
			fs.SetOutput(os.Stdout)
			verbose := fs.Bool("v", false, "enable verbose logging")

			// Register the subcommand flags in there, too.
			cmd.Register(fs)

			// Override the usage text to something nicer.
			resetUsage(fs, cmdName, cmd.Args(), cmd.LongHelp())

			if printCommandHelp {
				fs.Usage()
				return
			}

			// Parse the flags the user gave us.
			// flag package automatically prints usage and error message in err != nil
			// or if '-h' flag provided
			if err := fs.Parse(os.Args[2:]); err != nil {
				return
			}

			ctx := &Ctx{
				Verbose: *verbose,
			}

			// Run the command with the post-flag-processing args.
			if err := cmd.Run(ctx, fs.Args()); err != nil {
				fmt.Printf("%v\n", err)
				return
			}

			// Easy peasy livin' breezy.
			return
		}
	}

	fmt.Printf("getmoe: %s: no such command\n", cmdName)
	usage()
	return
}

func resetUsage(fs *flag.FlagSet, name, args, longHelp string) {
	var (
		hasFlags   bool
		flagBlock  bytes.Buffer
		flagWriter = tabwriter.NewWriter(&flagBlock, 0, 4, 2, ' ', 0)
	)
	fs.VisitAll(func(f *flag.Flag) {
		hasFlags = true
		// Default-empty string vars should read "(default: <none>)"
		// rather than the comparatively ugly "(default: )".
		defValue := f.DefValue
		if defValue == "" {
			defValue = "<none>"
		}
		fmt.Fprintf(flagWriter, "\t-%s\t%s (default: %s)\n", f.Name, f.Usage, defValue)
	})
	flagWriter.Flush()
	fs.Usage = func() {
		fmt.Printf("Usage: getmoe %s %s\n", name, args)
		fmt.Println()
		fmt.Println(strings.TrimSpace(longHelp))
		fmt.Println()
		if hasFlags {
			fmt.Println("Flags:")
			fmt.Println()
			fmt.Println(flagBlock.String())
		}
	}
}

// parseArgs determines the name of the dep command and whether the user asked for
// help to be printed.
func parseArgs(args []string) (cmdName string, printCmdUsage bool, exit bool) {
	isHelpArg := func() bool {
		return strings.Contains(strings.ToLower(args[1]), "help") || strings.ToLower(args[1]) == "-h"
	}

	switch len(args) {
	case 0, 1:
		exit = true
	case 2:
		if isHelpArg() {
			exit = true
		}
		cmdName = args[1]
	default:
		if isHelpArg() {
			cmdName = args[2]
			printCmdUsage = true
		} else {
			cmdName = args[1]
		}
	}
	return cmdName, printCmdUsage, exit
}
