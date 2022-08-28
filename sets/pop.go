package sets

import (
	"math/rand"
	"time"
)

func (s *Set) Pop() (element any) {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(s.data))

	element = s.data[index]
	s.setData(remove(s.data, index))
	return
}
