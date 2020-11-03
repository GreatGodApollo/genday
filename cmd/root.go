package cmd

import (
	"github.com/jawher/mow.cli"
	"os"
)

var verbose *bool

func Exec() {
	app := cli.App("genday", "Generate a curday.dat file")

	app.Spec = "[--version] [--verbose]"

	app.Version("v version", "genday " + VERSION)

	verbose = app.BoolOpt("V verbose", false, "Verbose debug mode")

	app.Command("json", "generate a file from JSON", cmdJSON)

	app.Run(os.Args)
}