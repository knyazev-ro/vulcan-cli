package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 1 {
		fmt.Println("Usage: create:{middleware|controller} {Name} {Module}")
		return
	}

	command := os.Args[1]
	switch command {
	case "create:middleware":
		HandleCreateMiddleware(os.Args)
	case "create:controller":
		HandleCreateController(os.Args)
	case "create:repository":
		HandleCreateRepository(os.Args)
	case "create:model":
		HandleCreateModel(os.Args)
	case "create:interface":
		HandleCreateInterface(os.Args)
	case "create:service":
		HandleCreateService(os.Args)
	case "create:config":
		HandleCreateConfig(os.Args)
	case "create:init":
		HandleInit(os.Args)
	case "remove:module":
		HandleRemoveModule(os.Args)
	default:
		fmt.Println("Unknown command:", command)
		return
	}
	fmt.Println("gerard is executed.")
}
