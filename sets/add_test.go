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
	require.Equal(t, 1, len(set.Data))
	require.Equal(t, "juca", set.Data[0])

	set.Add("juca")

	require.Equal(t, 1, len(set.Data))
	require.Equal(t, "juca", set.Data[0])
}
