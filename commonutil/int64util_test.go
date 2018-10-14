package commonutil

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Int64Util", func() {
	Describe("AssertInt64", func() {
		Context("assert given numeric-interface to int64", func() {
			Specify("int32", func() {
				var v int32 = 4
				var v64 int64 = int64(v)
				int64Val, err := AssertInt64(v)
				Expect(err).ToNot(HaveOccurred())
				Expect(int64Val).To(Equal(v64))
			})

			Specify("int64", func() {
				var v int64 = 78
				int64Val, err := AssertInt64(v)
				Expect(err).ToNot(HaveOccurred())
				Expect(int64Val).To(Equal(v))
			})

			Specify("float32", func() {
				var v float32 = 234
				var v64 int64 = int64(v)
				int64Val, err := AssertInt64(v)
				Expect(err).ToNot(HaveOccurred())
				Expect(int64Val).To(Equal(v64))
			})

			Specify("float64", func() {
				var v float64 = 457
				var v64 int64 = int64(v)
				int64Val, err := AssertInt64(v)
				Expect(err).ToNot(HaveOccurred())
				Expect(int64Val).To(Equal(v64))
			})
		})

		Context("pass non-numeric interface", func() {
			It("should return an error", func() {
				v := "invalid"
				int64Val, err := AssertInt64(v)
				Expect(err).To(HaveOccurred())
				Expect(int64Val).To(Equal(int64(0)))
			})
		})
	})
})
