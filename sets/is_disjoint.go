package sets

func (s *Set) IsDisjoint(set *Set) bool {
	return len(s.Intersection(set).Data()) == 0
}
