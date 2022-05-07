package slices

import "strings"

type StringSlice struct {
	slice []string
}

func NewStringSlice(slice []string) StringSlice {
	return StringSlice{
		slice: slice,
	}
}

func (ss StringSlice) Contains(element interface{}) bool {
	for _, el := range ss.slice {
		if el == element {
			return true
		}
	}

	return false
}

func (ss StringSlice) Join(separator string) string {
	return strings.Join(ss.slice, separator)
}
