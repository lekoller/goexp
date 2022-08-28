package sets

func (s *Set) IsSuperset(set *Set) bool {
	return set.IsSubset(s)
}
