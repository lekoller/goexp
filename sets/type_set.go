package sets

import (
	"reflect"
)

type Set struct {
	data []any
	kind reflect.Type
}

func NewSet[T comparable](elements ...T) (set *Set) {
	set = &Set{}
	set.setType(elements[0])

	appendToSet := func(m T) {
		set.data = append(set.data, m)
	}
	checkIfNew := func(m T) (first bool) {
		first = true
		for _, el := range set.data {
			if el == m {
				first = false
			}
		}
		return
	}

	for _, element := range elements {
		switch len(set.data) {
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
	arr := reflect.ValueOf(s.Data())

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == element {
			return false
		}
	}

	return true
}

func (s *Set) setType(element any) {
	s.kind = reflect.TypeOf(element)
}

func remove(set []any, index int) []any {
	set[index] = set[len(set)-1]
	return set[:len(set)-1]
}

func (s *Set) setData(data []any) {
	s.data = data
}

func (s *Set) Data() []any {
	return s.data
}

func (s *Set) Type() reflect.Type {
	return s.kind
}
