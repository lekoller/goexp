package sets_test

import (
	"testing"

	"github.com/lekoller/goexp/sets"
	"github.com/stretchr/testify/require"
)

func TestDiscard(t *testing.T) {
	x := sets.NewSet("apple", "banana", "cherry")

	x.Discard("apple")

	require.Equal(t, 2, len(x.Data()))
}
