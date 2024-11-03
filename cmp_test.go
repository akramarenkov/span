package span

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompareInc(t *testing.T) {
	expected := []Span[int]{{2, 6}, {7, 11}, {12, 16}}
	actual := []Span[int]{{12, 16}, {2, 6}, {7, 11}}
	intersect := []Span[int]{{1, 2}, {2, 3}, {1, 2}}

	require.NotEqual(t, expected, actual)

	slices.SortFunc(actual, CompareInc)
	require.Equal(t, expected, actual)

	require.Panics(t, func() { slices.SortFunc(intersect, CompareInc) })
}

func TestCompareDec(t *testing.T) {
	expected := []Span[int]{{16, 12}, {11, 7}, {6, 2}}
	actual := []Span[int]{{6, 2}, {16, 12}, {11, 7}}
	intersect := []Span[int]{{2, 1}, {3, 2}, {4, 2}}

	require.NotEqual(t, expected, actual)

	slices.SortFunc(actual, CompareDec)
	require.Equal(t, expected, actual)

	require.Panics(t, func() { slices.SortFunc(intersect, CompareDec) })
}

func TestCompare(t *testing.T) {
	expectedInc := []Span[int]{{2, 6}, {7, 11}, {12, 16}}
	actualInc := []Span[int]{{12, 16}, {2, 6}, {7, 11}}

	expectedDec := []Span[int]{{16, 12}, {11, 7}, {6, 2}}
	actualDec := []Span[int]{{6, 2}, {16, 12}, {11, 7}}

	intersectInc := []Span[int]{{1, 2}, {2, 3}, {1, 2}}
	intersectDec := []Span[int]{{2, 1}, {3, 2}, {4, 2}}
	difference := []Span[int]{{16, 6}, {7, 11}, {6, 2}}

	require.NotEqual(t, expectedInc, actualInc)
	require.NotEqual(t, expectedDec, actualDec)

	slices.SortFunc(actualInc, Compare)
	require.Equal(t, expectedInc, actualInc)

	slices.SortFunc(actualDec, Compare)
	require.Equal(t, expectedDec, actualDec)

	require.Panics(t, func() { slices.SortFunc(intersectInc, Compare) })
	require.Panics(t, func() { slices.SortFunc(intersectDec, Compare) })
	require.Panics(t, func() { slices.SortFunc(difference, Compare) })
}
