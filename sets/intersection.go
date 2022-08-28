package sets

func (s *Set) Intersection(sets ...*Set) (newSet *Set) {
	numberOfSets := len(sets)
	repetitionMap := make(map[any]int)
	newSet = &Set{}
	newSet.setType(sets[0].Data[0])

	for _, inSet := range sets {
		for _, inEl := range inSet.Data {
			for _, sEl := range s.Data {
				if inEl == sEl {
					times := 1
					if val, ok := repetitionMap[inEl]; ok {
						times = val + 1
					}
					if times == numberOfSets {
						newSet.Add(inEl)
					}
					repetitionMap[inEl] = times
				}
			}
		}
	}
	return
}

func (s *Set) IntersectionUpdate(sets ...*Set) {
	s.Data = s.Intersection(sets...).Data
}
