package main

import (
	"errors"
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
)

type AppEnv struct{}

func main() {
	plugin.Start(new(AppEnv))
}

func (a *AppEnv) GetEnvs(cli plugin.CliConnection) (string, error) {
	if loggedIn, _ := cli.IsLoggedIn(); loggedIn == false {
		return "", errors.New("oops")
	}
	return "", nil
}

func (a *AppEnv) Run(cli plugin.CliConnection, args []string) {
	if len(args) > 0 && args[0] == "CLI-MESSAGE-UNINSTALL" {
		return
	}

	a.GetEnvs(cli)

	fmt.Println("test")
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
