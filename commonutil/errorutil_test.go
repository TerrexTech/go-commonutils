package commonutil

import (
	"regexp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
)

var _ = Describe("ErrorUtils", func() {
	Describe("ErrorStackTrace", func() {
		It("should return StackTrace of provided error", func() {
			err := errors.Wrap(errors.New("some-error"), "some-context")
			st := ErrorStackTrace(err)

			//  Sample string that passes this regex: "errorutils_test.go:16"
			rgx := regexp.MustCompile(`errorutil_test\.go:[0-9]+`)
			match := rgx.FindString(st)

			// Check if the error has regex match, then its a StackTrace
			// and our test is good.
			Expect(match).ToNot(BeEmpty())
		})
	})
})
