package span

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInt(t *testing.T) {
	spans, err := Int[int8](0, 0, 3)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{0, 0}}, spans)

	spans, err = Int[int8](1, 8, 3)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{1, 3}, {4, 6}, {7, 8}}, spans)

	spans, err = Int[int8](1, 8, 9)
	require.NoError(t, err)
	require.Equal(
		t,
		[]Span[int8]{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}},
		spans,
	)

	spans, err = Int[int8](math.MinInt8, math.MaxInt8, 1)
	require.NoError(t, err)
	require.Equal(
		t,
		[]Span[int8]{{math.MinInt8, math.MaxInt8}},
		spans,
	)

	spans, err = Int[int8](math.MinInt8, math.MaxInt8, 2)
	require.NoError(t, err)
	require.Equal(
		t,
		[]Span[int8]{{math.MinInt8, -1}, {0, math.MaxInt8}},
		spans,
	)
}

func TestIntError(t *testing.T) {
	spans, err := Int(2, 1, 1)
	require.Error(t, err)
	require.Equal(t, []Span[int](nil), spans)

	spans, err = Int(1, 2, -1)
	require.Error(t, err)
	require.Equal(t, []Span[int](nil), spans)

	spans, err = Int(1, 2, 0)
	require.Error(t, err)
	require.Equal(t, []Span[int](nil), spans)
}
