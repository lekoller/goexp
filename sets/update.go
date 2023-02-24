package sets

func (s *Set) Update(sets ...Set) {
	*s = s.Union(sets...)
}
