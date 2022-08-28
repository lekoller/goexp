package sets_test

import (
	"log"
	"testing"

	"github.com/lekoller/goexp/sets"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	x := sets.NewSet("apple", "banana", "cherry")
	y := sets.NewSet("google", "microsoft", "apple")
	z := sets.NewSet("peach", "cherry")

	x.Update(y, z)

	log.Println(x.Data)
	require.Equal(t, 6, len(x.Data))
}
