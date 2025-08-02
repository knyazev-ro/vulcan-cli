package main

import (
	"flag"
	"os"
	"os/exec"
)

func HandleInit(args []string) {

	templatesFolder := "gerard/templates"
	flagSet := flag.NewFlagSet("args", flag.ContinueOnError)
	flagSet.Parse(args[2:])
	allArgs := flagSet.Args()

	if len(allArgs) < 1 {
		println("Error: missing module name")
		return
	}

	module := allArgs[0]
	if module == "" {
		println("Error: module name cannot be empty")
		return
	}
	println("Initializing module:", module)
	src := module + "/src"
	tests := module + "/tests"
	scripts := module + "/scripts"
	docker := module + "/docker"

	middlewares := src + "/middlewares"
	controllers := src + "/controllers"
	routes := src + "/routes"
	configs := src + "/configs"
	interfaces := src + "/interfaces"
	models := src + "/models"
	services := src + "/services"
	repositories := src + "/repositories"
	utils := src + "/utils"
	enums := src + "/enums"

	os.MkdirAll(src, 0755)
	os.MkdirAll(middlewares, 0755)
	os.MkdirAll(controllers, 0755)
	os.MkdirAll(routes, 0755)
	os.MkdirAll(configs, 0755)
	os.MkdirAll(interfaces, 0755)
	os.MkdirAll(models, 0755)
	os.MkdirAll(services, 0755)
	os.MkdirAll(repositories, 0755)
	os.MkdirAll(utils, 0755)
	os.MkdirAll(tests, 0755)
	os.MkdirAll(scripts, 0755)
	os.MkdirAll(docker, 0755)
	os.MkdirAll(enums, 0755)

	// You can also create a README.md or other initial files here
	readmeFile := module + "/README.md"
	readmeTmplFile := templatesFolder + "/readme.tmpl"
	path, err := ParseTemplate(readmeTmplFile, readmeFile, map[string]string{"Module": module})
	if err != nil {
		println("Error creating "+path+":", err.Error())
		return
	}

	//create routes.go based on template route.tmpl in gerard/templates
	routesFile := routes + "/routes.go"
	routesTmplFile := templatesFolder + "/route.tmpl"
	path, err = ParseTemplate(routesTmplFile, routesFile, map[string]string{"Module": module})
	if err != nil {
		println("Error creating "+path+":", err.Error())
		return
	}

	//run go mod init and tidy
	cmd := exec.Command("go", "mod", "init", module)
	cmd.Dir = module
	err = cmd.Run()

	if err != nil {
		println("Error initializing Go module:", err.Error())
		return
	}

	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = module
	err = cmd.Run()

	if err != nil {
		println("Error running go mod tidy:", err.Error())
		return
	}

	// create a .gitignore file from the template in gerard/templates
	gitignoreFile := module + "/.gitignore"
	gitignoreTmplFile := templatesFolder + "/gitignore.tmpl"
	path, err = ParseTemplate(gitignoreTmplFile, gitignoreFile, map[string]string{"Module": module})
	if err != nil {
		println("Error creating "+path+":", err.Error())
		return
	}

	//create a Dockerfile from the template in gerard/templates
	dockerfile := docker + "/Dockerfile"
	dockerTmplFile := templatesFolder + "/dockerfile.tmpl"
	path, err = ParseTemplate(dockerTmplFile, dockerfile, map[string]string{"Module": module})
	if err != nil {
		println("Error creating "+path+":", err.Error())
		return
	}

	// create env example file
	envExampleFile := module + "/.env.example"
	envExampleTmplFile := templatesFolder + "/env-example.tmpl"
	path, err = ParseTemplate(envExampleTmplFile, envExampleFile, map[string]string{"Module": module})
	if err != nil {
		println("Error creating "+path+":", err.Error())
		return
	}
	// create a config file from the template in gerard/templates
	configFile := configs + "/config.go"
	configTmplFile := templatesFolder + "/config.tmpl"
	path, err = ParseTemplate(configTmplFile, configFile, map[string]string{"Module": module})
	if err != nil {
		println("Error creating "+path+":", err.Error())
		return
	}

	cmd = exec.Command("git", "init")
	cmd.Dir = module
	err = cmd.Run()

	if err != nil {
		println("Error initializing git repository:", err.Error())
		return
	}
	// Here you would typically create the module directory structure
	// and possibly generate some initial files or configurations.
	// For now, we just print the module name.
	println("Module initialized successfully:", module)
	println("You can now start adding controllers, middlewares, and routes to your module.")

}
