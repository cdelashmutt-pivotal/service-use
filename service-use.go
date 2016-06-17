package main

import (
	"fmt"
	"github.com/cdelashmutt-pivotal/service-use/apihelper"
	"github.com/cloudfoundry/cli/plugin"
	"os"
)

type ServiceUsePlugin struct {
	apiHelper apihelper.CFAPIHelper
	cli       plugin.CliConnection
}

type service struct {
	label string
	plans []serviceplan
}

type serviceplan struct {
	name string
}

func (cmd *ServiceUsePlugin) ServiceUseCommand(args []string) {
	fmt.Println("---Getting service instances")

	if nil == cmd.cli {
		fmt.Println("ERROR: CLI Connection is nil!")
		os.Exit(1)
	}

	services, _ := cmd.getServices()

	for _, service := range services {
		fmt.Printf("Service %s was found.\n", service.label)
	}
}

func (cmd *ServiceUsePlugin) getServices() ([]service, error) {
	rawServices, err := cmd.apiHelper.GetServices(cmd.cli)
	if nil != err {
		return nil, err
	}

	var services = []service{}

	for _, s := range rawServices {

		serviceplans, err := cmd.getServicePlans(s.ServicePlansURL)
		if nil != err {
			return nil, err
		}

		services = append(services, service{
			label: s.Label,
			plans: serviceplans,
		})
	}
	return services, nil
}

func (cmd *ServiceUsePlugin) getServicePlans(servicePlansURL string) ([]serviceplan, error) {
	rawServicePlans, err := cmd.apiHelper.GetServicePlans(cmd.cli, servicePlansURL)
	if nil != err {
		return nil, err
	}

	var serviceplans = []serviceplan{}

	for _, sp := range rawServicePlans {

		serviceplans = append(serviceplans, serviceplan{
			name: sp.Name,
		})
	}
	return serviceplans, nil
}

func (cmd *ServiceUsePlugin) Run(cli plugin.CliConnection, args []string) {

	if args[0] == "service-use" {
		cmd.apiHelper = &apihelper.APIHelper{}
		cmd.cli = cli
		cmd.ServiceUseCommand(args)
	}
}

func (cmd *ServiceUsePlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "ServiceUsePlugin",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		Commands: []plugin.Command{
			plugin.Command{
				Name:     "service-use",
				HelpText: "Infomation about service instances and bound apps",

				// UsageDetails is optional
				// It is used to show help of usage of each command
				UsageDetails: plugin.Usage{
					Usage: "service-use\n   cf service-use",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(ServiceUsePlugin))
}
