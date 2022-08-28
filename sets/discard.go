package sets

func (s *Set) Discard(element any) {
	for i, el := range s.Data {
		if el == element {
			s.Data = remove(s.Data, i)
		}
	}
}
