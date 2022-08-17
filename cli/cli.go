package cli

import (
	"github.com/Sett17/dev/global"
	"github.com/alexeyco/simpletable"
	"github.com/i582/cfmt/cmd/cfmt"
	"strings"
)

func Parse(args []string) {
	for _, arg := range args {
		if !strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--") {
			split := strings.Split(arg, "+")
			op, err := global.FindOperation(split[0])
			if err != nil {
				Error(err, true)
			}
			op.SuppliedArgs = split[1:]
			global.CurrentOps = append(global.CurrentOps, op)
		} else {
			arg = strings.TrimLeft(arg, "-")
			for _, pArg := range ProgArgs {
				if pArg.Short == arg || pArg.Long == arg {
					defer pArg.Func()
				}
			}
		}
	}
}

func ParseForDebug(args []string) {
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") || strings.HasPrefix(arg, "--") {
			arg = strings.TrimLeft(arg, "-")
			if arg == DebugProgArg.Short || arg == DebugProgArg.Long {
				DebugProgArg.Func()
				return
			}
		}
	}
}

func ParseForHelp(args []string) {
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") || strings.HasPrefix(arg, "--") {
			arg = strings.TrimLeft(arg, "-")
			if arg == HelpProgArg.Short || arg == HelpProgArg.Long {
				HelpProgArg.Func()
				return
			}
		}
	}
}

func ParseForVersion(args []string) {
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") || strings.HasPrefix(arg, "--") {
			arg = strings.TrimLeft(arg, "-")
			if arg == VersionProgArg.Short || arg == VersionProgArg.Long {
				VersionProgArg.Func()
				return
			}
		}
	}
}

var roundTableStyle = &simpletable.Style{
	Border: &simpletable.BorderStyle{
		TopLeft:            "╭",
		Top:                "─",
		TopRight:           "╮",
		Right:              "│",
		BottomRight:        "╯",
		Bottom:             "─",
		BottomLeft:         "╰",
		Left:               "│",
		TopIntersection:    "┬",
		BottomIntersection: "┴",
	},
	Divider: &simpletable.DividerStyle{
		Left:         "├",
		Center:       "─",
		Right:        "┤",
		Intersection: "┼",
	},
	Cell: "│",
}

var Logo = cfmt.Sprint(`
   ________  _______   ___      ___ ________ ___  ___       _______
  |\   ___ \|\  ___ \ |\  \    /  /|\  _____\\  \|\  \     |\  ___ \
  \ \  \_|\ \ \   __/|\ \  \  /  / | \  \__/\ \  \ \  \    \ \   __/|
   \ \  \ \\ \ \  \_|/_\ \  \/  / / \ \   __\\ \  \ \  \    \ \  \_|/__
    \ \  \_\\ \ \  \_|\ \ \    / /   \ \  \_| \ \  \ \  \____\ \  \_|\ \
     \ \_______\ \_______\ \__/ /     \ \__\   \ \__\ \_______\ \_______\
      \|_______|\|_______|\|__|/       \|__|    \|__|\|_______|\|_______|
        by Sett   {{https://github.com/Sett17/Devfile}}::green
`)
