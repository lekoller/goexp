package sets

func (s *Set) Union(sets ...*Set) (newSet *Set) {
	newSet = s.Copy()
	for _, set := range sets {
		for _, element := range set.data {
			newSet.Add(element)
		}
	}
	return
}
