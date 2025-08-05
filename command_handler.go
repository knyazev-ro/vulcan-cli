package main

import (
	"flag"
	"fmt"
)

func HandleRunCommand(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: create:{middleware|controller} {Name} {Module}")
		return
	}

	flagSet := flag.NewFlagSet("args", flag.ContinueOnError)
	flagSet.Parse(args[2:])
	argsArr := flagSet.Args()

	command := args[1]
	switch command {
	case "create:middleware":
		HandleCreateMiddleware(argsArr)
	case "create:controller":
		HandleCreateController(argsArr)
	case "create:repository":
		HandleCreateRepository(argsArr)
	case "create:model":
		HandleCreateModel(argsArr)
	case "create:interface":
		HandleCreateInterface(argsArr)
	case "create:service":
		HandleCreateService(argsArr)
	case "create:config":
		HandleCreateConfig(argsArr)
	case "create:module":
		HandleCreateModule(argsArr)
	case "remove:module":
		HandleRemoveModule(argsArr)
	default:
		fmt.Println("Unknown command:", command)
		return
	}
}
