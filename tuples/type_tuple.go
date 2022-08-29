package tuples

type Tuple struct {
	data []any
}

func NewTuple(elements ...any) (tuple *Tuple) {
	var data []any
	data = append(data, elements...)

	tuple = &Tuple{data}
	return
}

func (s *Tuple) Data() []any {
	return s.data
}
