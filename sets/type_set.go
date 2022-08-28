package sets

import (
	"reflect"
)

type Set struct {
	Data []any
	Type reflect.Type
}

func NewSet[T comparable](elements ...T) (set *Set) {
	set = &Set{}
	set.setType(elements[0])

	appendToSet := func(m T) {
		set.Data = append(set.Data, m)
	}
	checkIfNew := func(m T) (first bool) {
		first = true
		for _, el := range set.Data {
			if el == m {
				first = false
			}
		}
		return
	}

	for _, element := range elements {
		switch len(set.Data) {
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
	arr := reflect.ValueOf(s.Data)

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == element {
			return false
		}
	}

	return true
}

func (s *Set) setType(element any) {
	s.Type = reflect.TypeOf(element)
}

func remove(set []any, index int) []any {
	set[index] = set[len(set)-1]
	return set[:len(set)-1]
}
