package main

import (
	"flag"
	"os"
)

func HandleRemoveModule(args []string) {

	var module string
	flagSet := flag.NewFlagSet("args", flag.ContinueOnError)
	flagSet.Parse(args[2:])

	names := flagSet.Args()
	module = names[0]

	if module == "" {
		println("Module name is required")
		return
	}

	err := os.RemoveAll(module)
	if err != nil {
		println("Error removing module:", err.Error())
		return
	}

	println("Module removed successfully:", module)
	os.Exit(0)
}
