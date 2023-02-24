package sets

func (s *Set) IsSubset(set Set) bool {
	size := len(*s)
	count := 0

	for _, element := range set {
		for _, el := range *s {
			if element == el {
				count++
			}
		}
	}

	return size == count
}
