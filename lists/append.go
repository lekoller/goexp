package lists

func (l *List) Append(element any) {
	*l = append(*l, element)
}
