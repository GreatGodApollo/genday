package cmd

import (
	"github.com/jawher/mow.cli"
	"os"
)

var verbose *bool

func Exec() {
	app := cli.App("genday", "Generate a curday.dat file")

	app.Spec = "[-v]"

	verbose = app.BoolOpt("v verbose", false, "Verbose debug mode")

	app.Command("json", "generate a file from JSON", cmdJSON)

	app.Run(os.Args)
}