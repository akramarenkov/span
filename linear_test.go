package span

import (
	"math"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLinearIncreasing(t *testing.T) {
	spans, err := Linear[int8](2, 5, 1)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{2, 2}, {3, 3}, {4, 4}, {5, 5}}, spans)

	spans, err = Linear[int8](2, 2, 5)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{2, 2}}, spans)

	spans, err = Linear[int8](2, 6, 5)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{2, 6}}, spans)

	spans, err = Linear[int8](2, 16, 5)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{2, 6}, {7, 11}, {12, 16}}, spans)

	spans, err = Linear[int8](2, 17, 5)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{2, 6}, {7, 11}, {12, 16}, {17, 17}}, spans)

	spans, err = Linear[int8](2, 18, 5)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{2, 6}, {7, 11}, {12, 16}, {17, 18}}, spans)

	spans, err = Linear[int8](2, 19, 5)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{2, 6}, {7, 11}, {12, 16}, {17, 19}}, spans)

	spans, err = Linear[int8](math.MaxInt8-14, math.MaxInt8, 5)
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

	spans, err = Linear[int8](math.MaxInt8-15, math.MaxInt8, 5)
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

	spans, err = Linear[int8](math.MaxInt8-16, math.MaxInt8, 5)
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

func TestLinearDecreasing(t *testing.T) {
	spans, err := Linear[int8](5, 2, 1)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{5, 5}, {4, 4}, {3, 3}, {2, 2}}, spans)

	spans, err = Linear[int8](6, 2, 5)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{6, 2}}, spans)

	spans, err = Linear[int8](16, 2, 5)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{16, 12}, {11, 7}, {6, 2}}, spans)

	spans, err = Linear[int8](17, 2, 5)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{17, 13}, {12, 8}, {7, 3}, {2, 2}}, spans)

	spans, err = Linear[int8](18, 2, 5)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{18, 14}, {13, 9}, {8, 4}, {3, 2}}, spans)

	spans, err = Linear[int8](19, 2, 5)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{19, 15}, {14, 10}, {9, 5}, {4, 2}}, spans)

	spans, err = Linear[int8](math.MinInt8+14, math.MinInt8, 5)
	require.NoError(t, err)
	require.Equal(
		t,
		[]Span[int8]{
			{math.MinInt8 + 14, math.MinInt8 + 10},
			{math.MinInt8 + 9, math.MinInt8 + 5},
			{math.MinInt8 + 4, math.MinInt8},
		},
		spans,
	)

	spans, err = Linear[int8](math.MinInt8+15, math.MinInt8, 5)
	require.NoError(t, err)
	require.Equal(
		t,
		[]Span[int8]{
			{math.MinInt8 + 15, math.MinInt8 + 11},
			{math.MinInt8 + 10, math.MinInt8 + 6},
			{math.MinInt8 + 5, math.MinInt8 + 1},
			{math.MinInt8, math.MinInt8}},
		spans,
	)

	spans, err = Linear[int8](math.MinInt8+16, math.MinInt8, 5)
	require.NoError(t, err)
	require.Equal(
		t,
		[]Span[int8]{
			{math.MinInt8 + 16, math.MinInt8 + 12},
			{math.MinInt8 + 11, math.MinInt8 + 7},
			{math.MinInt8 + 6, math.MinInt8 + 2},
			{math.MinInt8 + 1, math.MinInt8}},
		spans,
	)
}

func TestLinearError(t *testing.T) {
	spans, err := Linear(1, 2, -1)
	require.Error(t, err)
	require.Equal(t, []Span[int](nil), spans)

	spans, err = Linear(1, 2, 0)
	require.Error(t, err)
	require.Equal(t, []Span[int](nil), spans)
}

func TestLinearIsContinuousSigned(t *testing.T) {
	beginning, endings, widths := testLinearRangeSigned()

	for _, begin := range beginning {
		for _, end := range endings {
			for _, width := range widths {
				spans, err := Linear(begin, end, width)
				require.NoError(
					t,
					err,
					"begin: %v, end: %v, width: %v",
					begin,
					end,
					width,
				)

				err = IsContinuous(spans)
				require.NoError(
					t,
					err,
					"begin: %v, end: %v, width: %v",
					begin,
					end,
					width,
				)
			}
		}
	}
}

func TestLinearIsContinuousUnsigned(t *testing.T) {
	beginning, endings, widths := testLinearRangeUnsigned()

	for _, begin := range beginning {
		for _, end := range endings {
			for _, width := range widths {
				spans, err := Linear(begin, end, width)
				require.NoError(
					t,
					err,
					"begin: %v, end: %v, width: %v",
					begin,
					end,
					width,
				)

				err = IsContinuous(spans)
				require.NoError(
					t,
					err,
					"begin: %v, end: %v, width: %v",
					begin,
					end,
					width,
				)
			}
		}
	}
}

func TestLinearIsSortedSigned(t *testing.T) {
	beginning, endings, widths := testLinearRangeSigned()

	for _, begin := range beginning {
		for _, end := range endings {
			for _, width := range widths {
				spans, err := Linear(begin, end, width)
				require.NoError(
					t,
					err,
					"begin: %v, end: %v, width: %v",
					begin,
					end,
					width,
				)

				cmp := CompareInc[int8]

				if begin > end {
					cmp = CompareDec[int8]
				}

				sorted := slices.IsSortedFunc(spans, cmp)
				require.True(
					t,
					sorted,
					"begin: %v, end: %v, width: %v",
					begin,
					end,
					width,
				)
			}
		}
	}
}

func TestLinearIsSortedUnsigned(t *testing.T) {
	beginning, endings, widths := testLinearRangeUnsigned()

	for _, begin := range beginning {
		for _, end := range endings {
			for _, width := range widths {
				spans, err := Linear(begin, end, width)
				require.NoError(
					t,
					err,
					"begin: %v, end: %v, width: %v",
					begin,
					end,
					width,
				)

				cmp := CompareInc[uint8]

				if begin > end {
					cmp = CompareDec[uint8]
				}

				sorted := slices.IsSortedFunc(spans, cmp)
				require.True(
					t,
					sorted,
					"begin: %v, end: %v, width: %v",
					begin,
					end,
					width,
				)
			}
		}
	}
}

func testLinearRangeSigned() ([]int8, []int8, []int8) {
	beginning := []int8{
		-128, -127, -126, -125, -124, -123, -122, -121, -120, -119,
		-10, -9, -8, -7, -6, -5, -4, -3, -2, -1,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		118, 119, 120, 121, 122, 123, 124, 125, 126, 127,
	}

	endings := []int8{
		-128, -127, -126, -125, -124, -123, -122, -121, -120, -119,
		-10, -9, -8, -7, -6, -5, -4, -3, -2, -1,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		118, 119, 120, 121, 122, 123, 124, 125, 126, 127,
	}

	widths := []int8{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		118, 119, 120, 121, 122, 123, 124, 125, 126, 127,
	}

	return beginning, endings, widths
}

func testLinearRangeUnsigned() ([]uint8, []uint8, []uint8) {
	beginning := []uint8{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		246, 247, 248, 249, 250, 251, 252, 253, 254, 255,
	}

	endings := []uint8{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		246, 247, 248, 249, 250, 251, 252, 253, 254, 255,
	}

	widths := []uint8{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		246, 247, 248, 249, 250, 251, 252, 253, 254, 255,
	}

	return beginning, endings, widths
}

func BenchmarkLinear(b *testing.B) {
	expected := []Span[int]{{1, 1}, {2, 2}}

	var (
		spans []Span[int]
		err   error
	)

	for b.Loop() {
		spans, err = Linear(1, 2, 1)
	}

	require.NoError(b, err)
	require.Equal(b, expected, spans)
}
