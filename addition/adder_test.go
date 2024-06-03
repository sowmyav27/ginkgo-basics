package addition

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Adder", func() {
	Describe("Add", func() {
		Context("When numbers are positive", func() {
			It("Adds two numbers", func() {
				sum, err := Add(2, 3)
				Expect(err).NotTo(HaveOccurred())
				Expect(sum).To(Equal(5))

			})
		})
		Context("When numbers are negative", func() {
			It("Returns an error", func() {
				_, err := Add(-1, -1)
				Expect(err).To(HaveOccurred())

			})

		})
		Context("When numbers are equal to 0", func() {
			It("Returns zero", func() {
				sum, err := Add(0, 0)
				Expect(err).NotTo(HaveOccurred())
				Expect(sum).To(Equal(0))

			})

		})

	})

})
