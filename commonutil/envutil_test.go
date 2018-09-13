package commonutil

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ErrorUtils", func() {
	Describe("ValidateEnv", func() {
		var unsetEnv = func(envVars ...string) {
			for _, v := range envVars {
				os.Unsetenv(v)
			}
		}

		It("should not return any error if all specified env-vars are set", func() {
			// Just trying not to override any existing env-vars
			envVars := []string{
				"v113451e1234",
				"v213451e1234",
				"v313451e1234",
				"v413451e1234",
			}

			for _, v := range envVars {
				os.Setenv(v, v)
			}

			varName, err := ValidateEnv(envVars...)
			Expect(err).ToNot(HaveOccurred())
			Expect(varName).To(BeEmpty())
			unsetEnv(envVars...)
		})

		It("should return error if some env-vars is missing", func() {
			envVars := []string{
				"v113451e1234",
				"v213451e1234",
				"v513451e1234",
			}

			for _, v := range envVars {
				os.Setenv(v, v)
			}

			envVars = append(envVars, "v313451e1234")
			envVars = append(envVars, "v413451e1234")
			varName, err := ValidateEnv(envVars...)
			Expect(err).To(HaveOccurred())
			Expect(varName).To(Equal("v313451e1234"))
			unsetEnv(envVars...)
		})
	})
})
