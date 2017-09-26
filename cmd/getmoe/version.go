package main

import (
	"flag"
	"fmt"
	"runtime"
)

var (
	version    = "devel"
	buildDate  string
	commitHash string
)

const versionHelp = `Show the getmoe version information`

func (cmd *versionCommand) Name() string { return "version" }
func (cmd *versionCommand) Args() string {
	return ""
}
func (cmd *versionCommand) ShortHelp() string { return versionHelp }
func (cmd *versionCommand) LongHelp() string  { return versionHelp }
func (cmd *versionCommand) Hidden() bool      { return false }

func (cmd *versionCommand) Register(fs *flag.FlagSet) {}

func (cmd *versionCommand) Run(ctx *Ctx, args []string) error {
	fmt.Printf("getmoe v%s %s/%s %s", version, runtime.GOOS, runtime.GOARCH, runtime.Version())
	return nil
}

type versionCommand struct{}
