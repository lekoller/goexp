package sets

func (s *Set) Copy() (newSet *Set) {
	newSet = &Set{}
	newSet.Data = s.Data
	newSet.Type = s.Type
	return
}
