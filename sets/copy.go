package sets

func (s *Set) Copy() (newSet *Set) {
	newSet = &Set{}

	for _, el := range s.data {
		newSet.Add(el)
	}
	return
}
