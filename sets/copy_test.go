package sets_test

import (
	"testing"

	"github.com/lekoller/goexp/sets"
	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	x := sets.NewSet("apple", "banana", "cherry")
	y := x.Copy()

	require.Equal(t, 3, len(y.Data))
}
