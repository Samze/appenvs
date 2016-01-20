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
		fakeCliConnection.IsLoggedInStub = func() (bool, error) { return true, nil }

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
		BeforeEach(func() {
			fakeCliConnection.IsLoggedInStub = func() (bool, error) { return false, nil }
		})

		It("returns an error message", func() {
			_, err := appenv.GetEnvs(fakeCliConnection, []string{"app_name"})
			Expect(err).To(MatchError("You must login first!"))
		})
	})

	Context("when no app name is supplied", func() {
		It("returns an error message", func() {
			_, err := appenv.GetEnvs(fakeCliConnection, []string{})
			Expect(err).To(MatchError("You must specify an app name"))
		})
	})

	Context("when getting vcap_services", func() {
		appname := "APP_NAME"
		BeforeEach(func() {
			fakeCliConnection.CliCommandWithoutTerminalOutputReturns([]string{"something"}, nil)
			appenv.GetEnvs(fakeCliConnection, []string{appname})
		})

		It("calls cli with env", func() {
			Expect(fakeCliConnection.CliCommandWithoutTerminalOutputCallCount()).
				To(Not(BeZero()))
		})

		It("requests the correct app envs", func() {
			Expect(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(0)).
				To(Equal([]string{"env", appname}))
		})
	})

})
