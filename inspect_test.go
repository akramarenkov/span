package span

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsNotDiffSequencing(t *testing.T) {
	require.NoError(t, IsNotDiffSequencing[int](nil))
	require.NoError(t, IsNotDiffSequencing([]Span[int]{}))
	require.NoError(t, IsNotDiffSequencing([]Span[int]{{2, 6}}))
	require.NoError(t, IsNotDiffSequencing([]Span[int]{{2, 6}, {7, 11}}))
	require.NoError(t, IsNotDiffSequencing([]Span[int]{{2, 6}, {7, 11}, {12, 16}, {17, 18}}))
	require.NoError(t, IsNotDiffSequencing([]Span[int]{{2, 2}, {7, 11}, {12, 16}, {17, 18}}))
	require.NoError(t, IsNotDiffSequencing([]Span[int]{{2, 6}, {7, 7}, {12, 16}, {17, 18}}))

	require.Error(t, IsNotDiffSequencing([]Span[int]{{2, 6}, {11, 7}}))
	require.Error(t, IsNotDiffSequencing([]Span[int]{{2, 6}, {7, 11}, {12, 16}, {18, 17}}))
	require.Error(t, IsNotDiffSequencing([]Span[int]{{2, 2}, {7, 11}, {12, 16}, {18, 17}}))
	require.Error(t, IsNotDiffSequencing([]Span[int]{{2, 6}, {7, 7}, {12, 16}, {18, 17}}))
}

func TestIsTwoSpansNotDiffSequencing(t *testing.T) {
	require.NoError(t, isTwoSpansNotDiffSequencing(Span[int]{1, 2}, Span[int]{3, 4}))
	require.NoError(t, isTwoSpansNotDiffSequencing(Span[int]{2, 1}, Span[int]{4, 3}))
	require.NoError(t, isTwoSpansNotDiffSequencing(Span[int]{1, 1}, Span[int]{2, 2}))
	require.NoError(t, isTwoSpansNotDiffSequencing(Span[int]{2, 2}, Span[int]{1, 1}))
	require.NoError(t, isTwoSpansNotDiffSequencing(Span[int]{2, 2}, Span[int]{3, 4}))
	require.NoError(t, isTwoSpansNotDiffSequencing(Span[int]{3, 4}, Span[int]{2, 2}))
	require.NoError(t, isTwoSpansNotDiffSequencing(Span[int]{2, 2}, Span[int]{4, 3}))
	require.NoError(t, isTwoSpansNotDiffSequencing(Span[int]{4, 3}, Span[int]{2, 2}))

	require.Error(t, isTwoSpansNotDiffSequencing(Span[int]{1, 2}, Span[int]{4, 3}))
	require.Error(t, isTwoSpansNotDiffSequencing(Span[int]{2, 1}, Span[int]{3, 4}))
}

func TestIsNonDecreasing(t *testing.T) {
	require.NoError(t, IsNonDecreasing[int](nil))
	require.NoError(t, IsNonDecreasing([]Span[int]{}))
	require.NoError(t, IsNonDecreasing([]Span[int]{{2, 6}}))
	require.NoError(t, IsNonDecreasing([]Span[int]{{2, 6}, {7, 11}, {12, 16}, {17, 18}}))
	require.NoError(t, IsNonDecreasing([]Span[int]{{2, 2}, {7, 11}, {12, 16}, {17, 18}}))

	require.Error(t, IsNonDecreasing([]Span[int]{{6, 2}}))
	require.Error(t, IsNonDecreasing([]Span[int]{{2, 6}, {11, 7}, {12, 16}, {17, 18}}))
}

func TestIsNonIncreasing(t *testing.T) {
	require.NoError(t, IsNonIncreasing[int](nil))
	require.NoError(t, IsNonIncreasing([]Span[int]{}))
	require.NoError(t, IsNonIncreasing([]Span[int]{{6, 2}}))
	require.NoError(t, IsNonIncreasing([]Span[int]{{18, 17}, {16, 12}, {11, 7}, {6, 2}}))
	require.NoError(t, IsNonIncreasing([]Span[int]{{18, 17}, {16, 16}, {11, 7}, {6, 2}}))

	require.Error(t, IsNonIncreasing([]Span[int]{{2, 6}}))
	require.Error(t, IsNonIncreasing([]Span[int]{{18, 17}, {16, 12}, {7, 11}, {6, 2}}))
}

func TestIsNotIntersect(t *testing.T) {
	require.NoError(t, IsNotIntersect[int](nil))
	require.NoError(t, IsNotIntersect([]Span[int]{}))
	require.NoError(t, IsNotIntersect([]Span[int]{{2, 6}}))
	require.NoError(t, IsNotIntersect([]Span[int]{{2, 6}, {7, 11}}))
	require.NoError(t, IsNotIntersect([]Span[int]{{2, 6}, {11, 7}}))
	require.NoError(t, IsNotIntersect([]Span[int]{{2, 6}, {7, 11}, {12, 16}, {17, 18}}))
	require.NoError(t, IsNotIntersect([]Span[int]{{2, 2}, {7, 11}, {12, 16}, {17, 18}}))
	require.NoError(t, IsNotIntersect([]Span[int]{{2, 6}, {7, 7}, {12, 16}, {17, 18}}))
	require.NoError(t, IsNotIntersect([]Span[int]{{2, 6}, {7, 7}, {16, 12}, {17, 18}}))

	require.Error(t, IsNotIntersect([]Span[int]{{2, 6}, {6, 11}}))
	require.Error(t, IsNotIntersect([]Span[int]{{2, 8}, {11, 7}}))
	require.Error(t, IsNotIntersect([]Span[int]{{2, 6}, {7, 11}, {12, 16}, {16, 18}}))
	require.Error(t, IsNotIntersect([]Span[int]{{2, 6}, {6, 6}, {12, 16}, {16, 18}}))
	require.Error(t, IsNotIntersect([]Span[int]{{2, 6}, {7, 7}, {12, 16}, {7, 8}}))
	require.Error(t, IsNotIntersect([]Span[int]{{2, 6}, {7, 7}, {17, 12}, {17, 18}}))
	require.Error(t, IsNotIntersect([]Span[int]{{2, 6}, {7, 7}, {16, 3}, {17, 18}}))
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

	require.NoError(t, isTwoSpansNotIntersect(Span[int]{1, 2}, Span[int]{4, 3}))
	require.NoError(t, isTwoSpansNotIntersect(Span[int]{4, 3}, Span[int]{1, 2}))
	require.NoError(t, isTwoSpansNotIntersect(Span[int]{1, 2}, Span[int]{5, 3}))
	require.NoError(t, isTwoSpansNotIntersect(Span[int]{5, 3}, Span[int]{1, 2}))
	require.NoError(t, isTwoSpansNotIntersect(Span[int]{1, 2}, Span[int]{0, -2}))
	require.NoError(t, isTwoSpansNotIntersect(Span[int]{0, -2}, Span[int]{1, 2}))

	require.Error(t, isTwoSpansNotIntersect(Span[int]{1, 3}, Span[int]{4, 3}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{4, 3}, Span[int]{1, 3}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{1, 4}, Span[int]{5, 3}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{5, 3}, Span[int]{1, 4}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{0, 2}, Span[int]{0, -2}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{0, -2}, Span[int]{0, 2}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{-1, 2}, Span[int]{0, -2}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{0, -2}, Span[int]{-1, 2}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{1, 3}, Span[int]{2, 2}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{2, 2}, Span[int]{1, 3}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{1, 5}, Span[int]{3, 2}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{3, 2}, Span[int]{1, 5}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{5, 1}, Span[int]{2, 3}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{2, 3}, Span[int]{5, 1}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{1, 5}, Span[int]{4, 2}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{4, 2}, Span[int]{1, 5}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{5, 1}, Span[int]{2, 4}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{2, 4}, Span[int]{5, 1}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{2, 3}, Span[int]{3, 2}))
	require.Error(t, isTwoSpansNotIntersect(Span[int]{3, 2}, Span[int]{2, 3}))
}

func TestIsContinuous(t *testing.T) {
	require.NoError(t, IsContinuous[int](nil))
	require.NoError(t, IsContinuous([]Span[int]{}))
	require.NoError(t, IsContinuous([]Span[int]{{2, 6}}))
	require.NoError(t, IsContinuous([]Span[int]{{6, 2}}))
	require.NoError(t, IsContinuous([]Span[uint]{{2, 6}, {7, 11}, {12, 16}, {17, 18}}))

	require.Error(t, IsContinuous([]Span[uint]{{2, 6}, {7, 11}, {13, 16}, {17, 18}}))
	require.Error(t, IsContinuous([]Span[uint]{{2, 6}, {7, 11}, {10, 16}, {17, 18}}))
	require.Error(t, IsContinuous([]Span[uint]{{2, 6}, {7, 11}, {5, 16}, {17, 18}}))
	require.Error(t, IsContinuous([]Span[uint]{{2, 6}, {7, 11}, {2, 2}, {17, 18}}))
	require.Error(t, IsContinuous([]Span[uint]{{2, 6}, {7, 11}, {16, 12}, {17, 18}}))
	require.Error(t, IsContinuous([]Span[uint]{{2, 6}, {7, 11}, {12, 12}, {11, 11}}))
	require.Error(t, IsContinuous([]Span[uint]{{2, 6}, {7, 11}, {12, 12}, {11, 10}}))
	require.Error(t, IsContinuous([]Span[uint]{{2, 6}, {7, 11}, {12, 11}, {11, 12}}))
	require.Error(t, IsContinuous([]Span[uint]{{2, 2}, {3, 3}, {4, 4}, {3, 3}}))
}

func TestIsTwoSpansContinuous(t *testing.T) {
	require.NoError(t, isTwoSpansContinuous(Span[int]{-2, -1}, Span[int]{0, 1}))
	require.NoError(t, isTwoSpansContinuous(Span[int]{0, 1}, Span[int]{-2, -1}))
	require.NoError(t, isTwoSpansContinuous(Span[int]{1, 0}, Span[int]{-1, -2}))
	require.NoError(t, isTwoSpansContinuous(Span[int]{-1, -2}, Span[int]{1, 0}))
	require.NoError(
		t,
		isTwoSpansContinuous(
			Span[int]{math.MinInt, math.MaxInt - 1},
			Span[int]{math.MaxInt, math.MaxInt},
		),
	)
	require.NoError(
		t,
		isTwoSpansContinuous(
			Span[int]{math.MaxInt, math.MaxInt},
			Span[int]{math.MinInt, math.MaxInt - 1},
		),
	)

	require.NoError(t, isTwoSpansContinuous(Span[uint]{1, 2}, Span[uint]{3, 4}))
	require.NoError(t, isTwoSpansContinuous(Span[uint]{3, 4}, Span[uint]{1, 2}))
	require.NoError(t, isTwoSpansContinuous(Span[uint]{4, 3}, Span[uint]{2, 1}))
	require.NoError(t, isTwoSpansContinuous(Span[uint]{2, 1}, Span[uint]{4, 3}))
	require.NoError(
		t,
		isTwoSpansContinuous(
			Span[uint]{0, math.MaxUint - 1},
			Span[uint]{math.MaxUint, math.MaxUint},
		),
	)
	require.NoError(
		t,
		isTwoSpansContinuous(
			Span[uint]{math.MaxUint, math.MaxUint},
			Span[uint]{0, math.MaxUint - 1},
		),
	)

	require.Error(t, isTwoSpansContinuous(Span[int]{-2, -1}, Span[int]{1, 2}))
	require.Error(t, isTwoSpansContinuous(Span[int]{1, 2}, Span[int]{-2, -1}))
	require.Error(t, isTwoSpansContinuous(Span[int]{2, 1}, Span[int]{-1, -2}))
	require.Error(t, isTwoSpansContinuous(Span[int]{-1, -2}, Span[int]{2, 1}))
	require.Error(
		t,
		isTwoSpansContinuous(
			Span[int]{math.MinInt, math.MaxInt - 2},
			Span[int]{math.MaxInt, math.MaxInt},
		),
	)
	require.Error(
		t,
		isTwoSpansContinuous(
			Span[int]{math.MaxInt, math.MaxInt},
			Span[int]{math.MinInt, math.MaxInt - 2},
		),
	)

	require.Error(t, isTwoSpansContinuous(Span[uint]{0, 1}, Span[uint]{3, 4}))
	require.Error(t, isTwoSpansContinuous(Span[uint]{3, 4}, Span[uint]{0, 1}))
	require.Error(t, isTwoSpansContinuous(Span[uint]{4, 3}, Span[uint]{1, 0}))
	require.Error(t, isTwoSpansContinuous(Span[uint]{1, 0}, Span[uint]{4, 3}))
	require.Error(
		t,
		isTwoSpansContinuous(
			Span[uint]{0, math.MaxUint - 2},
			Span[uint]{math.MaxUint, math.MaxUint},
		),
	)
	require.Error(
		t,
		isTwoSpansContinuous(
			Span[uint]{math.MaxUint, math.MaxUint},
			Span[uint]{0, math.MaxUint - 2},
		),
	)

	require.Error(t, isTwoSpansContinuous(Span[int]{1, 2}, Span[int]{4, 3}))
	require.Error(t, isTwoSpansContinuous(Span[int]{4, 3}, Span[int]{1, 2}))
	require.Error(t, isTwoSpansContinuous(Span[int]{1, 2}, Span[int]{3, 2}))
	require.Error(t, isTwoSpansContinuous(Span[int]{3, 2}, Span[int]{1, 2}))
	require.Error(t, isTwoSpansContinuous(Span[int]{1, 2}, Span[int]{1, 3}))
	require.Error(t, isTwoSpansContinuous(Span[int]{1, 3}, Span[int]{1, 2}))
}
