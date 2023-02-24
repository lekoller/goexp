package tuples

type Tuple []any

func NewTuple(elements ...any) (tuple Tuple) {
	tuple = Tuple{}

	for _, element := range elements {
		tuple = append(tuple, element)
	}

	return
}
