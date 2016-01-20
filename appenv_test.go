package main

import (
	"github.com/cloudfoundry/cli/plugin/fakes"
	ioStub "github.com/cloudfoundry/cli/testhelpers/io"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("appenv", func() {
	var fakeCliConnection *fakes.FakeCliConnection
	var appenv *AppEnv

	BeforeEach(func() {
		fakeCliConnection = &fakes.FakeCliConnection{}
		appenv = &AppEnv{}
	})

	Context("when uninstalled", func() {
		It("does nothing", func() {
			output := ioStub.CaptureOutput(func() {
				appenv.Run(fakeCliConnection, []string{"CLI-MESSAGE-UNINSTALL"})
			})

			Expect(output).To(ConsistOf([]string{""}))
		})
	})

	Context("when not logged in", func() {
		It("returns an error", func() {
			fakeCliConnection.IsLoggedInStub = func() (bool, error) { return false, nil }
			_, err := appenv.GetEnvs(fakeCliConnection)
			Expect(err).To(Not(BeNil()))
		})
	})
})
