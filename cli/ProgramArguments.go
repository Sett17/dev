package cli

import (
	"dev/global"
	"github.com/i582/cfmt/cmd/cfmt"
	"os"
)

type progArgFunction func()

type ProgArg struct {
	Name  string
	Short string
	Long  string
	Help  string

	Func progArgFunction
}

var DebugProgArg = ProgArg{
	Name:  "Debug",
	Short: "d",
	Long:  "debug",
	Help:  "Enabled Debug output",
	Func: func() {
		global.Dbg = true
		Dbg("Debug output enabled")
	},
}

var HelpProgArg = ProgArg{
	Name:  "Help",
	Short: "h",
	Long:  "help",

	Func: help,
}

var VersionProgArg = ProgArg{
	Name:  "Version",
	Short: "v",
	Long:  "version",
	Help:  "Show the version of this program",
	Func: func() {
		cfmt.Printf("Devfile version %s\n", global.Version)
		os.Exit(0)
	},
}

var ProgArgs = []ProgArg{
	{
		Name:  "List",
		Short: "l",
		Long:  "list",
		Help:  "List all passed operations",
		Func:  list,
	},
	{
		Name:  "Info",
		Short: "i",
		Long:  "info",
		Help:  "Shows information about any operations, as well as the script",
		Func:  info,
	},
}
