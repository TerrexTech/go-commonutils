package commonutil

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Float64Util", func() {
	Describe("AssertFloat64", func() {
		Context("assert given numeric-interface to float64", func() {
			Specify("int32", func() {
				var v int32 = 4
				var v64 float64 = float64(v)
				float64Val, err := AssertFloat64(v)
				Expect(err).ToNot(HaveOccurred())
				Expect(float64Val).To(Equal(v64))
			})

			Specify("int64", func() {
				var v int64 = 78
				var v64 float64 = float64(v)
				float64Val, err := AssertFloat64(v)
				Expect(err).ToNot(HaveOccurred())
				Expect(float64Val).To(Equal(v64))
			})

			Specify("float32", func() {
				var v float32 = 234
				var v64 float64 = float64(v)
				float64Val, err := AssertFloat64(v)
				Expect(err).ToNot(HaveOccurred())
				Expect(float64Val).To(Equal(v64))
			})

			Specify("float64", func() {
				var v float64 = 457
				float64Val, err := AssertFloat64(v)
				Expect(err).ToNot(HaveOccurred())
				Expect(float64Val).To(Equal(v))
			})
		})

		Context("pass non-numeric interface", func() {
			It("should return an error", func() {
				v := "invalid"
				float64Val, err := AssertFloat64(v)
				Expect(err).To(HaveOccurred())
				Expect(float64Val).To(Equal(float64(0)))
			})
		})
	})
})
