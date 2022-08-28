package sets

func (s *Set) Difference(sets ...*Set) (newSet *Set) {
	newSet = s.Copy()
	newSet.DifferenceUpdate(sets...)
	return
}

func (s *Set) DifferenceUpdate(sets ...*Set) {
	for _, inSet := range sets {
		for _, inEl := range inSet.Data {
			for i, newEl := range s.Data {
				if inEl == newEl {
					s.Data = remove(s.Data, i)
				}
			}
		}
	}
}
