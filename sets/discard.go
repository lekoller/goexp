package sets

func (s *Set) Discard(element any) {
	for i, el := range s.Data() {
		if el == element {
			s.setData(remove(s.Data(), i))
		}
	}
}
