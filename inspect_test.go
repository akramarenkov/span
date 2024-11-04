package span

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsNotDiffSequencing(t *testing.T) {
	correct := []Span[int]{{2, 6}, {7, 11}, {12, 16}, {17, 18}}
	difference := []Span[int]{{2, 6}, {7, 11}, {12, 16}, {18, 17}}

	require.NoError(t, IsNotDiffSequencing[int](nil))
	require.NoError(t, IsNotDiffSequencing([]Span[int]{}))
	require.NoError(t, IsNotDiffSequencing([]Span[int]{{2, 6}}))
	require.NoError(t, IsNotDiffSequencing([]Span[int]{{2, 6}, {7, 11}}))
	require.NoError(t, IsNotDiffSequencing(correct))

	require.Error(t, IsNotDiffSequencing([]Span[int]{{2, 6}, {11, 7}}))
	require.Error(t, IsNotDiffSequencing(difference))
}

func TestIsTwoSpansNotDiffSequencing(t *testing.T) {
	require.NoError(t, isTwoSpansNotDiffSequencing(Span[int]{1, 2}, Span[int]{3, 4}))
	require.NoError(t, isTwoSpansNotDiffSequencing(Span[int]{2, 1}, Span[int]{4, 3}))

	require.Error(t, isTwoSpansNotDiffSequencing(Span[int]{1, 2}, Span[int]{4, 3}))
	require.Error(t, isTwoSpansNotDiffSequencing(Span[int]{2, 1}, Span[int]{3, 4}))
}

func TestIsIncreasing(t *testing.T) {
	increasing := []Span[int]{{2, 6}, {7, 11}, {12, 16}, {17, 18}}
	nonincreasing := []Span[int]{{2, 6}, {11, 7}, {12, 16}, {17, 18}}

	require.NoError(t, IsIncreasing([]Span[int]{{2, 6}}))
	require.NoError(t, IsIncreasing(increasing))

	require.Error(t, IsIncreasing[int](nil))
	require.Error(t, IsIncreasing([]Span[int]{}))
	require.Error(t, IsIncreasing([]Span[int]{{6, 2}}))
	require.Error(t, IsIncreasing(nonincreasing))
}

func TestIsDecreasing(t *testing.T) {
	decreasing := []Span[int]{{18, 17}, {16, 12}, {11, 7}, {6, 2}}
	nondecreasing := []Span[int]{{18, 17}, {16, 12}, {7, 11}, {6, 2}}

	require.NoError(t, IsDecreasing([]Span[int]{{6, 2}}))
	require.NoError(t, IsDecreasing(decreasing))

	require.Error(t, IsDecreasing[int](nil))
	require.Error(t, IsDecreasing([]Span[int]{}))
	require.Error(t, IsDecreasing([]Span[int]{{2, 6}}))
	require.Error(t, IsDecreasing(nondecreasing))
}

func TestIsNotIntersect(t *testing.T) {
	correct := []Span[int]{{2, 6}, {7, 11}, {12, 16}, {17, 18}}
	intersected := []Span[int]{{2, 6}, {7, 11}, {12, 16}, {16, 18}}

	require.NoError(t, IsNotIntersect[int](nil))
	require.NoError(t, IsNotIntersect([]Span[int]{}))
	require.NoError(t, IsNotIntersect([]Span[int]{{2, 6}}))
	require.NoError(t, IsNotIntersect([]Span[int]{{2, 6}, {7, 11}}))
	require.NoError(t, IsNotIntersect(correct))

	require.Error(t, IsNotIntersect([]Span[int]{{2, 6}, {6, 11}}))
	require.Error(t, IsNotIntersect(intersected))

	require.Error(t, IsNotIntersect([]Span[int]{{2, 6}, {11, 7}}))
}

func TestIsTwoSpansNotIntersect(t *testing.T) {
	require.NoError(t, isTwoSpansNotIntersect(Span[int]{1, 2}, Span[int]{3, 4}))
	require.NoError(t, isTwoSpansNotIntersect(Span[int]{3, 4}, Span[int]{1, 2}))
	require.NoError(t, isTwoSpansNotIntersect(Span[int]{2, 2}, Span[int]{3, 4}))
	require.NoError(t, isTwoSpansNotIntersect(Span[int]{3, 4}, Span[int]{2, 2}))
	require.NoError(t, isTwoSpansNotIntersect(Span[int]{2, 2}, Span[int]{3, 3}))
	require.NoError(t, isTwoSpansNotIntersect(Span[int]{3, 3}, Span[int]{2, 2}))

	require.Error(t, isTwoSpansNotIntersect(Span[int]{1, 2}, Span[int]{2, 4}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{2, 4}, Span[int]{1, 2}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{1, 3}, Span[int]{2, 4}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{2, 4}, Span[int]{1, 3}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{1, 4}, Span[int]{2, 4}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{2, 4}, Span[int]{1, 4}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{1, 5}, Span[int]{2, 4}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{2, 4}, Span[int]{1, 5}))

	require.NoError(t, isTwoSpansNotIntersect(Span[int]{2, 1}, Span[int]{4, 3}))
	require.NoError(t, isTwoSpansNotIntersect(Span[int]{4, 3}, Span[int]{2, 1}))
	require.NoError(t, isTwoSpansNotIntersect(Span[int]{2, 2}, Span[int]{4, 3}))
	require.NoError(t, isTwoSpansNotIntersect(Span[int]{4, 3}, Span[int]{2, 2}))

	require.Error(t, isTwoSpansNotIntersect(Span[int]{2, 1}, Span[int]{4, 2}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{4, 2}, Span[int]{2, 1}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{3, 1}, Span[int]{4, 2}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{4, 2}, Span[int]{3, 1}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{4, 1}, Span[int]{4, 2}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{4, 2}, Span[int]{4, 1}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{5, 1}, Span[int]{4, 2}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{4, 2}, Span[int]{5, 1}))

	require.Error(t, isTwoSpansNotIntersect(Span[int]{1, 2}, Span[int]{4, 3}))
}

func TestIsContinuous(t *testing.T) {
	continuous := []Span[uint]{{2, 6}, {7, 11}, {12, 16}, {17, 18}}
	discontinuous := []Span[uint]{{2, 6}, {7, 11}, {13, 16}, {17, 18}}
	intersected := []Span[uint]{{2, 6}, {7, 11}, {10, 16}, {17, 18}}

	require.NoError(t, IsContinuous[int](nil))
	require.NoError(t, IsContinuous([]Span[int]{}))
	require.NoError(t, IsContinuous([]Span[int]{{2, 6}}))
	require.NoError(t, IsContinuous([]Span[int]{{6, 2}}))
	require.NoError(t, IsContinuous(continuous))

	require.Error(t, IsContinuous(discontinuous))
	require.Error(t, IsContinuous(intersected))
}

func TestIsTwoSpansContinuous(t *testing.T) {
	require.NoError(t, isTwoSpansContinuous(Span[int]{-2, -1}, Span[int]{0, 1}))
	require.NoError(t, isTwoSpansContinuous(Span[int]{0, 1}, Span[int]{-2, -1}))
	require.NoError(t, isTwoSpansContinuous(Span[int]{1, 0}, Span[int]{-1, -2}))
	require.NoError(t, isTwoSpansContinuous(Span[int]{-1, -2}, Span[int]{1, 0}))
	require.NoError(t, isTwoSpansContinuous(Span[int]{math.MinInt, math.MaxInt - 1}, Span[int]{math.MaxInt, math.MaxInt}))
	require.NoError(t, isTwoSpansContinuous(Span[int]{math.MaxInt, math.MaxInt}, Span[int]{math.MinInt, math.MaxInt - 1}))

	require.NoError(t, isTwoSpansContinuous(Span[uint]{1, 2}, Span[uint]{3, 4}))
	require.NoError(t, isTwoSpansContinuous(Span[uint]{3, 4}, Span[uint]{1, 2}))
	require.NoError(t, isTwoSpansContinuous(Span[uint]{4, 3}, Span[uint]{2, 1}))
	require.NoError(t, isTwoSpansContinuous(Span[uint]{2, 1}, Span[uint]{4, 3}))
	require.NoError(t, isTwoSpansContinuous(Span[uint]{0, math.MaxUint - 1}, Span[uint]{math.MaxUint, math.MaxUint}))
	require.NoError(t, isTwoSpansContinuous(Span[uint]{math.MaxUint, math.MaxUint}, Span[uint]{0, math.MaxUint - 1}))

	require.Error(t, isTwoSpansContinuous(Span[int]{-2, -1}, Span[int]{1, 2}))
	require.Error(t, isTwoSpansContinuous(Span[int]{1, 2}, Span[int]{-2, -1}))
	require.Error(t, isTwoSpansContinuous(Span[int]{2, 1}, Span[int]{-1, -2}))
	require.Error(t, isTwoSpansContinuous(Span[int]{-1, -2}, Span[int]{2, 1}))
	require.Error(t, isTwoSpansContinuous(Span[int]{math.MinInt, math.MaxInt - 2}, Span[int]{math.MaxInt, math.MaxInt}))
	require.Error(t, isTwoSpansContinuous(Span[int]{math.MaxInt, math.MaxInt}, Span[int]{math.MinInt, math.MaxInt - 2}))

	require.Error(t, isTwoSpansContinuous(Span[uint]{0, 1}, Span[uint]{3, 4}))
	require.Error(t, isTwoSpansContinuous(Span[uint]{3, 4}, Span[uint]{0, 1}))
	require.Error(t, isTwoSpansContinuous(Span[uint]{4, 3}, Span[uint]{1, 0}))
	require.Error(t, isTwoSpansContinuous(Span[uint]{1, 0}, Span[uint]{4, 3}))
	require.Error(t, isTwoSpansContinuous(Span[uint]{0, math.MaxUint - 2}, Span[uint]{math.MaxUint, math.MaxUint}))
	require.Error(t, isTwoSpansContinuous(Span[uint]{math.MaxUint, math.MaxUint}, Span[uint]{0, math.MaxUint - 2}))

	require.Error(t, isTwoSpansContinuous(Span[int]{1, 2}, Span[int]{4, 3}))
	require.Error(t, isTwoSpansContinuous(Span[int]{1, 2}, Span[int]{1, 3}))
}
