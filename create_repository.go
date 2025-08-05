package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func HandleCreateRepository(args []string) {
	settings := LoadSettings()
	if settings == nil {
		println("Error loading settings")
		return
	}

	commandsActivity := settings.Commands

	if !commandsActivity.CreateRepository {
		fmt.Println("Repository creation is disabled in settings.")
		return
	}

	templates := settings.Templates
	directories := settings.GeneratedModuleFileStructure
	tmplFile := templates.Repository
	module := args[1]
	name := args[0]

	nameCamalCase, _ := CreateStructNameAndVar(name)
	data := map[string]string{
		"Name": nameCamalCase,
	}

	target := filepath.Join(module, directories.Repositories)
	os.MkdirAll(target, 0755)
	outFile := fmt.Sprintf("%s/%s.go", target, name)
	path, err := ParseTemplate(tmplFile, outFile, data)
	if err != nil {
		println("Error creating "+path+":", err.Error())
		return
	}

	fmt.Printf("Created repository: %s\n", outFile)
}
