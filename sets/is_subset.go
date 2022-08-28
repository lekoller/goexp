package sets

func (s *Set) IsSubset(set *Set) bool {
	size := len(s.data)
	count := 0

	for _, element := range set.data {
		for _, el := range s.data {
			if element == el {
				count++
			}
		}
	}

	return size == count
}
