package main

import (
	"fmt"
	"github.com/Sett17/dev/cli"
	"github.com/Sett17/dev/df"
	"github.com/Sett17/dev/global"
	"github.com/i582/cfmt/cmd/cfmt"
	"os"
)

func main() {
	cfmt.RegisterStyle("purple", func(s string) string {
		return cfmt.Sprintf("{{%s}}::#BA8EF7", s)
	})

	cli.ParseForHelp(os.Args[1:])
	cli.ParseForVersion(os.Args[1:])
	cli.ParseForDebug(os.Args[1:])
	ops, err := df.Read("devfile")
	if err != nil {
		cli.Error(fmt.Errorf("error reading devfile"), false)
	}
	global.Ops = ops

	cli.Parse(os.Args[1:])

	if len(global.CurrentOps) == 0 {
		cli.HelpProgArg.Func()
	}

	for _, op := range global.CurrentOps {
		err := df.Execute(op)
		if err != nil {
			cli.Error(fmt.Errorf("%s returned with %s", op.Name, err), true)
		}
	}
}
