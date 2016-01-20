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

func (a *AppEnv) GetEnvs(cli plugin.CliConnection, args []string) (string, error) {
	if loggedIn, _ := cli.IsLoggedIn(); loggedIn == false {
		return "", errors.New("You must login first!")
	}

	if len(args) <= 1 {
		return "", errors.New("You must specify an app name")
	}

	cliOut, err := a.GetAppEnvFromCli(cli, args[1])

	vcapApp, errJson := a.GetJson("VCAP_APPLICATION", cliOut)
	if errJson == nil {
		fmt.Println(formatExportEnvvar("VCAP_APPLICATION", vcapApp))
	}

	vcapServices, errJson := a.GetJson("VCAP_SERVICES", cliOut)
	if errJson == nil {
		fmt.Println(formatExportEnvvar("VCAP_SERVICES", vcapServices))
	}

	return "", err
}

func (a *AppEnv) Run(cli plugin.CliConnection, args []string) {
	if len(args) > 0 && args[0] == "CLI-MESSAGE-UNINSTALL" {
		return
	}

	_, err := a.GetEnvs(cli, args)

	panicOnError(err)
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
				HelpText: "hello",
			},
		},
	}
}
