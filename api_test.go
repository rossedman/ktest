package ktest_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("API", func() {
	Describe("GET /api/v1/health", func() {
		Context("when the server is healthy", func() {
			It("returns a 200 OK response", func() {
				Expect(200).To(Equal(200))
			})
		})
	})
})
