package sets_test

import (
	"log"
	"testing"

	"github.com/lekoller/goexp/sets"
	"github.com/stretchr/testify/require"
)

func TestPop(t *testing.T) {
	x := sets.NewSet("apple", "banana", "cherry", "orange")
	y := x.Copy()
	popped := x.Pop().(string)
	subset := sets.NewSet(popped)

	log.Println(x)
	log.Println(y)
	log.Println(subset.Data()...)
	require.Equal(t, 3, len(x.Data()))
	require.Equal(t, true, subset.IsSubset(y))
}
