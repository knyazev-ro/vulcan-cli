package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func HandleCreateModule(args []string) {
	if len(args) < 1 {
		ErrorPrintln("Error: missing module name")
		return
	}

	module := args[0]
	module, err := ValidateName(module)
	if err != nil {
		ErrorPrintln("Erro: invalid module name:", err.Error())
		return
	}

	isForce := Contains(args, func(e string) bool {
		return e == "--force"
	}) >= 0

	if _, err := os.Stat(module); !os.IsNotExist(err) && !isForce {
		println("Module already exists.")
		return
	}

	settings := LoadSettings()
	if settings == nil {
		ErrorPrintln("Error loading settings")
		return
	}

	templates := settings.Templates
	directories := settings.GeneratedModuleFileStructure

	println("Initializing module:", module)
	src := filepath.Join(module, directories.Src)
	tests := filepath.Join(module, directories.Tests)
	scripts := filepath.Join(module, directories.Scripts)
	docker := filepath.Join(module, directories.Docker)
	configs := filepath.Join(module, directories.Configs)
	docs := filepath.Join(module, directories.Docs)
	config_utils := filepath.Join(module, directories.ConfigUtils)
	githubWorkflows := filepath.Join(module, directories.GithubWorkflows)
	migrations := filepath.Join(module, directories.Migrations)

	middlewares := filepath.Join(module, directories.Middlewares)
	controllers := filepath.Join(module, directories.Controllers)
	routes := filepath.Join(module, directories.Routes)
	interfaces := filepath.Join(module, directories.Interfaces)
	models := filepath.Join(module, directories.Models)
	services := filepath.Join(module, directories.Services)
	repositories := filepath.Join(module, directories.Repositories)
	utils := filepath.Join(module, directories.Utils)
	enums := filepath.Join(module, directories.Enums)

	os.MkdirAll(src, 0755)
	os.MkdirAll(middlewares, 0755)
	os.MkdirAll(controllers, 0755)
	os.MkdirAll(routes, 0755)
	os.MkdirAll(docs, 0755)

	os.MkdirAll(configs, 0755)
	os.MkdirAll(config_utils, 0755)
	os.MkdirAll(githubWorkflows, 0755)

	os.MkdirAll(migrations, 0755)

	os.MkdirAll(interfaces, 0755)
	os.MkdirAll(models, 0755)
	os.MkdirAll(services, 0755)
	os.MkdirAll(repositories, 0755)
	os.MkdirAll(utils, 0755)
	os.MkdirAll(tests, 0755)
	os.MkdirAll(scripts, 0755)
	os.MkdirAll(docker, 0755)
	os.MkdirAll(enums, 0755)

	// You can also create a main.go or other initial files here
	mainFile := module + "/main.go"
	mainTmplFile := templates.Module
	path, err := ParseTemplate(mainTmplFile, mainFile, map[string]string{"Module": module}, args)
	if err != nil {
		ErrorPrintln("Error creating "+path+":", err.Error())
		return
	}

	appFile := module + "/app.go"
	appTmplFile := templates.App
	path, err = ParseTemplate(appTmplFile, appFile, map[string]string{
		"Route":      strings.Join(strings.Split(routes, "\\"), "/"),
		"FileName":   module,
		"Module":     module,
		"Migrations": strings.Join(strings.Split(migrations, "\\"), "/"),
		"Configs":    strings.Join(strings.Split(configs, "\\"), "/"),
	}, args)
	if err != nil {
		ErrorPrintln("Error creating "+path+":", err.Error())
		return
	}

	// You can also create a githubWorkflows or other initial files here
	githubWorkflowFile := githubWorkflows + "/workflow.yaml"
	githubWorkflowTmplFile := templates.GithubWorkflows
	path, err = ParseTemplate(githubWorkflowTmplFile, githubWorkflowFile, map[string]string{"Module": module}, args)
	if err != nil {
		ErrorPrintln("Error creating "+path+":", err.Error())
		return
	}

	// You can also create a README.md or other initial files here
	readmeFile := module + "/README.md"
	readmeTmplFile := templates.Readme
	path, err = ParseTemplate(readmeTmplFile, readmeFile, map[string]string{"Module": module}, args)
	if err != nil {
		ErrorPrintln("Error creating "+path+":", err.Error())
		return
	}

	//create routes.go based on template route.tmpl in vulcan/templates
	routesFile := routes + "/routes.go"
	routesTmplFile := templates.Route
	path, err = ParseTemplate(routesTmplFile, routesFile, map[string]string{"Module": module}, args)
	if err != nil {
		ErrorPrintln("Error creating "+path+":", err.Error())
		return
	}

	migrationsFile := migrations + "/migrations.go"
	migrationsTmplFile := templates.Migrations
	path, err = ParseTemplate(migrationsTmplFile, migrationsFile, map[string]string{"Module": module}, args)
	if err != nil {
		ErrorPrintln("Error creating "+path+":", err.Error())
		return
	}

	//run go mod init and tidy
	cmd := exec.Command("go", "mod", "init", module)
	cmd.Dir = module
	err = cmd.Run()

	if err != nil {
		ErrorPrintln("Error while running Go Mod Go module:", err.Error(), ". Skip this step...")
	}

	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = module
	err = cmd.Run()

	if err != nil {
		ErrorPrintln("Error running go mod tidy:", err.Error())
		return
	}

	// create a .gitignore file from the template in vulcan/templates
	gitignoreFile := module + "/.gitignore"
	gitignoreTmplFile := templates.GitIgnore
	path, err = ParseTemplate(gitignoreTmplFile, gitignoreFile, map[string]string{"Module": module}, args)
	if err != nil {
		ErrorPrintln("Error creating "+path+":", err.Error())
		return
	}

	//create a Dockerfile from the template in vulcan/templates
	dockerfile := docker + "/Dockerfile"
	dockerTmplFile := templates.Dockerfile
	path, err = ParseTemplate(dockerTmplFile, dockerfile, map[string]string{"Module": module}, args)
	if err != nil {
		ErrorPrintln("Error creating "+path+":", err.Error())
		return
	}

	// create env example file
	envExampleFile := module + "/.env.example"
	envExampleTmplFile := templates.EnvExample
	path, err = ParseTemplate(envExampleTmplFile, envExampleFile, map[string]string{"Module": module}, args)
	if err != nil {
		ErrorPrintln("Error creating "+path+":", err.Error())
		return
	}
	// create a config file from the template in vulcan/templates
	configFile := configs + "/config.go"
	configTmplFile := templates.ConfigBase
	path, err = ParseTemplate(configTmplFile, configFile, map[string]string{"Module": module}, args)
	if err != nil {
		ErrorPrintln("Error creating "+path+":", err.Error())
		return
	}

	configUtilsFile := config_utils + "/config_utils.go"
	configUtilsTmplFile := templates.ConfigUtils
	path, err = ParseTemplate(configUtilsTmplFile, configUtilsFile, map[string]string{"Module": module}, args)
	if err != nil {
		ErrorPrintln("Error creating "+path+":", err.Error())
		return
	}

	configDatabaseFile := configs + "/database.go"
	configDatabaseTmplFile := templates.ConfigDatabase
	path, err = ParseTemplate(configDatabaseTmplFile, configDatabaseFile, map[string]string{"Module": module}, args)
	if err != nil {
		ErrorPrintln("Error creating "+path+":", err.Error())
		return
	}

	cmd = exec.Command("git", "init")
	cmd.Dir = module
	err = cmd.Run()

	if err != nil {
		ErrorPrintln("Error initializing git repository:", err.Error())
		return
	}
	// Here you would typically create the module directory structure
	// and possibly generate some initial files or configurations.
	// For now, we just print the module name.
	SuccessPrintln("Module initialized successfully:", module)
	println("You can now start adding controllers, middlewares, and routes to your module.")

}
