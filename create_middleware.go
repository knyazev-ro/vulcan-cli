package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func HandleCreateMiddleware(args []string) {
	settings := LoadSettings()
	if settings == nil {
		println("Error loading settings")
		return
	}
	commandsActivity := settings.Commands

	if !commandsActivity.CreateMiddleware {
		fmt.Println("Middleware creation is disabled in settings.")
		return
	}

	templates := settings.Templates
	directories := settings.GeneratedModuleFileStructure
	tmplFile := templates.Middleware

	if len(args) < 2 {
		fmt.Println("Error: missing fields")
		return
	}
	module := args[1]
	name := args[0]

	nameCamalCase, nameVar := CreateStructNameAndVar(name)
	data := map[string]string{
		"Name":    nameCamalCase,
		"NameVar": nameVar,
	}

	target := filepath.Join(module, directories.Middlewares)
	os.MkdirAll(target, 0755)
	outFile := fmt.Sprintf("%s/%s.go", target, strings.ToLower(name))
	path, err := ParseTemplate(tmplFile, outFile, data)
	if err != nil {
		println("Error creating "+path+":", err.Error())
		return
	}

	fmt.Printf("Created middleware: %s\n", outFile)
}
