package tuples

import (
	"reflect"
)

func (t *Tuple) Count(element any) (sum int) {
	for _, el := range *t {
		if reflect.TypeOf(el) == reflect.TypeOf(element) && el == element {
			sum++
		}
	}
	return
}
