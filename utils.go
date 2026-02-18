package main

import (
	"os"
	"regexp"
	"strings"
	"text/template"
	"vulcan/templates"

	"github.com/fatih/camelcase"
	"github.com/fatih/color"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Contains[T any](arr []T, needle func(T) bool) int {
	for i := range arr {
		if needle(arr[i]) {
			return i
		}
	}
	return -1
}

func Filter[T any](ss []T, callback func(T) bool) (ret []T) {
	for _, s := range ss {
		if callback(s) {
			ret = append(ret, s)
		}
	}
	return
}

func ParseTemplate(templatePath string, outputFilePath string, data interface{}, args []string) (string, error) {

	isForce := Contains(args, func(x string) bool {
		return x == "--force"
	}) >= 0

	withTests := Contains(args, func(x string) bool {
		return x == "--with-tests"
	}) >= 0

	//check if output directory exists, if it exists then nothing
	if _, err := os.Stat(outputFilePath); !os.IsNotExist(err) && !isForce {
		WarningPrintln("Warning: output file already exists:", outputFilePath)
		return outputFilePath, os.ErrExist
	}

	dataMap := data.(map[string]string)
	if withTests {
		argsForTest := []string{dataMap["FileName"], dataMap["Module"]}
		if isForce {
			argsForTest = append(argsForTest, "--force")
		}
		HandleCreateTesting(argsForTest)
	}

	tmpl, err := template.ParseFS(templates.TemplatesFS, templatePath)
	if err != nil {
		ErrorPrintln("Error parsing template:", err.Error())
		return "", err
	}

	out, err := os.Create(outputFilePath)
	if err != nil {
		ErrorPrintln("Error creating:", err.Error())
		return "", err
	}

	defer out.Close()
	err = tmpl.Execute(out, data)
	if err != nil {
		ErrorPrintln("Error executing template:", err.Error())
		return "", err
	}

	return outputFilePath, nil
}

func Normalize(name string) (string, string, string, error) {

	r := strings.NewReplacer(
		"-", " ",
		"_", " ",
		"+", " ",
		"\\", " ",
		"|", " ",
		"/", " ",
		"_test", "",
	)
	name = r.Replace(name)
	nameSplit := strings.Join(strings.Fields(name), " ")
	splitCamelCase := strings.Fields(strings.Join(camelcase.Split(nameSplit), " "))
	name = strings.Join(splitCamelCase, "_")
	name = strings.ToLower(name)

	name, err := ValidateName(name)
	println(name)
	if err != nil {
		ErrorPrintln("Filename is invalid")
		return "", "", "", err
	}

	splitName := strings.Split(name, "_")
	for i, s := range splitName {
		splitName[i] = cases.Title(language.English).String(s)
	}
	data := map[string]string{
		"Name":    strings.Join(splitName, ""),
		"NameVar": strings.ToLower(splitName[0]) + strings.Join(splitName[1:], ""),
	}

	return data["Name"], data["NameVar"], name, nil
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

func ErrorPrintln(a ...any) {
	c := color.New(color.BgRed)
	c.Println(a...)
}

func SuccessPrintln(a ...any) {
	c := color.New(color.BgGreen)
	c.Println(a...)
}

func WarningPrintln(a ...any) {
	c := color.New(color.BgYellow)
	c.Println(a...)
}
