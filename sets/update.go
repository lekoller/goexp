package sets

func (s *Set) Update(sets ...*Set) {
	s.setData(s.Union(sets...).Data())
}
