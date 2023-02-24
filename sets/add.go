package sets

import (
	"errors"
)

func (s *Set) Add(element any) error {
	if !s.checkIfNew(element) {
		return errors.New("not a new element")
	}
	*s = append(*s, element)

	return nil
}
