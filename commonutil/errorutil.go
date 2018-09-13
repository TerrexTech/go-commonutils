package commonutil

import "fmt"

// ErrorStackTrace returns the StackTrace from the provided error.
// The provided error must be wrapped using package:
//  github.com/pkg/errors
// Example of wrapped error:
//  errors.Wrap(errors.New("some-error"), "some context")
// Errors that are not wrapped will not contain any StackTrace.
// This is intended to be used as a debugging utility, and the usage
// should be discarded once the issue has been debugged.
func ErrorStackTrace(err error) string {
	return fmt.Sprintf("%+v\n", err)
}
