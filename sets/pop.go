package sets

import (
	"math/rand"
	"time"
)

func (s *Set) Pop() (element any) {
	newSet := Set{}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(*s))

	for i, el := range *s {
		if i == index {
			element = el
		} else {
			newSet.Add(el)
		}
	}
	*s = newSet

	return
}
