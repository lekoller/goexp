package sets_test

import (
	"testing"

	"github.com/lekoller/goexp/sets"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	x := sets.NewSet("apple", "banana", "cherry")
	y := sets.NewSet("google", "microsoft", "apple")
	z := sets.NewSet("peach", "cherry")

	x.Update(y, z)

	require.Equal(t, 6, len(x))
}
