package df

import (
	"fmt"
	"github.com/Sett17/dev/cli"
	"github.com/Sett17/dev/global"
	"github.com/i582/cfmt/cmd/cfmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func Execute(op *global.Operation) error {
	cli.Info(fmt.Sprintf("Executing operation %s with options %v", op.Name, op.Options))

	if len(op.SuppliedArgs) > len(op.Arguments) {
		return fmt.Errorf("too many arguments supplied for operation %s", op.Name)
	}

	for _, opt := range op.Options {
		if opt == global.PRINT {
			cli.Info("Script:")
			for ln, line := range strings.Split(op.Script, "\n") {
				if len(line) != 0 {
					cfmt.Printf("  {{%d}}::purple %s\n", ln, line)
				}
			}
		}
	}

	cli.Separator()

	executor := "/bin/bash"
	extraArg := "-e"
	pattern := "devfile-"
	prefixLines := "shopt -s expand_aliases\nsource ~/.bash_aliases\n"
	suffixLines := ""
	if runtime.GOOS == "windows" {
		executor = "cmd"
		extraArg = "/c"
		pattern = "devfile-*.bat"
		prefixLines = "@echo off\n"
		suffixLines = "\n@echo on\n"
		for i, arg := range op.Arguments {
			if len(op.SuppliedArgs) > i {
				cli.Dbg(fmt.Sprintf("Adding DEV_%s variable", strings.ToUpper(arg)))
				prefixLines += fmt.Sprintf("set DEV_%s=%s\n", strings.ToUpper(arg), op.SuppliedArgs[i])
			}
		}
	} else {
		for i, arg := range op.Arguments {
			if len(op.SuppliedArgs) > i {
				cli.Dbg(fmt.Sprintf("Adding DEV_%s variable", strings.ToUpper(arg)))
				prefixLines += fmt.Sprintf("DEV_%s=%s\n", strings.ToUpper(arg), op.SuppliedArgs[i])
			}
		}
	}

	start := time.Now()

	f, err := os.CreateTemp("", pattern) // in Go version older than 1.17 you can use ioutil.TempFile
	if err != nil {
		return err
	}

	// close and remove the temporary file at the end of the program
	defer f.Close()
	defer os.Remove(f.Name())

	cli.Dbg(f.Name())

	op.Script = prefixLines + op.Script + suffixLines

	if _, err := f.WriteString(op.Script); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	cli.Dbg(fmt.Sprintf("Creating and writing script took: %s", time.Since(start)))

	start = time.Now()
	defer func() { cli.Info(fmt.Sprintf("Execution took: %s", time.Since(start))) }()

	filePath := f.Name()

	cli.Dbg(fmt.Sprintf("Executing %s %s %s", executor, extraArg, filePath))
	cmd := exec.Command(executor, extraArg, filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}
	cli.Separator()

	return nil
}
