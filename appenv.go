package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/cloudfoundry/cli/plugin"
)

type AppEnv struct{}

func main() {
	plugin.Start(new(AppEnv))
}

func (a *AppEnv) GetAppEnvFromCli(cli plugin.CliConnection, appName string) ([]string, error) {
	return cli.CliCommandWithoutTerminalOutput("env", appName)
}

func panicOnError(err error) {
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
}

func (a *AppEnv) GetJson(key string, out []string) (string, error) {
	for _, line := range out {
		if strings.Contains(line, key) {

			bytes := []byte(line)

			var j interface{}
			err := json.Unmarshal(bytes, &j)

			if err != nil {
				return "", err
			}

			m := j.(map[string]interface{})

			marshalledBytes, err := json.Marshal(m[key])

			if err != nil {
				return "", err
			}

			return string(marshalledBytes[:]), nil
		}
	}

	return "", errors.New(fmt.Sprintf("%s not found", key))
}

func formatExportEnvvar(key string, value string) string {
	return fmt.Sprintf("export %s='%s'", key, value)
}

func (a *AppEnv) GetJsonAndFormat(key string, cliOut []string) string {
	json, err := a.GetJson(key, cliOut)
	if err == nil {
		return formatExportEnvvar(key, json)
	}
	return ""
}

func (a *AppEnv) GetEnvs(cli plugin.CliConnection, args []string) ([]string, error) {
	if loggedIn, _ := cli.IsLoggedIn(); loggedIn == false {
		return nil, errors.New("You must login first!")
	}

	if len(args) <= 1 {
		return nil, errors.New("You must specify an app name")
	}

	cliOut, err := a.GetAppEnvFromCli(cli, args[1])

	vcapServicesExport := a.GetJsonAndFormat("VCAP_SERVICES", cliOut)
	vcapAppExport := a.GetJsonAndFormat("VCAP_APPLICATION", cliOut)

	envvars := []string{}

	if vcapServicesExport != "" {
		envvars = append(envvars, vcapServicesExport)
	}

	if vcapAppExport != "" {
		envvars = append(envvars, vcapAppExport)
	}

	return envvars, err
}

func (a *AppEnv) Run(cli plugin.CliConnection, args []string) {
	if len(args) > 0 && args[0] == "CLI-MESSAGE-UNINSTALL" {
		return
	}

	envVars, err := a.GetEnvs(cli, args)

	panicOnError(err)

	for _, env := range envVars {
		fmt.Println(env)
	}
}

func (a *AppEnv) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "appenvs",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 1,
			Build: 0,
		},
		Commands: []plugin.Command{
			plugin.Command{
				Name:     "appenvs",
				HelpText: "Export application environment variables locally.",
				UsageDetails: plugin.Usage{
					Usage: "appenvs APP_NAME",
				},
			},
		},
	}
}
