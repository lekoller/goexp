package sets

func (s *Set) IsSubset(set *Set) bool {
	size := len(s.Data)
	count := 0

	for _, element := range set.Data {
		for _, el := range s.Data {
			if element == el {
				count++
			}
		}
	}

	return size == count
}
