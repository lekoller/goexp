package sets_test

import (
	"testing"

	"github.com/lekoller/goexp/sets"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	set := sets.Set{}
	err := set.Add("juca")

	require.Nil(t, err)
	require.Equal(t, 1, len(set))
	require.Equal(t, "juca", set[0])

	set.Add("juca")

	require.Equal(t, 1, len(set))
	require.Equal(t, "juca", set[0])

	set.Add("carlos")

	require.Equal(t, 2, len(set))
	require.Contains(t, set, "carlos")
	require.Contains(t, set, "juca")
}
