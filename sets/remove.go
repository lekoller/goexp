package sets

import "errors"

func (s *Set) Remove(element any) error {
	sizeBefore := len(s.Data)
	s.Discard(element)
	sizeAfter := len(s.Data)

	if sizeBefore == sizeAfter {
		return errors.New("element not found")
	}
	return nil
}
