package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func HandleCreateController(args []string) {

	tmplFile := "gerard/templates/controller.tmpl"
	target := "/src/controllers"

	var module string
	flagSet := flag.NewFlagSet("args", flag.ContinueOnError)
	flagSet.Parse(args[2:])

	names := flagSet.Args()
	module = names[1]
	name := names[0]

	nameCamalCase, _ := CreateStructNameAndVar(name)
	data := map[string]string{
		"Name": nameCamalCase,
	}

	os.MkdirAll(module+target, 0755)
	outFile := fmt.Sprintf("%s%s/%s.go", module, target, strings.ToLower(name))
	path, err := ParseTemplate(tmplFile, outFile, data)
	if err != nil {
		println("Error creating "+path+":", err.Error())
		return
	}

	fmt.Printf("Created controller: %s\n", outFile)
}
