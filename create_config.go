package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func HandleCreateConfig(args []string) {

	settings := LoadSettings()
	if settings == nil {
		ErrorPrintln("Error loading settings")
		return
	}
	commandsActivity := settings.Commands

	if !commandsActivity.CreateMiddleware {
		WarningPrintln("Config creation is disabled in settings.")
		return
	}

	templates := settings.Templates
	directories := settings.GeneratedModuleFileStructure
	tmplFile := templates.Config

	module := args[1]
	name := args[0]

	nameCamalCase, nameVar, name, err := Normalize(name)

	if err != nil {
		ErrorPrintln("Filename is invalid.")
		return
	}

	data := map[string]string{
		"Module":  module,
		"Name":    nameCamalCase,
		"NameVar": nameVar,
	}
	target := filepath.Join(module, directories.Configs)
	os.MkdirAll(target, 0755)
	outFile := fmt.Sprintf("%s/%s.go", target, strings.ToLower(name))
	path, err := ParseTemplate(tmplFile, outFile, data, args)
	if err != nil {
		ErrorPrintln("Error creating "+path+":", err.Error())
		return
	}

	SuccessPrintln("Created config: ", outFile)
}
