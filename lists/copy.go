package lists

func (l *List) Copy() (new List) {
	new = List{}

	for _, el := range *l {
		new.Append(el)
	}
	return
}
