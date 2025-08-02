package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func HandleCreateMiddleware(args []string) {
	tmplFile := "gerard/templates/middleware.tmpl"
	block := "/src/middlewares"

	var module string
	flagSet := flag.NewFlagSet("args", flag.ContinueOnError)
	flagSet.Parse(args[2:])
	allArgs := flagSet.Args()

	if len(allArgs) < 2 {
		fmt.Println("Error: missing fields")
		return
	}
	module = allArgs[1]
	name := allArgs[0]

	nameCamalCase, nameVar := CreateStructNameAndVar(name)
	data := map[string]string{
		"Name":    nameCamalCase,
		"NameVar": nameVar,
	}

	// Создаём директорию если нет
	os.MkdirAll(module+block, 0755)
	outFile := fmt.Sprintf("%s%s/%s.go", module, block, strings.ToLower(name))
	path, err := ParseTemplate(tmplFile, outFile, data)
	if err != nil {
		println("Error creating "+path+":", err.Error())
		return
	}

	fmt.Printf("Created middleware: %s\n", outFile)
}
