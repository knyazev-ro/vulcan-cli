package main

import (
	"os"
)

func HandleRemoveModule(args []string) {

	module := args[0]

	if module == "" {
		ErrorPrintln("Module name is required")
		return
	}
	if module == "vulcan-cli" {
		WarningPrintln("Cannot remove the 'vulcan-cli'. Are you stupid?")
		return
	}
	err := os.RemoveAll(module)
	if err != nil {
		ErrorPrintln("Error removing module:", err.Error())
		return
	}

	ErrorPrintln("Module removed successfully:", module)
	os.Exit(0)
}
