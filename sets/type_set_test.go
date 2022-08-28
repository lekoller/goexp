package sets_test

import (
	"reflect"
	"testing"

	"github.com/lekoller/goexp/sets"
	"github.com/stretchr/testify/require"
)

func TestNewSet(t *testing.T) {
	set := sets.NewSet("like", "dislike", "like", "subscribe")

	require.Equal(t, 3, len(set.Data))
	require.Equal(t, "like", set.Data[0])
	require.Equal(t, "dislike", set.Data[1])
	require.Equal(t, "subscribe", set.Data[2])
	require.Equal(t, reflect.TypeOf(""), set.Type)

	err := set.Add("juca")
	require.Nil(t, err)
	require.Equal(t, 4, len(set.Data))
	require.Equal(t, "juca", set.Data[3])

	err = set.Add(3)
	require.Error(t, err)
	require.EqualError(t, err, err.Error(), "invalid type")

	err = set.Add("juca")
	require.Error(t, err)
	require.EqualError(t, err, err.Error(), "not a new element")

	anotherSet := sets.NewSet[float32](2, 1.2, 7.9)

	require.Equal(t, 3, len(anotherSet.Data))
	require.Equal(t, float32(2), anotherSet.Data[0])
	require.Equal(t, float32(1.2), anotherSet.Data[1])
	require.Equal(t, float32(7.9), anotherSet.Data[2])
}
