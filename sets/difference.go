package sets

func (s *Set) Difference(sets ...Set) (newSet Set) {
	newSet = s.Copy()
	newSet.DifferenceUpdate(sets...)
	return
}

func (s *Set) DifferenceUpdate(sets ...Set) {
	for _, inSet := range sets {
		for _, inEl := range inSet {
			for i, newEl := range *s {
				if inEl == newEl {
					*s = remove(*s, i)
				}
			}
		}
	}
}
