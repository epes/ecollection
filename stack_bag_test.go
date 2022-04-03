package ecollection_test

import (
	"github.com/epes/ecollection"
	"testing"
)

import (
	"github.com/stretchr/testify/require"
)

func Test_Number_Stack_Bag(t *testing.T) {
	bag := ecollection.NewFullNumberStackBag[int](2)

	t0, err := bag.Take()
	require.NoError(t, err)
	require.EqualValues(t, 1, t0)

	t1, err := bag.Take()
	require.NoError(t, err)
	require.EqualValues(t, 0, t1)

	t2, err := bag.Take()
	require.Equal(t, ecollection.ErrEmpty, err)
	require.Zero(t, t2)

	err = bag.Put(t0)
	require.NoError(t, err)

	err = bag.Put(t1)
	require.NoError(t, err)

	err = bag.Put(t1)
	require.Equal(t, ecollection.ErrFull, err)

	t3, err := bag.Take()
	require.NoError(t, err)
	require.EqualValues(t, 0, t3)

	t4, err := bag.Take()
	require.NoError(t, err)
	require.EqualValues(t, 1, t4)
}
