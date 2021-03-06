package commonutil

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StringUtils", func() {
	Describe("GetHostsFromEnv", func() {
		It("should parse the env var and get hosts", func() {
			// Should also handle spaces after comma; we test that, too
			hosts := *ParseHosts("192.168.1.1:8080, 172.17.1.1,10.15.1.1")

			expectedHosts := []string{
				"192.168.1.1:8080",
				"172.17.1.1",
				"10.15.1.1",
			}
			Expect(
				AreElementsInSlice(hosts, expectedHosts),
			).To(BeTrue())
		})

		It("should return blank slice if empty string is passed", func() {
			hosts := *ParseHosts("")
			Expect(hosts).To(Equal([]string{}))
			Expect(hosts).ToNot(BeNil())
		})
	})

	Describe("StandardizeSpaces", func() {
		It("should standardize spaces in a string", func() {
			s := `
			Some string    with irregular

			spaces
		`
			expected := "Some string with irregular spaces"
			Expect(StandardizeSpaces(s)).To(Equal(expected))
		})
	})
})
