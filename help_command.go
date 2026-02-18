package main

import "fmt"

func GetHelp() {
	info := `
Available commands:

run-test:
    vulcan.exe create:module <module_name>              - Create a new module (example: jobber)
    vulcan.exe create:controller <name> <module>        - Create a controller in a module
    vulcan.exe create:middleware <name> <module>        - Create middleware in a module
    vulcan.exe create:model <name> <module>             - Create a model in a module
    vulcan.exe create:repository <name> <module>        - Create a repository in a module
    vulcan.exe create:service <name> <module>           - Create a service in a module
    vulcan.exe create:interface <name> <module>         - Create an interface in a module
    vulcan.exe create:config <name> <module>            - Create a config file in a module
    vulcan.exe help                                      - Show this help message

destroy:
    vulcan.exe remove:module <module_name>               - Remove the specified module

Usage:
    Replace <module_name> and <name> with your desired names.
`
	fmt.Println(info)
}
