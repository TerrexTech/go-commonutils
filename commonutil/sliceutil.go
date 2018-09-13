package commonutil

import (
	"reflect"
)

// IsElementInSlice checks if the specified element is present in slice.
// Returns false if a non-slice argument is provided as first parameter.
func IsElementInSlice(slice interface{}, in interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		return false
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == in {
			return true
		}
	}

	return false
}

// IndexInSlice returns the index of element in specified slice.
// Returns false if a non-slice argument is provided as first parameter.
func IndexInSlice(slice interface{}, in interface{}) int {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		return -1
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == in {
			return i
		}
	}

	return -1
}

// AreElementsInSlice checks if the elements in specified slice
// are present in the given slice.
// Note: All elements must be present. This function simply runs
// two loops with O^2 complexity, not ideal for large slices.
// Returns false if either of the parameters is a non-slice.
func AreElementsInSlice(slice interface{}, in interface{}) bool {
	s := reflect.ValueOf(slice)
	i := reflect.ValueOf(in)

	if s.Kind() != reflect.Slice || i.Kind() != reflect.Slice {
		return false
	}

	for si := 0; si < s.Len(); si++ {
		sv := s.Index(si).Interface()
		isFound := false

		for ii := 0; ii < i.Len(); ii++ {
			if i.Index(ii).Interface() == sv {
				isFound = true
				break
			}
		}

		if !isFound {
			return false
		}
	}

	return true
}

// AreElementsInSliceStrict is same as AreElementsInSlice, except
// this also checks if the provided slices have same length.
func AreElementsInSliceStrict(slice interface{}, in interface{}) bool {
	s := reflect.ValueOf(slice)
	i := reflect.ValueOf(in)

	if s.Kind() != reflect.Slice || i.Kind() != reflect.Slice {
		return false
	}

	if s.Len() != i.Len() {
		return false
	}

	return AreElementsInSlice(slice, in)
}
