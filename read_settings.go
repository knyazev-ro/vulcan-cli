package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

type GeneratedModuleFileStructure struct {
	Repositories    string `yaml:"repositories"`
	Controllers     string `yaml:"controllers"`
	Middlewares     string `yaml:"middlewares"`
	Models          string `yaml:"models"`
	Services        string `yaml:"services"`
	Interfaces      string `yaml:"interfaces"`
	Routes          string `yaml:"routes"`
	Utils           string `yaml:"utils"`
	Enums           string `yaml:"enums"`
	Configs         string `yaml:"configs"`
	Tests           string `yaml:"tests"`
	Scripts         string `yaml:"scripts"`
	Docker          string `yaml:"docker"`
	Docs            string `yaml:"docs"`
	ConfigUtils     string `yaml:"config_utils"`
	Src             string `yaml:"src"`
	GithubWorkflows string `yaml:"github-workflows"`
}

type Template struct {
	Repository      string `yaml:"repository"`
	Controller      string `yaml:"controller"`
	Service         string `yaml:"service"`
	Interface       string `yaml:"interface"`
	Model           string `yaml:"model"`
	Middleware      string `yaml:"middleware"`
	Route           string `yaml:"route"`
	Enum            string `yaml:"enum"`
	Dockerfile      string `yaml:"dockerfile"`
	EnvExample      string `yaml:"env-example"`
	Module          string `yaml:"module"`
	Readme          string `yaml:"readme"`
	GitIgnore       string `yaml:"gitignore"`
	Config          string `yaml:"config"`
	ConfigBase      string `yaml:"config-base"`
	ConfigDatabase  string `yaml:"config-database"`
	ConfigServer    string `yaml:"config-server"`
	ConfigUtils     string `yaml:"config-utils"`
	GithubWorkflows string `yaml:"github-workflows"`
}

type Command struct {
	CreateInit       bool `yaml:"create-init"`
	CreateModel      bool `yaml:"create-model"`
	CreateConfig     bool `yaml:"create-config"`
	CreateService    bool `yaml:"create-service"`
	CreateInterface  bool `yaml:"create-interface"`
	CreateController bool `yaml:"create-controller"`
	CreateMiddleware bool `yaml:"create-middleware"`
	CreateRepository bool `yaml:"create-repository"`
	RemoveModule     bool `yaml:"remove-module"`
}

type Settings struct {
	GeneratedModuleFileStructure GeneratedModuleFileStructure `yaml:"generated-file-structure"`
	Templates                    Template                     `yaml:"templates"`
	Commands                     Command                      `yaml:"commands"`
}

func DefaultSettings() *Settings {
	settings := Settings{
		GeneratedModuleFileStructure: GeneratedModuleFileStructure{
			Repositories:    "src/repositories",
			Controllers:     "src/controllers",
			Middlewares:     "src/middlewares",
			Models:          "src/models",
			Services:        "src/services",
			Interfaces:      "src/interfaces",
			Routes:          "src/routes",
			Utils:           "src/utils",
			Enums:           "src/enums",
			Configs:         "configs",
			Tests:           "tests",
			Scripts:         "scripts",
			Docker:          "docker",
			Docs:            "docs",
			ConfigUtils:     "configs/utils",
			Src:             "src",
			GithubWorkflows: ".github/workflows",
		},
		Templates: Template{
			Repository:      "gerard-cli/templates/repository.tmpl",
			Controller:      "gerard-cli/templates/controller.tmpl",
			Service:         "gerard-cli/templates/service.tmpl",
			Interface:       "gerard-cli/templates/interface.tmpl",
			Model:           "gerard-cli/templates/model.tmpl",
			Middleware:      "gerard-cli/templates/middleware.tmpl",
			Route:           "gerard-cli/templates/route.tmpl",
			Enum:            "gerard-cli/templates/enum.tmpl",
			Dockerfile:      "gerard-cli/templates/dockerfile.tmpl",
			EnvExample:      "gerard-cli/templates/env-example.tmpl",
			GitIgnore:       "gerard-cli/templates/gitignore.tmpl",
			ConfigBase:      "gerard-cli/templates/config_base.tmpl",
			ConfigDatabase:  "gerard-cli/templates/config_database.tmpl",
			ConfigServer:    "gerard-cli/templates/config_server.tmpl",
			ConfigUtils:     "gerard-cli/templates/config_utils.tmpl",
			GithubWorkflows: "gerard-cli/templates/github-workflows",
		},

		Commands: Command{
			CreateInit:       true,
			CreateModel:      true,
			CreateConfig:     true,
			CreateService:    true,
			CreateInterface:  true,
			CreateController: true,
			CreateMiddleware: true,
			CreateRepository: true,
			RemoveModule:     true,
		},
	}
	return &settings
}

func LoadSettings() *Settings {

	settings, err := os.ReadFile("gerard-cli/settings.yaml")

	if err != nil {
		println("Error reading settings.yaml:", err.Error())
		println("Using default settings")
		return DefaultSettings()
	}

	var settingsConfig Settings
	err = yaml.Unmarshal(settings, &settingsConfig)
	if err != nil {
		println("Error unmarshalling settings.yaml:", err.Error())
		println("Using default settings")
		return DefaultSettings()
	}
	return &settingsConfig
}
