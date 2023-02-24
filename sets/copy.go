package sets

func (s *Set) Copy() (newSet Set) {
	newSet = Set{}

	for _, el := range *s {
		newSet.Add(el)
	}
	return
}
