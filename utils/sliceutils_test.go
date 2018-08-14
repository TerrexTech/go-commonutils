package utils

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// TestSliceUtils tests the critical slice-utils functions
func TestSliceUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SliceUtils Suite")
}

var _ = Describe("SliceUtils", func() {
	var (
		s []string
		i []int
	)

	BeforeEach(func() {
		s = []string{"e1", "e2", "e3", "e4"}
		i = []int{1, 4, 7, 8}
	})

	It("should check if element is in slice", func() {
		Expect(IsElementInSlice(s, "e3")).To(BeTrue())
		Expect(IsElementInSlice(s, "e5")).To(BeFalse())

		Expect(IsElementInSlice(i, 4)).To(BeTrue())
		Expect(IsElementInSlice(i, 6)).To(BeFalse())
	})

	It("should return the index of element in slice", func() {
		Expect(IndexInSlice(s, "e2")).To(Equal(1))
		Expect(IndexInSlice(s, "e5")).To(Equal(-1))

		Expect(IndexInSlice(i, 7)).To(Equal(2))
		Expect(IndexInSlice(i, 5)).To(Equal(-1))
	})

	It("should check if all elements are in a slice", func() {
		// All elements
		s2 := []string{"e3", "e2", "e4", "e1"}
		Expect(AreElementsInSlice(s, s2)).To(BeTrue())
		i2 := []int{1, 7, 4, 8}
		Expect(AreElementsInSlice(i, i2)).To(BeTrue())

		// An invalid element
		s3 := []string{"e3", "e6", "e4", "e1"}
		Expect(AreElementsInSlice(s, s3)).To(BeFalse())
		i3 := []int{1, 5, 7, 8}
		Expect(AreElementsInSlice(i, i3)).To(BeFalse())

		// A missing element
		s4 := []string{"e3", "e2", "e4"}
		Expect(AreElementsInSlice(s, s4)).To(BeFalse())
		i4 := []int{1, 7, 4}
		Expect(AreElementsInSlice(i, i4)).To(BeFalse())
	})

	It("should return negative results if non-slice argument is passed here slice is expected", func() {
		Expect(IsElementInSlice(1, "s")).To(BeFalse())
		Expect(IndexInSlice(6, 0)).To(Equal(-1))
		Expect(AreElementsInSlice("a", 5)).To(BeFalse())
		Expect(AreElementsInSlice(s, -1)).To(BeFalse())
	})

	It("should return negative results if data-types mismatch", func() {
		Expect(IsElementInSlice(s, 1)).To(BeFalse())
		Expect(IndexInSlice(i, "g")).To(Equal(-1))
		Expect(AreElementsInSlice(s, i)).To(BeFalse())
	})
})
