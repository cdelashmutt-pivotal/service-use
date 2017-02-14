package apihelper

import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/krujos/cfcurl"
	"strconv"
)

//CFAPIHelper to wrap cf curl results
type CFAPIHelper interface {
	GetServices(cli plugin.CliConnection) ([]Service, error)
	GetServicePlans(cli plugin.CliConnection, plansURL string) ([]ServicePlan, error)
	GetServiceInstances(cli plugin.CliConnection, serviceInstancesURL string) ([]ServiceInstance, error)
	GetSpace(cli plugin.CliConnection, spaceURL string) (Space, error)
	GetOrganization(cli plugin.CliConnection, organizationURL string) (Organization, error)
	GetOrgManagers(cli plugin.CliConnection, orgManagersURL string) ([]OrgManager, error)
}

//APIHelper implementation
type APIHelper struct{}

//Function type to simplify processing paged results
type process func(metadata map[string]interface{}, entity map[string]interface{}) interface{}

//Base method to process paged results from API calls
func processPagedResults(cli plugin.CliConnection, url string, fn process) ([]interface{}, error) {

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

//Service representation
type Service struct {
	URL             string
	Label           string
	ServicePlansURL string
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

	retVal := make([]Service, len(services))
	for i := range services {
		retVal[i] = services[i].(Service)
	}

	if nil != err {
		return nil, err
	}

	return retVal, nil
}

//ServicePlan representation
type ServicePlan struct {
	Name                string
	URL                 string
	ServiceInstancesURL string
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

	retVal := make([]ServicePlan, len(serviceplans))
	for i := range serviceplans {
		retVal[i] = serviceplans[i].(ServicePlan)
	}

	if nil != err {
		return nil, err
	}

	return retVal, nil
}

//ServiceInstance representation
type ServiceInstance struct {
	Name               string
	URL                string
	SpaceURL           string
	ServiceBindingsURL string
	ServiceKeysURL     string
	RoutesURL          string
}

//GetServiceInstances returns a struct that represents critical fields in the JSON
func (api *APIHelper) GetServiceInstances(cli plugin.CliConnection, serviceInstancesURL string) ([]ServiceInstance, error) {
	serviceinstances, err := processPagedResults(cli, serviceInstancesURL, func(metadata map[string]interface{}, entity map[string]interface{}) interface{} {
		return ServiceInstance{
			Name:               entity["name"].(string),
			URL:                metadata["url"].(string),
			SpaceURL:           entity["space_url"].(string),
			ServiceBindingsURL: entity["service_bindings_url"].(string),
			ServiceKeysURL:     entity["service_keys_url"].(string),
			RoutesURL:          entity["routes_url"].(string),
		}
	})

	retVal := make([]ServiceInstance, len(serviceinstances))
	for i := range serviceinstances {
		retVal[i] = serviceinstances[i].(ServiceInstance)
	}

	if nil != err {
		return nil, err
	}

	return retVal, nil
}

type Space struct {
	Name            string
	URL             string
	OrganizationURL string
}

//GetSpace returns a struct that represents critical fields in the JSON
func (api *APIHelper) GetSpace(cli plugin.CliConnection, spaceURL string) (Space, error) {
	theJSON, err := cfcurl.Curl(cli, spaceURL)
	if nil != err {
		return Space{}, err
	}

	entity := theJSON["entity"].(map[string]interface{})
	metadata := theJSON["metadata"].(map[string]interface{})

	space := Space{
		Name:            entity["name"].(string),
		URL:             metadata["url"].(string),
		OrganizationURL: entity["organization_url"].(string),
	}

	return space, nil
}

type Organization struct {
	Name string
	URL  string
        ManagersURL string
}

//GetOrganization returns a struct that represents critical fields in the JSON
func (api *APIHelper) GetOrganization(cli plugin.CliConnection, organizationURL string) (Organization, error) {
	theJSON, err := cfcurl.Curl(cli, organizationURL)
	if nil != err {
		return Organization{}, err
	}

	entity := theJSON["entity"].(map[string]interface{})
	metadata := theJSON["metadata"].(map[string]interface{})

	organization := Organization{
		Name: entity["name"].(string),
		URL:  metadata["url"].(string),
		ManagersURL: entity["managers_url"].(string),
	}

	return organization, nil
}

type OrgManager struct {
	UserName string
}

//GetOrgManagers returns critical fields from the returned JSON
func (api *APIHelper) GetOrgManagers(cli plugin.CliConnection, orgManagersURL string) ([]OrgManager, error) {
	orgmanagers, err := processPagedResults(cli, orgManagersURL, func(metadata map[string]interface{}, entity map[string]interface{}) interface{} {
		return OrgManager{
			UserName:           entity["username"].(string),
		}
	})

	retVal := make([]OrgManager, len(orgmanagers))
	for i := range orgmanagers {
		retVal[i] = orgmanagers[i].(OrgManager)
	}

	if nil != err {
		return nil, err
	}

	return retVal, nil
}
