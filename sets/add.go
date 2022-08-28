package sets

import (
	"errors"
	"reflect"
)

func (s *Set) Add(element any) error {
	switch s.Type() {
	case nil:
		s.setType(element)
	default:
		if reflect.TypeOf(element) != s.kind {
			return errors.New("invalid type")
		}
	}

	if !s.checkIfNew(element) {
		return errors.New("not a new element")
	}
	s.setData(append(s.data, element))

	return nil
}
