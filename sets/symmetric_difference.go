package sets

func (s *Set) SymmetricDifference(set *Set) (newSet *Set) {
	newSet = s.Union(set)

	for _, element := range set.data {
		for _, el := range s.data {
			if element == el {
				newSet.Discard(el)
			}
		}
	}
	return
}

func (s *Set) SymmetricDifferenceUpdate(set *Set) {
	s.setData(s.SymmetricDifference(set).data)
}
