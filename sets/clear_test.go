package sets_test

import (
	"testing"

	"github.com/lekoller/goexp/sets"
	"github.com/stretchr/testify/require"
)

func TestClear(t *testing.T) {
	x := sets.NewSet("apple", "banana", "cherry", "orange")
	x.Clear()

	require.Equal(t, 0, len(x))
}
