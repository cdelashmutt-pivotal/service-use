package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestServiceUsePlugin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ServiceUsePlugin Suite")
}
