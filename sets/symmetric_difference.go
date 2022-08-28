package sets

func (s *Set) SymmetricDifference(set *Set) (newSet *Set) {
	newSet = s.Union(set)

	for _, element := range set.Data {
		for _, el := range s.Data {
			if element == el {
				newSet.Discard(el)
			}
		}
	}
	return
}

func (s *Set) SymmetricDifferenceUpdate(set *Set) {
	s.Data = s.SymmetricDifference(set).Data
}
