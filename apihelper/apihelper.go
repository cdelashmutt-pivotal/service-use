package apihelper

import (
	"strconv"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/krujos/cfcurl"
)

//CFAPIHelper to wrap cf curl results
type CFAPIHelper interface {
	GetServices() ([]Service, error)
	GetServicePlans(plansURL string) ([]ServicePlan, error)
	GetServiceInstances(serviceInstancesURL string) ([]ServiceInstance, error)
	GetSpace(spaceURL string) (Space, error)
	GetOrganization(organizationURL string) (Organization, error)
	GetOrgManagers(orgManagersURL string) ([]OrgManager, error)
}

//APIHelper implementation
type APIHelper struct {
	cli plugin.CliConnection
}

func New(cli plugin.CliConnection) CFAPIHelper {
	return &APIHelper{cli}
}

//Function type to simplify processing paged results
type process func(metadata map[string]interface{}, entity map[string]interface{}) interface{}

//Base method to process paged results from API calls
func (api *APIHelper) processPagedResults(url string, fn process) ([]interface{}, error) {

	theJSON, err := cfcurl.Curl(api.cli, url)
	if nil != err {
		return nil, err
	}

	pages := int(theJSON["total_pages"].(float64))
	var objects []interface{}
	for i := 1; i <= pages; i++ {
		if 1 != i {
			theJSON, err = cfcurl.Curl(api.cli, url+"?page="+strconv.Itoa(i))
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
func (api *APIHelper) GetServices() ([]Service, error) {
	services, err := api.processPagedResults("/v2/services", func(metadata map[string]interface{}, entity map[string]interface{}) interface{} {
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
func (api *APIHelper) GetServicePlans(plansURL string) ([]ServicePlan, error) {
	serviceplans, err := api.processPagedResults(plansURL, func(metadata map[string]interface{}, entity map[string]interface{}) interface{} {
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
func (api *APIHelper) GetServiceInstances(serviceInstancesURL string) ([]ServiceInstance, error) {
	serviceinstances, err := api.processPagedResults(serviceInstancesURL, func(metadata map[string]interface{}, entity map[string]interface{}) interface{} {
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
func (api *APIHelper) GetSpace(spaceURL string) (Space, error) {
	theJSON, err := cfcurl.Curl(api.cli, spaceURL)
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
	Name        string
	URL         string
	ManagersURL string
}

//GetOrganization returns a struct that represents critical fields in the JSON
func (api *APIHelper) GetOrganization(organizationURL string) (Organization, error) {
	theJSON, err := cfcurl.Curl(api.cli, organizationURL)
	if nil != err {
		return Organization{}, err
	}

	entity := theJSON["entity"].(map[string]interface{})
	metadata := theJSON["metadata"].(map[string]interface{})

	organization := Organization{
		Name:        entity["name"].(string),
		URL:         metadata["url"].(string),
		ManagersURL: entity["managers_url"].(string),
	}

	return organization, nil
}

type OrgManager struct {
	UserName string
}

//GetOrgManagers returns critical fields from the returned JSON
func (api *APIHelper) GetOrgManagers(orgManagersURL string) ([]OrgManager, error) {
	orgmanagers, err := api.processPagedResults(orgManagersURL, func(metadata map[string]interface{}, entity map[string]interface{}) interface{} {
		username, _ := entity["username"].(string)
		return OrgManager{
			UserName: username,
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
