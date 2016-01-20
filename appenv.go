package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
)

type AppEnv struct{}

func main() {
	plugin.Start(new(AppEnv))
}

func (a *AppEnv) GetEnvs(cli plugin.CliConnection, args []string) (string, error) {
	if loggedIn, _ := cli.IsLoggedIn(); loggedIn == false {
		return "", errors.New("You must login first!")
	}

	if len(args) <= 1 {
		return "", errors.New("You must specify an app name")
	}

	_, err := cli.CliCommandWithoutTerminalOutput("env", args[1])

	return "", err
}

func (a *AppEnv) Run(cli plugin.CliConnection, args []string) {
	if len(args) > 0 && args[0] == "CLI-MESSAGE-UNINSTALL" {
		return
	}

	_, err := a.GetEnvs(cli, args)

	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
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
				HelpText: "hello",
			},
		},
	}
}
