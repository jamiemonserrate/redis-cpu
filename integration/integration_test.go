package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Integration", func() {

	It("should pass", func() {
		Expect(1).To(Equal(1))
	})

})
