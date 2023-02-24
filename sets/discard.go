package sets

func (s *Set) Discard(element any) {
	for i, el := range *s {
		if el == element {
			*s = (remove(*s, i))
		}
	}
}
