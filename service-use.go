package main

import (
	"encoding/json"
	"fmt"
	"github.com/cloudfoundry/cli/plugin"
	"strings"
	"time"
)

/*
*	This is the struct implementing the interface defined by the core CLI. It can
*	be found at  "github.com/cloudfoundry/cli/plugin/plugin.go"
*
 */
type ServiceUsePlugin struct{}

/*
*	This function must be implemented by any plugin because it is part of the
*	plugin interface defined by the core CLI.
*
*	Run(....) is the entry point when the core CLI is invoking a command defined
*	by a plugin. The first parameter, plugin.CliConnection, is a struct that can
*	be used to invoke cli commands. The second paramter, args, is a slice of
*	strings. args[0] will be the name of the command, and will be followed by
*	any additional arguments a cli user typed in.
*
*	Any error handling should be handled with the plugin itself (this means printing
*	user facing errors). The CLI will exit 0 if the plugin exits 0 and will exit
*	1 should the plugin exits nonzero.
 */
func (c *ServiceUsePlugin) Run(cliConnection plugin.CliConnection, args []string) {
	// Ensure that we called the command basic-plugin-command
	if args[0] == "service-use" {
		fmt.Println("---Getting service instances")
		var stringResults []string
		var error error
		stringResults, error = cliConnection.CliCommandWithoutTerminalOutput("curl", "/v2/service_instances")
		if error != nil {
			fmt.Println("Error getting service instances: ", error)
		} else {
			result := []byte(strings.Join(stringResults, ""))
			pagedResponse := &PagedResponse{}
			json.Unmarshal(result, &pagedResponse)
			for pagedResponse != nil {
				fmt.Println("Looking at ", len(pagedResponse.Resources), " services")
				for _, serviceInstance := range pagedResponse.Resources {
					fmt.Println(serviceInstance.Entity.Name)
				}
				if pagedResponse.NextUrl != "" {
					stringResults, error = cliConnection.CliCommandWithoutTerminalOutput("curl", pagedResponse.NextUrl)
					if error != nil {
						fmt.Println("Error getting service instances: ", error)
						pagedResponse = nil
					} else {
						pagedResponse = &PagedResponse{}
						result = []byte(strings.Join(stringResults, ""))
						json.Unmarshal(result, &pagedResponse)
					}
				} else {
					pagedResponse = nil
				}
			}
		}
	}
}

/*
*	This function must be implemented as part of the	plugin interface
*	defined by the core CLI.
*
*	GetMetadata() returns a PluginMetadata struct. The first field, Name,
*	determines the name of the plugin which should generally be without spaces.
*	If there are spaces in the name a user will need to properly quote the name
*	during uninstall otherwise the name will be treated as seperate arguments.
*	The second value is a slice of Command structs. Our slice only contains one
*	Command Struct, but could contain any number of them. The first field Name
*	defines the command `cf basic-plugin-command` once installed into the CLI. The
*	second field, HelpText, is used by the core CLI to display help information
*	to the user in the core commands `cf help`, `cf`, or `cf -h`.
 */
func (c *ServiceUsePlugin) GetMetadata() plugin.PluginMetadata {
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

/*
* Unlike most Go programs, the `Main()` function will not be used to run all of the
* commands provided in your plugin. Main will be used to initialize the plugin
* process, as well as any dependencies you might require for your
* plugin.
 */
func main() {
	// Any initialization for your plugin can be handled here
	//
	// Note: to run the plugin.Start method, we pass in a pointer to the struct
	// implementing the interface defined at "github.com/cloudfoundry/cli/plugin/plugin.go"
	//
	// Note: The plugin's main() method is invoked at install time to collect
	// metadata. The plugin will exit 0 and the Run([]string) method will not be
	// invoked.
	plugin.Start(new(ServiceUsePlugin))
	// Plugin code should be written in the Run([]string) method,
	// ensuring the plugin environment is bootstrapped.
}

type PagedResponse struct {
	TotalResults int              `json:"total_results"`
	TotalPages   int              `json:"total_pages"`
	PrevUrl      string           `json:"prev_url"`
	NextUrl      string           `json:"next_url"`
	Resources    []MetaDataEntity `json:"resources"`
}

type MetaDataEntity struct {
	MetaData MetaData `json:"metadata"`
	Entity   Entity   `json:"entity"`
}

type MetaData struct {
	GUID      string `json:"guid"`
	URL       string `json:"url"`
	CreatedAt Date   `json:"created_at"`
	UpdatedAt Date   `json:"updated_at"`
}

type Entity struct {
	Name               string
	Credentials        interface{}
	ServicePlanGUID    string      `json:"service_plan_guid"`
	SpaceGUID          string      `json:"space_guid"`
	GatewayData        interface{} `json:"gateway_data"`
	DashboardURL       string      `json:"dashboard_url"`
	Type               string
	LastOperation      LastOperation `json:"last_operation"`
	SpaceURL           string        `json:"space_url"`
	ServicePlanURL     string        `json:"service_plan_url"`
	ServiceBindingsURL string        `json:"service_bindings_url"`
}

type LastOperation struct {
	Type        string
	State       string
	Description string
	UpdatedAt   Date `json:"updated_at"`
}

type Date struct{ time.Time }

func (d *Date) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("date should be a string, got %s", data)
	}
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return fmt.Errorf("invalid date: %v", err)
	}
	d.Time = t
	return nil
}
