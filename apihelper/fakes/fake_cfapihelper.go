// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/cdelashmutt-pivotal/service-use/apihelper"
)

type FakeCFAPIHelper struct {
	GetServicesStub        func() ([]apihelper.Service, error)
	getServicesMutex       sync.RWMutex
	getServicesArgsForCall []struct{}
	getServicesReturns     struct {
		result1 []apihelper.Service
		result2 error
	}
	getServicesReturnsOnCall map[int]struct {
		result1 []apihelper.Service
		result2 error
	}
	GetServicePlansStub        func(plansURL string) ([]apihelper.ServicePlan, error)
	getServicePlansMutex       sync.RWMutex
	getServicePlansArgsForCall []struct {
		plansURL string
	}
	getServicePlansReturns struct {
		result1 []apihelper.ServicePlan
		result2 error
	}
	getServicePlansReturnsOnCall map[int]struct {
		result1 []apihelper.ServicePlan
		result2 error
	}
	GetServiceInstancesStub        func(serviceInstancesURL string) ([]apihelper.ServiceInstance, error)
	getServiceInstancesMutex       sync.RWMutex
	getServiceInstancesArgsForCall []struct {
		serviceInstancesURL string
	}
	getServiceInstancesReturns struct {
		result1 []apihelper.ServiceInstance
		result2 error
	}
	getServiceInstancesReturnsOnCall map[int]struct {
		result1 []apihelper.ServiceInstance
		result2 error
	}
	GetSpaceStub        func(spaceURL string) (apihelper.Space, error)
	getSpaceMutex       sync.RWMutex
	getSpaceArgsForCall []struct {
		spaceURL string
	}
	getSpaceReturns struct {
		result1 apihelper.Space
		result2 error
	}
	getSpaceReturnsOnCall map[int]struct {
		result1 apihelper.Space
		result2 error
	}
	GetOrganizationStub        func(organizationURL string) (apihelper.Organization, error)
	getOrganizationMutex       sync.RWMutex
	getOrganizationArgsForCall []struct {
		organizationURL string
	}
	getOrganizationReturns struct {
		result1 apihelper.Organization
		result2 error
	}
	getOrganizationReturnsOnCall map[int]struct {
		result1 apihelper.Organization
		result2 error
	}
	GetOrgManagersStub        func(orgManagersURL string) ([]apihelper.OrgManager, error)
	getOrgManagersMutex       sync.RWMutex
	getOrgManagersArgsForCall []struct {
		orgManagersURL string
	}
	getOrgManagersReturns struct {
		result1 []apihelper.OrgManager
		result2 error
	}
	getOrgManagersReturnsOnCall map[int]struct {
		result1 []apihelper.OrgManager
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCFAPIHelper) GetServices() ([]apihelper.Service, error) {
	fake.getServicesMutex.Lock()
	ret, specificReturn := fake.getServicesReturnsOnCall[len(fake.getServicesArgsForCall)]
	fake.getServicesArgsForCall = append(fake.getServicesArgsForCall, struct{}{})
	fake.recordInvocation("GetServices", []interface{}{})
	fake.getServicesMutex.Unlock()
	if fake.GetServicesStub != nil {
		return fake.GetServicesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getServicesReturns.result1, fake.getServicesReturns.result2
}

func (fake *FakeCFAPIHelper) GetServicesCallCount() int {
	fake.getServicesMutex.RLock()
	defer fake.getServicesMutex.RUnlock()
	return len(fake.getServicesArgsForCall)
}

func (fake *FakeCFAPIHelper) GetServicesReturns(result1 []apihelper.Service, result2 error) {
	fake.GetServicesStub = nil
	fake.getServicesReturns = struct {
		result1 []apihelper.Service
		result2 error
	}{result1, result2}
}

func (fake *FakeCFAPIHelper) GetServicesReturnsOnCall(i int, result1 []apihelper.Service, result2 error) {
	fake.GetServicesStub = nil
	if fake.getServicesReturnsOnCall == nil {
		fake.getServicesReturnsOnCall = make(map[int]struct {
			result1 []apihelper.Service
			result2 error
		})
	}
	fake.getServicesReturnsOnCall[i] = struct {
		result1 []apihelper.Service
		result2 error
	}{result1, result2}
}

func (fake *FakeCFAPIHelper) GetServicePlans(plansURL string) ([]apihelper.ServicePlan, error) {
	fake.getServicePlansMutex.Lock()
	ret, specificReturn := fake.getServicePlansReturnsOnCall[len(fake.getServicePlansArgsForCall)]
	fake.getServicePlansArgsForCall = append(fake.getServicePlansArgsForCall, struct {
		plansURL string
	}{plansURL})
	fake.recordInvocation("GetServicePlans", []interface{}{plansURL})
	fake.getServicePlansMutex.Unlock()
	if fake.GetServicePlansStub != nil {
		return fake.GetServicePlansStub(plansURL)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getServicePlansReturns.result1, fake.getServicePlansReturns.result2
}

func (fake *FakeCFAPIHelper) GetServicePlansCallCount() int {
	fake.getServicePlansMutex.RLock()
	defer fake.getServicePlansMutex.RUnlock()
	return len(fake.getServicePlansArgsForCall)
}

func (fake *FakeCFAPIHelper) GetServicePlansArgsForCall(i int) string {
	fake.getServicePlansMutex.RLock()
	defer fake.getServicePlansMutex.RUnlock()
	return fake.getServicePlansArgsForCall[i].plansURL
}

func (fake *FakeCFAPIHelper) GetServicePlansReturns(result1 []apihelper.ServicePlan, result2 error) {
	fake.GetServicePlansStub = nil
	fake.getServicePlansReturns = struct {
		result1 []apihelper.ServicePlan
		result2 error
	}{result1, result2}
}

func (fake *FakeCFAPIHelper) GetServicePlansReturnsOnCall(i int, result1 []apihelper.ServicePlan, result2 error) {
	fake.GetServicePlansStub = nil
	if fake.getServicePlansReturnsOnCall == nil {
		fake.getServicePlansReturnsOnCall = make(map[int]struct {
			result1 []apihelper.ServicePlan
			result2 error
		})
	}
	fake.getServicePlansReturnsOnCall[i] = struct {
		result1 []apihelper.ServicePlan
		result2 error
	}{result1, result2}
}

func (fake *FakeCFAPIHelper) GetServiceInstances(serviceInstancesURL string) ([]apihelper.ServiceInstance, error) {
	fake.getServiceInstancesMutex.Lock()
	ret, specificReturn := fake.getServiceInstancesReturnsOnCall[len(fake.getServiceInstancesArgsForCall)]
	fake.getServiceInstancesArgsForCall = append(fake.getServiceInstancesArgsForCall, struct {
		serviceInstancesURL string
	}{serviceInstancesURL})
	fake.recordInvocation("GetServiceInstances", []interface{}{serviceInstancesURL})
	fake.getServiceInstancesMutex.Unlock()
	if fake.GetServiceInstancesStub != nil {
		return fake.GetServiceInstancesStub(serviceInstancesURL)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getServiceInstancesReturns.result1, fake.getServiceInstancesReturns.result2
}

func (fake *FakeCFAPIHelper) GetServiceInstancesCallCount() int {
	fake.getServiceInstancesMutex.RLock()
	defer fake.getServiceInstancesMutex.RUnlock()
	return len(fake.getServiceInstancesArgsForCall)
}

func (fake *FakeCFAPIHelper) GetServiceInstancesArgsForCall(i int) string {
	fake.getServiceInstancesMutex.RLock()
	defer fake.getServiceInstancesMutex.RUnlock()
	return fake.getServiceInstancesArgsForCall[i].serviceInstancesURL
}

func (fake *FakeCFAPIHelper) GetServiceInstancesReturns(result1 []apihelper.ServiceInstance, result2 error) {
	fake.GetServiceInstancesStub = nil
	fake.getServiceInstancesReturns = struct {
		result1 []apihelper.ServiceInstance
		result2 error
	}{result1, result2}
}

func (fake *FakeCFAPIHelper) GetServiceInstancesReturnsOnCall(i int, result1 []apihelper.ServiceInstance, result2 error) {
	fake.GetServiceInstancesStub = nil
	if fake.getServiceInstancesReturnsOnCall == nil {
		fake.getServiceInstancesReturnsOnCall = make(map[int]struct {
			result1 []apihelper.ServiceInstance
			result2 error
		})
	}
	fake.getServiceInstancesReturnsOnCall[i] = struct {
		result1 []apihelper.ServiceInstance
		result2 error
	}{result1, result2}
}

func (fake *FakeCFAPIHelper) GetSpace(spaceURL string) (apihelper.Space, error) {
	fake.getSpaceMutex.Lock()
	ret, specificReturn := fake.getSpaceReturnsOnCall[len(fake.getSpaceArgsForCall)]
	fake.getSpaceArgsForCall = append(fake.getSpaceArgsForCall, struct {
		spaceURL string
	}{spaceURL})
	fake.recordInvocation("GetSpace", []interface{}{spaceURL})
	fake.getSpaceMutex.Unlock()
	if fake.GetSpaceStub != nil {
		return fake.GetSpaceStub(spaceURL)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getSpaceReturns.result1, fake.getSpaceReturns.result2
}

func (fake *FakeCFAPIHelper) GetSpaceCallCount() int {
	fake.getSpaceMutex.RLock()
	defer fake.getSpaceMutex.RUnlock()
	return len(fake.getSpaceArgsForCall)
}

func (fake *FakeCFAPIHelper) GetSpaceArgsForCall(i int) string {
	fake.getSpaceMutex.RLock()
	defer fake.getSpaceMutex.RUnlock()
	return fake.getSpaceArgsForCall[i].spaceURL
}

func (fake *FakeCFAPIHelper) GetSpaceReturns(result1 apihelper.Space, result2 error) {
	fake.GetSpaceStub = nil
	fake.getSpaceReturns = struct {
		result1 apihelper.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeCFAPIHelper) GetSpaceReturnsOnCall(i int, result1 apihelper.Space, result2 error) {
	fake.GetSpaceStub = nil
	if fake.getSpaceReturnsOnCall == nil {
		fake.getSpaceReturnsOnCall = make(map[int]struct {
			result1 apihelper.Space
			result2 error
		})
	}
	fake.getSpaceReturnsOnCall[i] = struct {
		result1 apihelper.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeCFAPIHelper) GetOrganization(organizationURL string) (apihelper.Organization, error) {
	fake.getOrganizationMutex.Lock()
	ret, specificReturn := fake.getOrganizationReturnsOnCall[len(fake.getOrganizationArgsForCall)]
	fake.getOrganizationArgsForCall = append(fake.getOrganizationArgsForCall, struct {
		organizationURL string
	}{organizationURL})
	fake.recordInvocation("GetOrganization", []interface{}{organizationURL})
	fake.getOrganizationMutex.Unlock()
	if fake.GetOrganizationStub != nil {
		return fake.GetOrganizationStub(organizationURL)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getOrganizationReturns.result1, fake.getOrganizationReturns.result2
}

func (fake *FakeCFAPIHelper) GetOrganizationCallCount() int {
	fake.getOrganizationMutex.RLock()
	defer fake.getOrganizationMutex.RUnlock()
	return len(fake.getOrganizationArgsForCall)
}

func (fake *FakeCFAPIHelper) GetOrganizationArgsForCall(i int) string {
	fake.getOrganizationMutex.RLock()
	defer fake.getOrganizationMutex.RUnlock()
	return fake.getOrganizationArgsForCall[i].organizationURL
}

func (fake *FakeCFAPIHelper) GetOrganizationReturns(result1 apihelper.Organization, result2 error) {
	fake.GetOrganizationStub = nil
	fake.getOrganizationReturns = struct {
		result1 apihelper.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeCFAPIHelper) GetOrganizationReturnsOnCall(i int, result1 apihelper.Organization, result2 error) {
	fake.GetOrganizationStub = nil
	if fake.getOrganizationReturnsOnCall == nil {
		fake.getOrganizationReturnsOnCall = make(map[int]struct {
			result1 apihelper.Organization
			result2 error
		})
	}
	fake.getOrganizationReturnsOnCall[i] = struct {
		result1 apihelper.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeCFAPIHelper) GetOrgManagers(orgManagersURL string) ([]apihelper.OrgManager, error) {
	fake.getOrgManagersMutex.Lock()
	ret, specificReturn := fake.getOrgManagersReturnsOnCall[len(fake.getOrgManagersArgsForCall)]
	fake.getOrgManagersArgsForCall = append(fake.getOrgManagersArgsForCall, struct {
		orgManagersURL string
	}{orgManagersURL})
	fake.recordInvocation("GetOrgManagers", []interface{}{orgManagersURL})
	fake.getOrgManagersMutex.Unlock()
	if fake.GetOrgManagersStub != nil {
		return fake.GetOrgManagersStub(orgManagersURL)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getOrgManagersReturns.result1, fake.getOrgManagersReturns.result2
}

func (fake *FakeCFAPIHelper) GetOrgManagersCallCount() int {
	fake.getOrgManagersMutex.RLock()
	defer fake.getOrgManagersMutex.RUnlock()
	return len(fake.getOrgManagersArgsForCall)
}

func (fake *FakeCFAPIHelper) GetOrgManagersArgsForCall(i int) string {
	fake.getOrgManagersMutex.RLock()
	defer fake.getOrgManagersMutex.RUnlock()
	return fake.getOrgManagersArgsForCall[i].orgManagersURL
}

func (fake *FakeCFAPIHelper) GetOrgManagersReturns(result1 []apihelper.OrgManager, result2 error) {
	fake.GetOrgManagersStub = nil
	fake.getOrgManagersReturns = struct {
		result1 []apihelper.OrgManager
		result2 error
	}{result1, result2}
}

func (fake *FakeCFAPIHelper) GetOrgManagersReturnsOnCall(i int, result1 []apihelper.OrgManager, result2 error) {
	fake.GetOrgManagersStub = nil
	if fake.getOrgManagersReturnsOnCall == nil {
		fake.getOrgManagersReturnsOnCall = make(map[int]struct {
			result1 []apihelper.OrgManager
			result2 error
		})
	}
	fake.getOrgManagersReturnsOnCall[i] = struct {
		result1 []apihelper.OrgManager
		result2 error
	}{result1, result2}
}

func (fake *FakeCFAPIHelper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getServicesMutex.RLock()
	defer fake.getServicesMutex.RUnlock()
	fake.getServicePlansMutex.RLock()
	defer fake.getServicePlansMutex.RUnlock()
	fake.getServiceInstancesMutex.RLock()
	defer fake.getServiceInstancesMutex.RUnlock()
	fake.getSpaceMutex.RLock()
	defer fake.getSpaceMutex.RUnlock()
	fake.getOrganizationMutex.RLock()
	defer fake.getOrganizationMutex.RUnlock()
	fake.getOrgManagersMutex.RLock()
	defer fake.getOrgManagersMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCFAPIHelper) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ apihelper.CFAPIHelper = new(FakeCFAPIHelper)
