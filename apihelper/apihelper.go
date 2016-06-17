package apihelper

import (
	"strconv"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/krujos/cfcurl"
)

//Service representation
type Service struct {
	URL             string
	Label           string
	ServicePlansURL string
}

//ServicePlan representation
type ServicePlan struct {
	URL                 string
	Name                string
	ServiceInstancesURL string
}

//CFAPIHelper to wrap cf curl results
type CFAPIHelper interface {
	GetServices(plugin.CliConnection) ([]Service, error)
	GetServicePlans(cli plugin.CliConnection, plansURL string) ([]ServicePlan, error)
}

//APIHelper implementation
type APIHelper struct{}

//Function type to simplify processing paged results
type process func(metadata map[string]interface{}, entity map[string]interface{}) interface{}

//Base method to process paged results from API calls
func processPagedResults(cli plugin.CliConnection, url string, fn process) (interface{}, error) {

	theJSON, err := cfcurl.Curl(cli, url)
	if nil != err {
		return nil, err
	}

	pages := int(theJSON["total_pages"].(float64))
	var objects []interface{}
	for i := 1; i <= pages; i++ {
		if 1 != i {
			theJSON, err = cfcurl.Curl(cli, url+"?page="+strconv.Itoa(i))
		}
		for _, o := range theJSON["resources"].([]interface{}) {
			theObj := o.(map[string]interface{})
			entity := theObj["entity"].(map[string]interface{})
			metadata := theObj["metadata"].(map[string]interface{})
			objects = append(objects, fn(metadata, entity))
		}

	}
	return objects, nil
}

//GetServices returns a struct that represents critical fields in the JSON
func (api *APIHelper) GetServices(cli plugin.CliConnection) ([]Service, error) {
	services, err := processPagedResults(cli, "/v2/services", func(metadata map[string]interface{}, entity map[string]interface{}) interface{} {
		return Service{
			Label:           entity["label"].(string),
			URL:             metadata["url"].(string),
			ServicePlansURL: entity["service_plans_url"].(string),
		}
	})

	if nil != err {
		return nil, err
	}

	return services.([]Service), nil
}

//GetServices returns a struct that represents critical fields in the JSON
func (api *APIHelper) GetServicePlans(cli plugin.CliConnection, plansURL string) ([]ServicePlan, error) {
	serviceplans, err := processPagedResults(cli, plansURL, func(metadata map[string]interface{}, entity map[string]interface{}) interface{} {
		return ServicePlan{
			Name:                entity["name"].(string),
			URL:                 metadata["url"].(string),
			ServiceInstancesURL: entity["service_instances_url"].(string),
		}
	})
	if nil != err {
		return nil, err
	}

	return serviceplans.([]ServicePlan), nil
}
