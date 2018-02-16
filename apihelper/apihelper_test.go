package apihelper

import (
	"bufio"
	"os"

	"github.com/cloudfoundry/cli/plugin/pluginfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func slurp(filename string) []string {
	var b []string
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		b = append(b, scanner.Text())
	}
	return b
}

var _ = Describe("service-use", func() {
	var api CFAPIHelper
	var fakeCliConnection *pluginfakes.FakeCliConnection

	BeforeEach(func() {
		fakeCliConnection = &pluginfakes.FakeCliConnection{}
		api = New(fakeCliConnection)
	})

	Describe("Get services", func() {
		var servicesJSON []string

		BeforeEach(func() {
			servicesJSON = slurp("test-data/services.json")
		})

		It("should return two services", func() {
			fakeCliConnection.CliCommandWithoutTerminalOutputReturns(servicesJSON, nil)
			services, _ := api.GetServices()
			Expect(len(services)).To(Equal(2))
		})
	})

	Describe("Get Org Managers", func() {
		It("should not blow up with empty result", func() {
			var emptyResultJSON = slurp("test-data/empty-result.json")
			fakeCliConnection.CliCommandWithoutTerminalOutputReturns(emptyResultJSON, nil)
			orgManagers, _ := api.GetOrgManagers("/v2/organization/1234/managers")
			Expect(len(orgManagers)).To(Equal(0))
		})

		It("should not blow up with empty username", func() {
			var emptyUserNameResultJSON = slurp("test-data/orgmanagers-blankusername.json")
			fakeCliConnection.CliCommandWithoutTerminalOutputReturns(emptyUserNameResultJSON, nil)
			orgManagers, _ := api.GetOrgManagers("/v2/organization/1234/managers")
			Expect(len(orgManagers)).To(Equal(1))
			Expect(orgManagers[0].UserName).To(BeEmpty())
		})
	})
})
