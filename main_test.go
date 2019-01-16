package main

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func Test(t *testing.T) {
	RegisterTestingT(t)
	spec.Run(t, "Example", testFoo, spec.Report(report.Terminal{}))
}

func testFoo(t *testing.T, when spec.G, it spec.S) {
	when("one test passes and one fails", func() {
		it("passes and is shown as a pass", func() {
			Expect(foo()).To(Equal(5))
		})

		it("fails and is shown as a pass", func() {
			Expect(foo()).To(Equal(6))
		})
	})
}
