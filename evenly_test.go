package span

import (
	"math"
	"slices"
	"testing"

	"github.com/akramarenkov/safe"
	"github.com/stretchr/testify/require"
)

func TestEvenly(t *testing.T) {
	testEvenlyIncreasing(t)
	testEvenlyDecreasing(t)
}

func testEvenlyIncreasing(t *testing.T) {
	spans, err := Evenly[int8](0, 0, 3)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{0, 0}}, spans)

	spans, err = Evenly[int8](1, 6, 3)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{1, 2}, {3, 4}, {5, 6}}, spans)

	spans, err = Evenly[int8](1, 7, 3)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{1, 3}, {4, 5}, {6, 7}}, spans)

	spans, err = Evenly[int8](1, 8, 3)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{1, 3}, {4, 6}, {7, 8}}, spans)

	spans, err = Evenly[int8](1, 9, 3)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{1, 3}, {4, 6}, {7, 9}}, spans)

	spans, err = Evenly[int8](1, 8, 9)
	require.NoError(t, err)
	require.Equal(
		t,
		[]Span[int8]{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}},
		spans,
	)

	spans, err = Evenly[int8](1, 8, 15)
	require.NoError(t, err)
	require.Equal(
		t,
		[]Span[int8]{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}},
		spans,
	)

	spans, err = Evenly[int8](math.MinInt8, math.MaxInt8, 1)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{math.MinInt8, math.MaxInt8}}, spans)

	spans, err = Evenly[int8](math.MinInt8, math.MaxInt8, 2)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{math.MinInt8, -1}, {0, math.MaxInt8}}, spans)
}

func testEvenlyDecreasing(t *testing.T) {
	spans, err := Evenly[int8](6, 1, 3)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{6, 5}, {4, 3}, {2, 1}}, spans)

	spans, err = Evenly[int8](7, 1, 3)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{7, 5}, {4, 3}, {2, 1}}, spans)

	spans, err = Evenly[int8](8, 1, 3)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{8, 6}, {5, 3}, {2, 1}}, spans)

	spans, err = Evenly[int8](9, 1, 3)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{9, 7}, {6, 4}, {3, 1}}, spans)

	spans, err = Evenly[int8](8, 1, 9)
	require.NoError(t, err)
	require.Equal(
		t,
		[]Span[int8]{{8, 8}, {7, 7}, {6, 6}, {5, 5}, {4, 4}, {3, 3}, {2, 2}, {1, 1}},
		spans,
	)

	spans, err = Evenly[int8](8, 1, 15)
	require.NoError(t, err)
	require.Equal(
		t,
		[]Span[int8]{{8, 8}, {7, 7}, {6, 6}, {5, 5}, {4, 4}, {3, 3}, {2, 2}, {1, 1}},
		spans,
	)

	spans, err = Evenly[int8](math.MaxInt8, math.MinInt8, 1)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{math.MaxInt8, math.MinInt8}}, spans)

	spans, err = Evenly[int8](math.MaxInt8, math.MinInt8, 2)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{math.MaxInt8, 0}, {-1, math.MinInt8}}, spans)
}

func TestEvenlyError(t *testing.T) {
	spans, err := Evenly(1, 2, -1)
	require.Error(t, err)
	require.Equal(t, []Span[int](nil), spans)

	spans, err = Evenly(1, 2, 0)
	require.Error(t, err)
	require.Equal(t, []Span[int](nil), spans)
}

func TestEvenlyInspect(t *testing.T) {
	for begin := range safe.Inc[int8](math.MinInt8, math.MaxInt8) {
		for end := range safe.Inc[int8](math.MinInt8, math.MaxInt8) {
			for quantity := range safe.Inc[int8](1, math.MaxInt8) {
				spans, err := Evenly(begin, end, quantity)
				if err != nil {
					require.NoError(
						t,
						err,
						"begin: %v, end: %v, quantity: %v",
						begin,
						end,
						quantity,
					)
				}

				// conditions are duplicated to performance reasons
				if err := IsContinuous(spans); err != nil {
					require.NoError(
						t,
						err,
						"begin: %v, end: %v, quantity: %v",
						begin,
						end,
						quantity,
					)
				}
			}
		}
	}

	for begin := range safe.Inc[uint8](0, math.MaxUint8) {
		for end := range safe.Inc[uint8](0, math.MaxUint8) {
			for quantity := range safe.Inc[uint8](1, math.MaxUint8) {
				spans, err := Evenly(begin, end, quantity)
				if err != nil {
					require.NoError(
						t,
						err,
						"begin: %v, end: %v, quantity: %v",
						begin,
						end,
						quantity,
					)
				}

				if err := IsContinuous(spans); err != nil {
					require.NoError(
						t,
						err,
						"begin: %v, end: %v, quantity: %v",
						begin,
						end,
						quantity,
					)
				}
			}
		}
	}
}

func TestEvenlyIsSorted(t *testing.T) {
	for begin := range safe.Inc[int8](math.MinInt8, math.MaxInt8) {
		for end := range safe.Inc[int8](math.MinInt8, math.MaxInt8) {
			for quantity := range safe.Inc[int8](1, math.MaxInt8) {
				spans, err := Evenly(begin, end, quantity)
				if err != nil {
					require.NoError(
						t,
						err,
						"begin: %v, end: %v, quantity: %v",
						begin,
						end,
						quantity,
					)
				}

				if sorted := slices.IsSortedFunc(spans, Compare); !sorted {
					require.True(
						t,
						sorted,
						"begin: %v, end: %v, quantity: %v",
						begin,
						end,
						quantity,
					)
				}
			}
		}
	}
}

func TestEvenlyIsSortedSelective(t *testing.T) {
	for begin := range safe.Inc[int8](math.MinInt8, math.MaxInt8) {
		for end := range safe.Inc[int8](math.MinInt8, math.MaxInt8) {
			for quantity := range safe.Inc[int8](1, math.MaxInt8) {
				spans, err := Evenly(begin, end, quantity)
				if err != nil {
					require.NoError(
						t,
						err,
						"begin: %v, end: %v, quantity: %v",
						begin,
						end,
						quantity,
					)
				}

				cmp := CompareInc[int8]
				if begin > end {
					cmp = CompareDec[int8]
				}

				if sorted := slices.IsSortedFunc(spans, cmp); !sorted {
					require.True(
						t,
						sorted,
						"begin: %v, end: %v, quantity: %v",
						begin,
						end,
						quantity,
					)
				}
			}
		}
	}
}

func BenchmarkEvenly(b *testing.B) {
	expected := []Span[int]{{1, 1}, {2, 2}}

	var (
		spans []Span[int]
		err   error
	)

	for range b.N {
		spans, err = Evenly(1, 2, 2)
	}

	require.NoError(b, err)
	require.Equal(b, expected, spans)
}
