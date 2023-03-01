package lists

import (
	"reflect"
)

func (l *List) Count(element any) (sum int) {
	for _, el := range *l {
		if reflect.TypeOf(el) == reflect.TypeOf(element) && el == element {
			sum++
		}
	}
	return
}
