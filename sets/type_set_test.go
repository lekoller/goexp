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
}
