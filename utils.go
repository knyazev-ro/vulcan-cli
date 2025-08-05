package main

import (
	"os"
	"regexp"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ParseTemplate(templatePath string, outputFilePath string, data interface{}) (string, error) {

	//check if output directory exists, if it exists then nothing
	if _, err := os.Stat(outputFilePath); !os.IsNotExist(err) {
		println("Warning: output file already exists:", outputFilePath)
		return outputFilePath, nil
	}

	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		println("Error: template not found:", templatePath)
		return "", err
	}
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		println("Error parsing template:", err.Error())
		return "", err
	}

	out, err := os.Create(outputFilePath)
	if err != nil {
		println("Error creating:", err.Error())
		return "", err
	}

	defer out.Close()
	err = tmpl.Execute(out, data)
	if err != nil {
		println("Error executing template:", err.Error())
		return "", err
	}

	return outputFilePath, nil
}

func CreateStructNameAndVar(name string) (string, string) {
	splitName := strings.Split(name, "_")
	for i, s := range splitName {
		splitName[i] = cases.Title(language.English).String(s)
	}
	data := map[string]string{
		"Name":    strings.Join(splitName, ""),
		"NameVar": strings.ToLower(splitName[0]) + strings.Join(splitName[1:], ""),
	}

	return data["Name"], data["NameVar"]
}

func ValidateName(module string) (string, error) {

	module = strings.TrimSpace(module)
	module = strings.ToLower(module)

	pattern := `^[a-z][a-z0-9]*(?:_[a-z0-9]+)*$`
	re := regexp.MustCompile(pattern)
	if !re.MatchString(module) {
		return "", os.ErrInvalid
	}

	return module, nil
}
