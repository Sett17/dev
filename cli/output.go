package cli

import (
	"github.com/Sett17/dev/global"
	"github.com/i582/cfmt/cmd/cfmt"
	"os"
	"strings"
)

func printPrefix() {
	cfmt.Print("{{DEV:}}::lightBlue ")
}

func Dbg(msg string) {
	if global.Dbg {
		printPrefix()
		cfmt.Printf("{{%s}}::gray\n", msg)
	}
}

func Info(msg string) {
	printPrefix()
	cfmt.Printf("{{%s}}::green\n", msg)
}

func Warning(msg string) {
	printPrefix()
	cfmt.Printf("{{%s}}::bgYellow|black\n", msg)
}

func Error(err error, exit bool) {
	cfmt.Printf("{{ERROR:}}::bgRed\n{{%s}}::red\n", err.Error())
	if exit {
		os.Exit(1)
	}
}

func Separator() {
	cfmt.Printf("{{%s}}::gray\n", strings.Repeat("─", 80))
}
