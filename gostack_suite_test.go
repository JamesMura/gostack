package gostack_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"

    "testing"
)

func TestGostack(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Gostack Suite")
}
