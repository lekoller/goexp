package sets

import (
	"errors"
	"reflect"
)

func (s *Set) Add(element any) error {
	if reflect.TypeOf(element) != s.Type {
		return errors.New("invalid type")
	}
	if !s.checkIfNew(element) {
		return errors.New("not a new element")
	}
	s.Data = append(s.Data, element)

	return nil
}
