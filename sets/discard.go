package sets

func (s *Set) Discard(element any) {
	for i, el := range s.Data {
		if el == element {
			s.Data = remove(s.Data, i)
		}
	}
}

func remove(set []any, index int) []any {
	set[index] = set[len(set)-1]
	return set[:len(set)-1]
}
