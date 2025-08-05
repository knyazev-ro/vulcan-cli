package main

import (
	"os"
)

func HandleRemoveModule(args []string) {

	module := args[0]

	if module == "" {
		println("Module name is required")
		return
	}
	if module == "gerard-cli" {
		println("Cannot remove the 'gerard-cli'. Are you stupid?")
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
