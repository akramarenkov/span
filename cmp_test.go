package span

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompareInc(t *testing.T) {
	expected := []Span[int]{{2, 6}, {7, 11}, {12, 16}}
	actual := []Span[int]{{12, 16}, {2, 6}, {7, 11}}

	slices.SortFunc(actual, CompareInc)
	require.Equal(t, expected, actual)

	intersect := []Span[int]{{2, 2}, {2, 2}, {2, 2}}
	differenceSequencing := []Span[int]{{16, 12}, {11, 7}, {6, 2}}

	require.Panics(t, func() { slices.SortFunc(intersect, CompareInc) })
	require.Panics(t, func() { slices.SortFunc(differenceSequencing, CompareInc) })
}

func TestCompareDec(t *testing.T) {
	expected := []Span[int]{{16, 12}, {11, 7}, {6, 2}}
	actual := []Span[int]{{6, 2}, {16, 12}, {11, 7}}

	slices.SortFunc(actual, CompareDec)
	require.Equal(t, expected, actual)

	intersect := []Span[int]{{2, 2}, {2, 2}, {2, 2}}
	differenceSequencing := []Span[int]{{2, 6}, {7, 11}, {12, 16}}

	require.Panics(t, func() { slices.SortFunc(intersect, CompareDec) })
	require.Panics(t, func() { slices.SortFunc(differenceSequencing, CompareDec) })
}

func TestCompare(t *testing.T) {
	testCompareIncreasing(t)
	testCompareDecreasing(t)
	testCompareUndetermined(t)
}

func testCompareIncreasing(t *testing.T) {
	expected := []Span[int]{{2, 6}, {7, 11}, {12, 16}}
	actual := []Span[int]{{12, 16}, {2, 6}, {7, 11}}

	slices.SortFunc(actual, Compare)
	require.Equal(t, expected, actual)
}

func testCompareDecreasing(t *testing.T) {
	expected := []Span[int]{{16, 12}, {11, 7}, {6, 2}}
	actual := []Span[int]{{6, 2}, {16, 12}, {11, 7}}

	slices.SortFunc(actual, Compare)
	require.Equal(t, expected, actual)
}

func testCompareUndetermined(t *testing.T) {
	expected := []Span[int]{{-127, -127}, {-128, -128}, {-126, -126}}
	actual := []Span[int]{{-127, -127}, {-128, -128}, {-126, -126}}

	slices.SortFunc(actual, Compare)
	require.Equal(t, expected, actual)
}

func TestComparePanic(t *testing.T) {
	intersectIncreasing := []Span[int]{{1, 2}, {2, 3}, {1, 2}}
	intersectDecreasing := []Span[int]{{2, 1}, {3, 2}, {4, 2}}
	differenceSequencing := []Span[int]{{16, 6}, {7, 11}, {6, 2}}

	require.Panics(t, func() { slices.SortFunc(intersectIncreasing, Compare) })
	require.Panics(t, func() { slices.SortFunc(intersectDecreasing, Compare) })
	require.Panics(t, func() { slices.SortFunc(differenceSequencing, Compare) })
}
