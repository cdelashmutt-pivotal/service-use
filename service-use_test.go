package main

import (
	"github.com/cdelashmutt-pivotal/service-use/apihelper"
	"github.com/cdelashmutt-pivotal/service-use/apihelper/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ServiceUse", func() {
	var fakeAPI *fakes.FakeCFAPIHelper
	var cmd *ServiceUsePlugin

	BeforeEach(func() {
		fakeAPI = &fakes.FakeCFAPIHelper{}
		cmd = &ServiceUsePlugin{apiHelper: fakeAPI}
	})

	Describe("Get services composes the values correctly", func() {
		service := apihelper.Service{
			URL:             "/v2/services/1234",
			Label:           "FakeService",
			ServicePlansURL: "/v2/services/1234/serviceplans",
		}

		BeforeEach(func() {
			fakeAPI.GetServicesReturns([]apihelper.Service{service}, nil)
		})

		It("should return a service", func() {
			services, err := cmd.getServices()
			Expect(err).To(BeNil())
			Expect(len(services)).To(Equal(1))
		})

	})

})
