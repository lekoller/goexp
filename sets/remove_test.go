package sets_test

import (
	"testing"

	"github.com/lekoller/goexp/sets"
	"github.com/stretchr/testify/require"
)

func TestRemove(t *testing.T) {
	x := sets.NewSet("apple", "banana", "cherry")

	x.Remove("apple")

	require.Equal(t, 2, len(x))

	err := x.Remove("apple")

	require.Error(t, err)
	require.Equal(t, 2, len(x))
}
