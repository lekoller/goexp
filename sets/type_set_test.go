package sets_test

import (
	"testing"

	"github.com/lekoller/goexp/sets"
	"github.com/stretchr/testify/require"
)

func TestNewSet(t *testing.T) {
	set := sets.NewSet("like", "dislike", "like", "subscribe")

	require.Equal(t, 3, len(set))
	require.Equal(t, "like", set[0])
	require.Equal(t, "dislike", set[1])
	require.Equal(t, "subscribe", set[2])

	err := set.Add("juca")
	require.Nil(t, err)
	require.Equal(t, 4, len(set))
	require.Equal(t, "juca", set[3])

	err = set.Add(3)
	require.Nil(t, err)

	err = set.Add("juca")
	require.Error(t, err)
	require.EqualError(t, err, err.Error(), "not a new element")

	anotherSet := sets.NewSet[float32](2, 1.2, 7.9)

	require.Equal(t, 3, len(anotherSet))
	require.Equal(t, float32(2), anotherSet[0])
	require.Equal(t, float32(1.2), anotherSet[1])
	require.Equal(t, float32(7.9), anotherSet[2])
}
