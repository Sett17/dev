package cli

import (
	"fmt"
	"github.com/Sett17/dev/global"
	"github.com/alexeyco/simpletable"
	"github.com/i582/cfmt/cmd/cfmt"
	"os"
	"strings"
)

func info() {
	for i, op := range global.CurrentOps {
		fmt.Println()
		printPrefix()
		cfmt.Printf("{{Information about %d. operation}}::purple {{%s}}::underline|purple\n", i+1, op.Name)
		cfmt.Printf("  {{Options:}}::underline\n")
		for _, opt := range op.Options {
			cfmt.Printf("    {{%s}}::purple\n", opt.Name)
		}
		if len(op.Options) == 0 {
			cfmt.Println("    {{/}}::purple")
		}
		cfmt.Printf("  {{Arguments:}}::underline\n")
		for _, arg := range op.Arguments {
			cfmt.Printf("    {{%s}}::purple\n", arg)
		}
		if len(op.Arguments) == 0 {
			cfmt.Println("    {{/}}::purple")
		}
		cfmt.Printf("  {{Script:}}::underline\n")
		for ln, line := range strings.Split(op.Script, "\n") {
			if len(line) != 0 {
				cfmt.Printf("    {{%d}}::purple %s\n", ln, line)
			}
		}
	}
	os.Exit(0)
}

func list() {
	table := simpletable.New()
	table.Header = &simpletable.Header{Cells: []*simpletable.Cell{
		{Text: cfmt.Sprint("{{Name}}::purple")},
		{Text: cfmt.Sprint("{{Options}}::purple")},
		{Text: cfmt.Sprint("{{Args}}::purple")},
		{Text: cfmt.Sprint("{{Lines}}::purple")},
		{Text: cfmt.Sprint("{{Description}}::purple")},
	}}
	for _, op := range global.Ops {
		r := []*simpletable.Cell{
			{Text: op.Name},
			{Text: fmt.Sprintf("%v", op.Options)},
			{Text: fmt.Sprintf("%v", op.Arguments)},
			{Text: fmt.Sprintf("%d", strings.Count(op.Script, "\n")-1)},
			{Text: op.Description},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}
	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Text: cfmt.Sprint("{{use 'dev -i OPERATION' to show the underlying script}}::gray"), Span: 5, Align: simpletable.AlignCenter},
	}}
	table.SetStyle(roundTableStyle)

	cfmt.Println("\n{{All parsed operations}}::purple")
	fmt.Println(table.String())
	os.Exit(0)
}

func help() {
	cfmt.Printf(`
{{Usage:}}::underline dev [OPTIONS] OPERATIONS+ARGUMENT..
`)
	fmt.Println(Logo)
	cfmt.Printf(`
Tired of the weird quirks of make? Annoyed of making typos in long chained
commands, or getting to them via reverse-i-search? Well, here is a solution
that comes as just an executable for each platform and with an extensive
help command.

{{Operation Options:}}::underline
 These options can be used as letters in the Devfile

`)
	for _, opt := range global.OpsOptions {
		cfmt.Printf("  {{%s\t%s}}::purple\t{{%s}}::gray\n", opt.Short, opt.Name, opt.Help)
	}

	cfmt.Printf(`

{{Options:}}::underline
`)
	for _, arg := range ProgArgs {
		cfmt.Printf("  {{-%s}}::purple, {{--%s}}::purple\t{{%s}}::gray\n", arg.Short, arg.Long, arg.Help)
	}
	cfmt.Printf("  {{-%s}}::purple, {{--%s}}::purple\t{{%s}}::gray\n", DebugProgArg.Short, DebugProgArg.Long, DebugProgArg.Help)
	cfmt.Printf("  {{-%s}}::purple, {{--%s}}::purple\t{{%s}}::gray\n", "h", "help", "Show this help message")
	cfmt.Printf("  {{-%s}}::purple, {{--%s}}::purple\t{{%s}}::gray\n", VersionProgArg.Short, VersionProgArg.Long, VersionProgArg.Help)

	cfmt.Printf(`

{{Operations:}}::underline
  {{OPERATIONS+ARGUMENT...}}::purple  {{Which operation to execute.}}::gray
                          {{Arguments should be in same order as in Devfile}}::gray
`)
	os.Exit(0)
}
