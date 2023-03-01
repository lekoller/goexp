package lists

type List []any

func NewList(elements ...any) (list List) {
	list = List{}

	for _, element := range elements {
		list = append(list, element)
	}

	return
}
