package commonutil

import "errors"

// AssertFloat64 attempts to type-assert the provided interface-argument to int64.
// The provided interface-argument must be a numeric-type or an error is thrown.
func AssertFloat64(num interface{}) (float64, error) {
	switch num.(type) {
	case int32:
		v := num.(int32)
		return float64(v), nil
	case int64:
		v := num.(int64)
		return float64(v), nil
	case float32:
		v := num.(float32)
		return float64(v), nil
	case float64:
		return num.(float64), nil
	default:
		return 0, errors.New("assertFloat64: Unexpected data-type")
	}
}
