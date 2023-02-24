package sets

import (
	"reflect"
)

type Set []any

func NewSet[T comparable](elements ...T) (set Set) {
	set = Set{}

	appendToSet := func(m T) {
		set = append(set, m)
	}
	checkIfNew := func(m T) (first bool) {
		first = true
		for _, el := range set {
			if reflect.TypeOf(el) == reflect.TypeOf(m) && el == m {
				first = false
			}
		}
		return
	}

	for _, element := range elements {
		switch len(set) {
		case 0:
			appendToSet(element)
		default:
			if checkIfNew(element) {
				appendToSet(element)
			}
		}
	}
	return
}

func (s Set) checkIfNew(element any) bool {
	arr := reflect.ValueOf(s)

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == element {
			return false
		}
	}

	return true
}

func remove(set []any, index int) []any {
	set[index] = set[len(set)-1]
	return set[:len(set)-1]
}
