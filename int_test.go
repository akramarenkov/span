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

	spans, err = Int[int8](1, 6, 3)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{1, 2}, {3, 4}, {5, 6}}, spans)

	spans, err = Int[int8](1, 7, 3)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{1, 3}, {4, 5}, {6, 7}}, spans)

	spans, err = Int[int8](1, 8, 3)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{1, 3}, {4, 6}, {7, 8}}, spans)

	spans, err = Int[int8](1, 9, 3)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{1, 3}, {4, 6}, {7, 9}}, spans)

	spans, err = Int[int8](1, 8, 9)
	require.NoError(t, err)
	require.Equal(
		t,
		[]Span[int8]{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}},
		spans,
	)

	spans, err = Int[int8](1, 8, 15)
	require.NoError(t, err)
	require.Equal(
		t,
		[]Span[int8]{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}},
		spans,
	)

	spans, err = Int[int8](math.MinInt8, math.MaxInt8, 1)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{math.MinInt8, math.MaxInt8}}, spans)

	spans, err = Int[int8](math.MinInt8, math.MaxInt8, 2)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{math.MinInt8, -1}, {0, math.MaxInt8}}, spans)
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

func TestIntWidth(t *testing.T) {
	spans, err := IntWidth[int8](2, 5, 1)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{2, 2}, {3, 3}, {4, 4}, {5, 5}}, spans)

	spans, err = IntWidth[int8](2, 2, 5)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{2, 2}}, spans)

	spans, err = IntWidth[int8](2, 6, 5)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{2, 6}}, spans)

	spans, err = IntWidth[int8](2, 16, 5)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{2, 6}, {7, 11}, {12, 16}}, spans)

	spans, err = IntWidth[int8](2, 17, 5)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{2, 6}, {7, 11}, {12, 16}, {17, 17}}, spans)

	spans, err = IntWidth[int8](2, 18, 5)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{2, 6}, {7, 11}, {12, 16}, {17, 18}}, spans)

	spans, err = IntWidth[int8](2, 19, 5)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{2, 6}, {7, 11}, {12, 16}, {17, 19}}, spans)

	spans, err = IntWidth[int8](math.MaxInt8-14, math.MaxInt8, 5)
	require.NoError(t, err)
	require.Equal(
		t,
		[]Span[int8]{
			{math.MaxInt8 - 14, math.MaxInt8 - 10},
			{math.MaxInt8 - 9, math.MaxInt8 - 5},
			{math.MaxInt8 - 4, math.MaxInt8},
		},
		spans,
	)

	spans, err = IntWidth[int8](math.MaxInt8-15, math.MaxInt8, 5)
	require.NoError(t, err)
	require.Equal(
		t,
		[]Span[int8]{
			{math.MaxInt8 - 15, math.MaxInt8 - 11},
			{math.MaxInt8 - 10, math.MaxInt8 - 6},
			{math.MaxInt8 - 5, math.MaxInt8 - 1},
			{math.MaxInt8, math.MaxInt8}},
		spans,
	)

	spans, err = IntWidth[int8](math.MaxInt8-16, math.MaxInt8, 5)
	require.NoError(t, err)
	require.Equal(
		t,
		[]Span[int8]{
			{math.MaxInt8 - 16, math.MaxInt8 - 12},
			{math.MaxInt8 - 11, math.MaxInt8 - 7},
			{math.MaxInt8 - 6, math.MaxInt8 - 2},
			{math.MaxInt8 - 1, math.MaxInt8}},
		spans,
	)
}

func TestIntWidthError(t *testing.T) {
	spans, err := IntWidth(2, 1, 1)
	require.Error(t, err)
	require.Equal(t, []Span[int](nil), spans)

	spans, err = IntWidth(1, 2, -1)
	require.Error(t, err)
	require.Equal(t, []Span[int](nil), spans)

	spans, err = IntWidth(1, 2, 0)
	require.Error(t, err)
	require.Equal(t, []Span[int](nil), spans)
}

func BenchmarkInt(b *testing.B) {
	var (
		spans []Span[int]
		err   error
	)

	for range b.N {
		spans, err = Int(1, 2, 2)
	}

	require.NotNil(b, spans)
	require.NoError(b, err)
}

func BenchmarkIntWidth(b *testing.B) {
	var (
		spans []Span[int]
		err   error
	)

	for range b.N {
		spans, err = IntWidth(1, 2, 1)
	}

	require.NotNil(b, spans)
	require.NoError(b, err)
}
