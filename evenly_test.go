package span

import (
	"math"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEvenlyIncreasing(t *testing.T) {
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

func TestEvenlyDecreasing(t *testing.T) {
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

func TestEvenlyQuantityExceedsLengthIncreasing(t *testing.T) {
	spans, err := Evenly[int8](1, 6, 3)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{1, 2}, {3, 4}, {5, 6}}, spans)
	require.Equal(t, 3, cap(spans))

	spans, err = Evenly[int8](1, 6, 4)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{1, 2}, {3, 4}, {5, 5}, {6, 6}}, spans)
	require.Equal(t, 4, cap(spans))

	spans, err = Evenly[int8](1, 6, 5)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{1, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}}, spans)
	require.Equal(t, 5, cap(spans))

	spans, err = Evenly[int8](1, 6, 6)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}}, spans)
	require.Equal(t, 6, cap(spans))

	spans, err = Evenly[int8](1, 6, 7)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}}, spans)
	require.Equal(t, 6, cap(spans))

	spans, err = Evenly[int8](1, 6, 8)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}}, spans)
	require.Equal(t, 6, cap(spans))

	spans, err = Evenly[int8](1, 6, 9)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}}, spans)
	require.Equal(t, 6, cap(spans))

	spans, err = Evenly[int8](1, 6, math.MaxInt8)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}}, spans)
	require.Equal(t, 6, cap(spans))
}

func TestEvenlyQuantityExceedsLengthDecreasing(t *testing.T) {
	spans, err := Evenly[int8](6, 1, 3)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{6, 5}, {4, 3}, {2, 1}}, spans)
	require.Equal(t, 3, cap(spans))

	spans, err = Evenly[int8](6, 1, 4)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{6, 5}, {4, 3}, {2, 2}, {1, 1}}, spans)
	require.Equal(t, 4, cap(spans))

	spans, err = Evenly[int8](6, 1, 5)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{6, 5}, {4, 4}, {3, 3}, {2, 2}, {1, 1}}, spans)
	require.Equal(t, 5, cap(spans))

	spans, err = Evenly[int8](6, 1, 6)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{6, 6}, {5, 5}, {4, 4}, {3, 3}, {2, 2}, {1, 1}}, spans)
	require.Equal(t, 6, cap(spans))

	spans, err = Evenly[int8](6, 1, 7)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{6, 6}, {5, 5}, {4, 4}, {3, 3}, {2, 2}, {1, 1}}, spans)
	require.Equal(t, 6, cap(spans))

	spans, err = Evenly[int8](6, 1, 8)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{6, 6}, {5, 5}, {4, 4}, {3, 3}, {2, 2}, {1, 1}}, spans)
	require.Equal(t, 6, cap(spans))

	spans, err = Evenly[int8](6, 1, 9)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{6, 6}, {5, 5}, {4, 4}, {3, 3}, {2, 2}, {1, 1}}, spans)
	require.Equal(t, 6, cap(spans))

	spans, err = Evenly[int8](6, 1, math.MaxInt8)
	require.NoError(t, err)
	require.Equal(t, []Span[int8]{{6, 6}, {5, 5}, {4, 4}, {3, 3}, {2, 2}, {1, 1}}, spans)
	require.Equal(t, 6, cap(spans))
}

func TestEvenlyError(t *testing.T) {
	spans, err := Evenly(1, 2, -1)
	require.Error(t, err)
	require.Equal(t, []Span[int](nil), spans)

	spans, err = Evenly(1, 2, 0)
	require.Error(t, err)
	require.Equal(t, []Span[int](nil), spans)
}

func TestEvenlyIsContinuousSigned(t *testing.T) {
	beginning, endings, quantities := testEvenlyRangeSigned()

	for _, begin := range beginning {
		for _, end := range endings {
			for _, quantity := range quantities {
				spans, err := Evenly(begin, end, quantity)
				require.NoError(
					t,
					err,
					"begin: %v, end: %v, quantity: %v",
					begin,
					end,
					quantity,
				)

				err = IsContinuous(spans)
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

func TestEvenlyIsContinuousUnsigned(t *testing.T) {
	beginning, endings, quantities := testEvenlyRangeUnsigned()

	for _, begin := range beginning {
		for _, end := range endings {
			for _, quantity := range quantities {
				spans, err := Evenly(begin, end, quantity)
				require.NoError(
					t,
					err,
					"begin: %v, end: %v, quantity: %v",
					begin,
					end,
					quantity,
				)

				err = IsContinuous(spans)
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

func TestEvenlyIsSortedSigned(t *testing.T) {
	beginning, endings, quantities := testEvenlyRangeSigned()

	for _, begin := range beginning {
		for _, end := range endings {
			for _, quantity := range quantities {
				spans, err := Evenly(begin, end, quantity)
				require.NoError(
					t,
					err,
					"begin: %v, end: %v, quantity: %v",
					begin,
					end,
					quantity,
				)

				cmp := CompareInc[int8]

				if begin > end {
					cmp = CompareDec[int8]
				}

				sorted := slices.IsSortedFunc(spans, cmp)
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

func TestEvenlyIsSortedUnsigned(t *testing.T) {
	beginning, endings, quantities := testEvenlyRangeUnsigned()

	for _, begin := range beginning {
		for _, end := range endings {
			for _, quantity := range quantities {
				spans, err := Evenly(begin, end, quantity)
				require.NoError(
					t,
					err,
					"begin: %v, end: %v, quantity: %v",
					begin,
					end,
					quantity,
				)

				cmp := CompareInc[uint8]

				if begin > end {
					cmp = CompareDec[uint8]
				}

				sorted := slices.IsSortedFunc(spans, cmp)
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

func TestEvenSigned(t *testing.T) {
	beginning, endings, quantities := testEvenlyRangeSigned()

	for _, begin := range beginning {
		for _, end := range endings {
			for _, quantity := range quantities {
				expected, err := Evenly(begin, end, quantity)
				require.NoError(
					t,
					err,
					"begin: %v, end: %v, quantity: %v",
					begin,
					end,
					quantity,
				)

				actual := make([]Span[int8], len(expected))

				for id, span := range Even(begin, end, quantity) {
					actual[id] = span
				}

				require.Equal(
					t,
					expected,
					actual,
					"begin: %v, end: %v, quantity: %v",
					begin,
					end,
					quantity,
				)
			}
		}
	}
}

func TestEvenUnsigned(t *testing.T) {
	beginning, endings, quantities := testEvenlyRangeUnsigned()

	for _, begin := range beginning {
		for _, end := range endings {
			for _, quantity := range quantities {
				expected, err := Evenly(begin, end, quantity)
				require.NoError(
					t,
					err,
					"begin: %v, end: %v, quantity: %v",
					begin,
					end,
					quantity,
				)

				actual := make([]Span[uint8], len(expected))

				for id, span := range Even(begin, end, quantity) {
					actual[id] = span
				}

				require.Equal(
					t,
					expected,
					actual,
					"begin: %v, end: %v, quantity: %v",
					begin,
					end,
					quantity,
				)
			}
		}
	}
}

func testEvenlyRangeSigned() ([]int8, []int8, []int8) {
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

	quantities := []int8{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		118, 119, 120, 121, 122, 123, 124, 125, 126, 127,
	}

	return beginning, endings, quantities
}

func testEvenlyRangeUnsigned() ([]uint8, []uint8, []uint8) {
	beginning := []uint8{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		246, 247, 248, 249, 250, 251, 252, 253, 254, 255,
	}

	endings := []uint8{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		246, 247, 248, 249, 250, 251, 252, 253, 254, 255,
	}

	quantities := []uint8{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		246, 247, 248, 249, 250, 251, 252, 253, 254, 255,
	}

	return beginning, endings, quantities
}

func TestEvenBreak(*testing.T) {
	for range Even(1, 2, 2) {
		break
	}

	for range Even(2, 1, 2) {
		break
	}
}

func TestEvenPanic(t *testing.T) {
	require.Panics(t,
		func() {
			for id := range Even(1, 2, -1) {
				_ = id
			}
		},
	)

	require.Panics(t,
		func() {
			for id := range Even(1, 2, 0) {
				_ = id
			}
		},
	)
}

func TestEvenSlices(t *testing.T) {
	testEvenSlices(t, []string(nil), [][]string(nil), 1)
	testEvenSlices(t, []string{}, [][]string(nil), 1)

	testEvenSlices(t, []string{"1"}, [][]string{{"1"}}, 1)
	testEvenSlices(t, []string{"1"}, [][]string{{"1"}}, 2)
	testEvenSlices(t, []string{"1"}, [][]string{{"1"}}, 3)
	testEvenSlices(t, []string{"1"}, [][]string{{"1"}}, 4)
	testEvenSlices(t, []string{"1"}, [][]string{{"1"}}, 5)
	testEvenSlices(t, []string{"1"}, [][]string{{"1"}}, 6)

	testEvenSlices(t, []string{"1", "2"}, [][]string{{"1", "2"}}, 1)
	testEvenSlices(t, []string{"1", "2"}, [][]string{{"1"}, {"2"}}, 2)
	testEvenSlices(t, []string{"1", "2"}, [][]string{{"1"}, {"2"}}, 3)
	testEvenSlices(t, []string{"1", "2"}, [][]string{{"1"}, {"2"}}, 4)
	testEvenSlices(t, []string{"1", "2"}, [][]string{{"1"}, {"2"}}, 5)
	testEvenSlices(t, []string{"1", "2"}, [][]string{{"1"}, {"2"}}, 6)

	testEvenSlices(t, []string{"1", "2", "3"}, [][]string{{"1", "2", "3"}}, 1)
	testEvenSlices(t, []string{"1", "2", "3"}, [][]string{{"1", "2"}, {"3"}}, 2)
	testEvenSlices(t, []string{"1", "2", "3"}, [][]string{{"1"}, {"2"}, {"3"}}, 3)
	testEvenSlices(t, []string{"1", "2", "3"}, [][]string{{"1"}, {"2"}, {"3"}}, 4)
	testEvenSlices(t, []string{"1", "2", "3"}, [][]string{{"1"}, {"2"}, {"3"}}, 5)
	testEvenSlices(t, []string{"1", "2", "3"}, [][]string{{"1"}, {"2"}, {"3"}}, 6)
}

func testEvenSlices(t *testing.T, divisible []string, expected [][]string, quantity int) {
	var actual [][]string //nolint:prealloc // To correctly reflect the actual behavior

	for sub := range EvenSlices(divisible, quantity) {
		actual = append(actual, sub)
	}

	require.Equal(t, expected, actual)
}

func TestEvenSlicesPanic(t *testing.T) {
	require.Panics(t,
		func() {
			for sub := range EvenSlices([]string{}, -1) {
				_ = sub
			}
		},
	)

	require.Panics(t,
		func() {
			for sub := range EvenSlices([]string{}, 0) {
				_ = sub
			}
		},
	)
}

func BenchmarkEvenly(b *testing.B) {
	expected := []Span[int]{{1, 1}, {2, 2}}

	var (
		spans []Span[int]
		err   error
	)

	for b.Loop() {
		spans, err = Evenly(1, 2, 2)
	}

	require.NoError(b, err)
	require.Equal(b, expected, spans)
}

func BenchmarkEven(b *testing.B) {
	expected := []Span[int]{{1, 1}, {2, 2}}

	var spans []Span[int]

	for b.Loop() {
		spans = make([]Span[int], 0, 2)

		for _, span := range Even(1, 2, 2) {
			spans = append(spans, span)
		}
	}

	require.Equal(b, expected, spans)
}

func BenchmarkEvenNoRealloc(b *testing.B) {
	expected := []Span[int]{{1, 1}, {2, 2}}

	spans := make([]Span[int], 0, 2)

	for b.Loop() {
		spans = spans[:0]

		for _, span := range Even(1, 2, 2) {
			spans = append(spans, span)
		}
	}

	require.Equal(b, expected, spans)
}

func BenchmarkEvenSlices(b *testing.B) {
	divisible := []string{"1", "2", "3", "4"}
	expected := [][]string{{"1", "2"}, {"3", "4"}}

	divided := make([][]string, 0, 2)

	for b.Loop() {
		divided = divided[:0]

		for sub := range EvenSlices(divisible, 2) {
			divided = append(divided, sub)
		}
	}

	require.Equal(b, expected, divided)
}
