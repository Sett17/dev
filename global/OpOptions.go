package global

import (
	"fmt"
	"strings"
)

type opOption struct {
	Name  string
	Short string
	Long  string
	Help  string
}

func (o opOption) String() string {
	return fmt.Sprintf("%s", o.Name)
}

var PRINT = opOption{"PRINT", "p", "print", "Print the script before executing"}
var TIME = opOption{"TIME", "t", "time", "Print the time taken to execute the script"}
var QUIET = opOption{"QUIET", "q", "quiet", "Don't print any STDOUT output"}

//var WSL = opOption{"WINDOWS", "w", "wsl", "Use Windows to execute the script, only applicable with WSL"}

var OpsOptions = []opOption{PRINT, TIME, QUIET}

func ParseOption(op *Operation, text string) error {
	args := op.Arguments
	opts := op.Options
	if strings.HasPrefix(text, "+") {
		args = append(args, strings.ToUpper(strings.TrimPrefix(text, "+")))
	} else {
		found := false
		for _, opt := range OpsOptions {
			if strings.ToLower(text) == opt.Short {
				opts = append(opts, opt)
				found = true
			}
		}
		if !found && text != "" {
			return fmt.Errorf("Unknown option %s in %s", text, op.Name)
		}
	}
	op.Arguments = args
	op.Options = opts
	return nil
}
