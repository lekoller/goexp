package sets

import (
	"math/rand"
	"time"
)

func (s *Set) Pop() (element any) {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(s.Data))

	element = s.Data[index]
	s.Data = remove(s.Data, index)
	return
}
