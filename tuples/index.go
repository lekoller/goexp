package tuples

func (t *Tuple) Index(element any) (index int) {
	for i, el := range t.data {
		if el == element {
			index = i
			return
		}
	}
	return -1
}
