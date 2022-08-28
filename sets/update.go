package sets

func (s *Set) Update(sets ...*Set) {
	s.Data = s.Union(sets...).Data
}
