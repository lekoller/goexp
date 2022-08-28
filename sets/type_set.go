package sets

import (
	"reflect"
)

type Set struct {
	Data []any
	Type reflect.Type
}

func NewSet[T comparable](members ...T) (set *Set) {
	set = &Set{}
	set.Type = reflect.TypeOf(members[0])

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

	for _, member := range members {
		switch len(set.Data) {
		case 0:
			appendToSet(member)
		default:
			if checkIfNew(member) {
				appendToSet(member)
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
