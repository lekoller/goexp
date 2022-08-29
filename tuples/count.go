package tuples

func (t *Tuple) Count(element any) (sum int) {
	for _, el := range t.data {
		if el == element {
			sum++
		}
	}
	return
}
