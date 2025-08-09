package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func HandleCreateRepository(args []string) {
	settings := LoadSettings()
	if settings == nil {
		ErrorPrintln("Error loading settings")
		return
	}

	commandsActivity := settings.Commands

	if !commandsActivity.CreateRepository {
		WarningPrintln("Repository creation is disabled in settings.")
		return
	}

	templates := settings.Templates
	directories := settings.GeneratedModuleFileStructure
	tmplFile := templates.Repository
	module := args[1]
	name := args[0]

	nameCamalCase, _, name, err := Normalize(name)
	if err != nil {
		ErrorPrintln("Filename is invalid.")
		return
	}
	data := map[string]string{
		"Name": nameCamalCase,
	}

	target := filepath.Join(module, directories.Repositories)
	os.MkdirAll(target, 0755)
	outFile := fmt.Sprintf("%s/%s.go", target, name)
	path, err := ParseTemplate(tmplFile, outFile, data, args)
	if err != nil {
		ErrorPrintln("Error creating "+path+":", err.Error())
		return
	}

	SuccessPrintln("Created repository: ", outFile)
}
