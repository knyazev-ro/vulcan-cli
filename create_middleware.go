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
		ErrorPrintln("Error loading settings")
		return
	}
	commandsActivity := settings.Commands

	if !commandsActivity.CreateMiddleware {
		WarningPrintln("Middleware creation is disabled in settings.")
		return
	}

	templates := settings.Templates
	directories := settings.GeneratedModuleFileStructure
	tmplFile := templates.Middleware

	if len(args) < 2 {
		ErrorPrintln("Error: missing fields")
		return
	}
	module := args[1]
	name := args[0]

	nameCamalCase, nameVar, name, err := Normalize(name)
	if err != nil {
		ErrorPrintln("Filename is invalid.")
		return
	}
	data := map[string]string{
		"Name":    nameCamalCase,
		"NameVar": nameVar,
	}

	target := filepath.Join(module, directories.Middlewares)
	os.MkdirAll(target, 0755)
	outFile := fmt.Sprintf("%s/%s.go", target, strings.ToLower(name))
	path, err := ParseTemplate(tmplFile, outFile, data, args)
	if err != nil {
		ErrorPrintln("Error creating "+path+":", err.Error())
		return
	}

	SuccessPrintln("Created middleware: ", outFile)
}
