package lists

import (
	"reflect"
)

func (l *List) Extend(other ...any) {
	if len(other) == 1 && reflect.TypeOf(other[0]).String() == "string" {
		if els, ok := other[0].([]interface{}); ok {
			for _, el := range els {
				l.Append(el)
			}
		}
		return
	}
	for _, el := range other {
		l.Append(el)
	}
}
