package sets

func (s *Set) SymmetricDifference(set Set) (newSet Set) {
	newSet = s.Union(set)

	for _, element := range set {
		for _, el := range *s {
			if element == el {
				newSet.Discard(el)
			}
		}
	}
	return
}

func (s *Set) SymmetricDifferenceUpdate(set Set) {
	*s = s.SymmetricDifference(set)
}
