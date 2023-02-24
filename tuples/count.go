package tuples

import (
	"log"
	"reflect"
)

func (t *Tuple) Count(element any) (sum int) {
	log.Println(*t)
	for _, el := range *t {
		println()
		log.Println(element)
		log.Println(el)
		println()
		if reflect.TypeOf(el) == reflect.TypeOf(element) && el == element {
			sum++
		}
	}
	return
}
