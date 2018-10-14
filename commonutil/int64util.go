package commonutil

import "errors"

// AssertInt64 attempts to type-assert the provided interface-argument to int64.
// The provided interface-argument must be a numeric-type or an error is thrown.
func AssertInt64(num interface{}) (int64, error) {
	switch num.(type) {
	case int32:
		v := num.(int32)
		return int64(v), nil
	case int64:
		return num.(int64), nil
	case float32:
		v := num.(float32)
		return int64(v), nil
	case float64:
		v := num.(float64)
		return int64(v), nil
	default:
		return 0, errors.New("assertInt64: Unexpected data-type")
	}
}
