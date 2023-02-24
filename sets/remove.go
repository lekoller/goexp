package sets

import "errors"

func (s *Set) Remove(element any) error {
	sizeBefore := len(*s)
	s.Discard(element)
	sizeAfter := len(*s)

	if sizeBefore == sizeAfter {
		return errors.New("element not found")
	}
	return nil
}
