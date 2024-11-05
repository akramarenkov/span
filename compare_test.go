package span

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompareInc(t *testing.T) {
	testCompareFunc(
		t,
		CompareInc,
		[]Span[int]{{1, 1}, {2, 2}, {3, 3}},
		[]Span[int]{{3, 3}, {1, 1}, {2, 2}},
	)

	testCompareFunc(
		t,
		CompareInc,
		[]Span[int]{{2, 6}, {7, 11}, {12, 16}},
		[]Span[int]{{12, 16}, {2, 6}, {7, 11}},
	)

	testCompareFunc(
		t,
		CompareInc,
		[]Span[int]{{2, 6}, {8, 12}, {13, 17}},
		[]Span[int]{{13, 17}, {2, 6}, {8, 12}},
	)

	testCompareFunc(
		t,
		CompareInc,
		[]Span[int]{{2, 6}, {8, 12}, {14, 18}},
		[]Span[int]{{14, 18}, {2, 6}, {8, 12}},
	)

	testCompareFunc(
		t,
		CompareInc,
		[]Span[int]{{2, 6}, {9, 13}, {16, 20}},
		[]Span[int]{{16, 20}, {2, 6}, {9, 13}},
	)

	intersect := []Span[int]{{2, 2}, {2, 2}, {2, 2}}
	differenceSequencing := []Span[int]{{16, 12}, {11, 7}, {6, 2}}

	require.Panics(t, func() { slices.SortFunc(intersect, CompareInc) })
	require.Panics(t, func() { slices.SortFunc(differenceSequencing, CompareInc) })
}

func TestCompareDec(t *testing.T) {
	testCompareFunc(
		t,
		CompareDec,
		[]Span[int]{{3, 3}, {2, 2}, {1, 1}},
		[]Span[int]{{2, 2}, {1, 1}, {3, 3}},
	)

	testCompareFunc(
		t,
		CompareDec,
		[]Span[int]{{16, 12}, {11, 7}, {6, 2}},
		[]Span[int]{{6, 2}, {16, 12}, {11, 7}},
	)

	testCompareFunc(
		t,
		CompareDec,
		[]Span[int]{{17, 13}, {12, 8}, {6, 2}},
		[]Span[int]{{6, 2}, {17, 13}, {12, 8}},
	)

	testCompareFunc(
		t,
		CompareDec,
		[]Span[int]{{18, 14}, {12, 8}, {6, 2}},
		[]Span[int]{{6, 2}, {18, 14}, {12, 8}},
	)

	testCompareFunc(
		t,
		CompareDec,
		[]Span[int]{{20, 16}, {13, 9}, {6, 2}},
		[]Span[int]{{6, 2}, {20, 16}, {13, 9}},
	)

	intersect := []Span[int]{{2, 2}, {2, 2}, {2, 2}}
	differenceSequencing := []Span[int]{{2, 6}, {7, 11}, {12, 16}}

	require.Panics(t, func() { slices.SortFunc(intersect, CompareDec) })
	require.Panics(t, func() { slices.SortFunc(differenceSequencing, CompareDec) })
}

func testCompareFunc(
	t *testing.T,
	compare func(Span[int], Span[int]) int,
	expected []Span[int],
	actual []Span[int],
) {
	require.NotEqual(t, expected, actual)
	slices.SortFunc(actual, compare)
	require.Equal(t, expected, actual)
}

func TestSearchInc(t *testing.T) {
	spans := []Span[int]{{1, 1}, {2, 3}, {4, 6}, {8, 11}}

	id, found := slices.BinarySearchFunc(spans, Span[int]{0, 0}, SearchInc)
	require.False(t, found)
	require.Equal(t, 0, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{0, 1}, SearchInc)
	require.False(t, found)
	require.Equal(t, 0, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{1, 1}, SearchInc)
	require.True(t, found)
	require.Equal(t, 0, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{1, 2}, SearchInc)
	require.False(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{0, 2}, SearchInc)
	require.False(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{0, 3}, SearchInc)
	require.False(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{-1, 3}, SearchInc)
	require.False(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{2, 2}, SearchInc)
	require.True(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{2, 3}, SearchInc)
	require.True(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{3, 3}, SearchInc)
	require.True(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{2, 4}, SearchInc)
	require.False(t, found)
	require.Equal(t, 2, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{1, 4}, SearchInc)
	require.False(t, found)
	require.Equal(t, 2, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{1, 5}, SearchInc)
	require.False(t, found)
	require.Equal(t, 2, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{0, 5}, SearchInc)
	require.False(t, found)
	require.Equal(t, 2, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{4, 4}, SearchInc)
	require.True(t, found)
	require.Equal(t, 2, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{4, 5}, SearchInc)
	require.True(t, found)
	require.Equal(t, 2, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{4, 6}, SearchInc)
	require.True(t, found)
	require.Equal(t, 2, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{5, 5}, SearchInc)
	require.True(t, found)
	require.Equal(t, 2, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{5, 6}, SearchInc)
	require.True(t, found)
	require.Equal(t, 2, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{6, 6}, SearchInc)
	require.True(t, found)
	require.Equal(t, 2, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{4, 7}, SearchInc)
	require.False(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{5, 7}, SearchInc)
	require.False(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{6, 7}, SearchInc)
	require.False(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{3, 4}, SearchInc)
	require.False(t, found)
	require.Equal(t, 2, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{3, 5}, SearchInc)
	require.False(t, found)
	require.Equal(t, 2, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{3, 6}, SearchInc)
	require.False(t, found)
	require.Equal(t, 2, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{3, 7}, SearchInc)
	require.False(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{7, 7}, SearchInc)
	require.False(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{8, 8}, SearchInc)
	require.True(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{8, 9}, SearchInc)
	require.True(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{8, 10}, SearchInc)
	require.True(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{8, 11}, SearchInc)
	require.True(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{9, 9}, SearchInc)
	require.True(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{9, 10}, SearchInc)
	require.True(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{9, 11}, SearchInc)
	require.True(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{10, 10}, SearchInc)
	require.True(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{10, 11}, SearchInc)
	require.True(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{11, 11}, SearchInc)
	require.True(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{8, 12}, SearchInc)
	require.False(t, found)
	require.Equal(t, 4, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{9, 12}, SearchInc)
	require.False(t, found)
	require.Equal(t, 4, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{10, 12}, SearchInc)
	require.False(t, found)
	require.Equal(t, 4, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{11, 12}, SearchInc)
	require.False(t, found)
	require.Equal(t, 4, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{7, 8}, SearchInc)
	require.False(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{7, 9}, SearchInc)
	require.False(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{7, 10}, SearchInc)
	require.False(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{7, 11}, SearchInc)
	require.False(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{7, 12}, SearchInc)
	require.False(t, found)
	require.Equal(t, 4, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{6, 8}, SearchInc)
	require.False(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{6, 9}, SearchInc)
	require.False(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{6, 10}, SearchInc)
	require.False(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{6, 11}, SearchInc)
	require.False(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{6, 12}, SearchInc)
	require.False(t, found)
	require.Equal(t, 4, id)
}

func TestSearchIncPanic(t *testing.T) {
	differenceSequencing := []Span[int]{{16, 12}, {11, 7}, {6, 2}}

	require.Panics(t, func() { slices.SortFunc(differenceSequencing, SearchInc) })
}

func TestSearchDec(t *testing.T) {
	spans := []Span[int]{{11, 8}, {6, 4}, {3, 2}, {1, 1}}

	id, found := slices.BinarySearchFunc(spans, Span[int]{0, 0}, SearchDec)
	require.False(t, found)
	require.Equal(t, 4, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{1, 0}, SearchDec)
	require.False(t, found)
	require.Equal(t, 4, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{1, 1}, SearchDec)
	require.True(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{2, 1}, SearchDec)
	require.False(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{2, 0}, SearchDec)
	require.False(t, found)
	require.Equal(t, 4, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{3, 0}, SearchDec)
	require.False(t, found)
	require.Equal(t, 4, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{3, -1}, SearchDec)
	require.False(t, found)
	require.Equal(t, 4, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{2, 2}, SearchDec)
	require.True(t, found)
	require.Equal(t, 2, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{3, 2}, SearchDec)
	require.True(t, found)
	require.Equal(t, 2, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{3, 3}, SearchDec)
	require.True(t, found)
	require.Equal(t, 2, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{4, 2}, SearchDec)
	require.False(t, found)
	require.Equal(t, 2, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{4, 1}, SearchDec)
	require.False(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{5, 1}, SearchDec)
	require.False(t, found)
	require.Equal(t, 3, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{5, 0}, SearchDec)
	require.False(t, found)
	require.Equal(t, 4, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{4, 4}, SearchDec)
	require.True(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{5, 4}, SearchDec)
	require.True(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{5, 4}, SearchDec)
	require.True(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{5, 5}, SearchDec)
	require.True(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{6, 5}, SearchDec)
	require.True(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{6, 6}, SearchDec)
	require.True(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{7, 4}, SearchDec)
	require.False(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{7, 5}, SearchDec)
	require.False(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{7, 6}, SearchDec)
	require.False(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{4, 3}, SearchDec)
	require.False(t, found)
	require.Equal(t, 2, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{5, 3}, SearchDec)
	require.False(t, found)
	require.Equal(t, 2, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{6, 3}, SearchDec)
	require.False(t, found)
	require.Equal(t, 2, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{7, 3}, SearchDec)
	require.False(t, found)
	require.Equal(t, 2, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{7, 7}, SearchDec)
	require.False(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{8, 8}, SearchDec)
	require.True(t, found)
	require.Equal(t, 0, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{9, 8}, SearchDec)
	require.True(t, found)
	require.Equal(t, 0, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{10, 8}, SearchDec)
	require.True(t, found)
	require.Equal(t, 0, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{11, 8}, SearchDec)
	require.True(t, found)
	require.Equal(t, 0, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{9, 9}, SearchDec)
	require.True(t, found)
	require.Equal(t, 0, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{10, 9}, SearchDec)
	require.True(t, found)
	require.Equal(t, 0, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{11, 9}, SearchDec)
	require.True(t, found)
	require.Equal(t, 0, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{10, 10}, SearchDec)
	require.True(t, found)
	require.Equal(t, 0, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{11, 10}, SearchDec)
	require.True(t, found)
	require.Equal(t, 0, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{11, 11}, SearchDec)
	require.True(t, found)
	require.Equal(t, 0, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{12, 8}, SearchDec)
	require.False(t, found)
	require.Equal(t, 0, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{12, 9}, SearchDec)
	require.False(t, found)
	require.Equal(t, 0, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{12, 10}, SearchDec)
	require.False(t, found)
	require.Equal(t, 0, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{12, 11}, SearchDec)
	require.False(t, found)
	require.Equal(t, 0, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{8, 7}, SearchDec)
	require.False(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{9, 7}, SearchDec)
	require.False(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{10, 7}, SearchDec)
	require.False(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{11, 7}, SearchDec)
	require.False(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{12, 7}, SearchDec)
	require.False(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{8, 6}, SearchDec)
	require.False(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{9, 6}, SearchDec)
	require.False(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{10, 6}, SearchDec)
	require.False(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{11, 6}, SearchDec)
	require.False(t, found)
	require.Equal(t, 1, id)

	id, found = slices.BinarySearchFunc(spans, Span[int]{12, 6}, SearchDec)
	require.False(t, found)
	require.Equal(t, 1, id)
}

func TestSearchDecPanic(t *testing.T) {
	differenceSequencing := []Span[int]{{2, 6}, {7, 11}, {12, 16}}

	require.Panics(t, func() { slices.SortFunc(differenceSequencing, SearchDec) })
}
