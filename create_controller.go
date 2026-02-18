package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func HandleCreateController(args []string) {

	settings := LoadSettings()
	if settings == nil {
		ErrorPrintln("Error loading settings")
		return
	}

	commandsActivity := settings.Commands

	if !commandsActivity.CreateMiddleware {
		WarningPrintln("Controller creation is disabled in settings.")
		return
	}

	templates := settings.Templates
	directories := settings.GeneratedModuleFileStructure
	tmplFile := templates.Controller

	name := args[0]
	module := args[1]

	module, err := ValidateName(module)
	if err != nil {
		ErrorPrintln("Erro: invalid module name:", err.Error())
		return
	}

	nameCamalCase, nameVar, name, err := Normalize(name)
	if err != nil {
		ErrorPrintln("Filename is invalid.")
		return
	}
	data := map[string]string{
		"Name":     nameCamalCase,
		"NameVar":  nameVar,
		"FileName": name,
		"Module":   module,
	}

	target := filepath.Join(module, directories.Controllers)
	os.MkdirAll(target, 0755)
	outFile := fmt.Sprintf("%s/%s.go", target, strings.ToLower(name))
	path, err := ParseTemplate(tmplFile, outFile, data, args)
	if err != nil {
		ErrorPrintln("Error creating "+path+":", err.Error())
		return
	}

	SuccessPrintln("Created controller: ", outFile)
}
