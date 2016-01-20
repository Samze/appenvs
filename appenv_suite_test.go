package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAppEnv(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "AppEnv suite")
}
