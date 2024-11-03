package span

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsTwoSpansIntersectInc(t *testing.T) {
	require.NoError(t, isTwoSpansIntersectInc(Span[int]{1, 2}, Span[int]{3, 4}))
	require.NoError(t, isTwoSpansIntersectInc(Span[int]{3, 4}, Span[int]{1, 2}))
	require.NoError(t, isTwoSpansIntersectInc(Span[int]{2, 2}, Span[int]{3, 4}))
	require.NoError(t, isTwoSpansIntersectInc(Span[int]{3, 4}, Span[int]{2, 2}))
	require.NoError(t, isTwoSpansIntersectInc(Span[int]{2, 2}, Span[int]{3, 3}))
	require.NoError(t, isTwoSpansIntersectInc(Span[int]{3, 3}, Span[int]{2, 2}))

	require.Error(t, isTwoSpansIntersectInc(Span[int]{1, 2}, Span[int]{2, 4}))
	require.Error(t, isTwoSpansIntersectInc(Span[int]{2, 4}, Span[int]{1, 2}))
	require.Error(t, isTwoSpansIntersectInc(Span[int]{1, 3}, Span[int]{2, 4}))
	require.Error(t, isTwoSpansIntersectInc(Span[int]{2, 4}, Span[int]{1, 3}))
	require.Error(t, isTwoSpansIntersectInc(Span[int]{1, 4}, Span[int]{2, 4}))
	require.Error(t, isTwoSpansIntersectInc(Span[int]{2, 4}, Span[int]{1, 4}))
	require.Error(t, isTwoSpansIntersectInc(Span[int]{1, 5}, Span[int]{2, 4}))
	require.Error(t, isTwoSpansIntersectInc(Span[int]{2, 4}, Span[int]{1, 5}))
}

func TestIsTwoSpansIntersectDec(t *testing.T) {
	require.NoError(t, isTwoSpansIntersectDec(Span[int]{2, 1}, Span[int]{4, 3}))
	require.NoError(t, isTwoSpansIntersectDec(Span[int]{4, 3}, Span[int]{2, 1}))
	require.NoError(t, isTwoSpansIntersectDec(Span[int]{2, 2}, Span[int]{4, 3}))
	require.NoError(t, isTwoSpansIntersectDec(Span[int]{4, 3}, Span[int]{2, 2}))
	require.NoError(t, isTwoSpansIntersectInc(Span[int]{2, 2}, Span[int]{3, 3}))
	require.NoError(t, isTwoSpansIntersectInc(Span[int]{3, 3}, Span[int]{2, 2}))

	require.Error(t, isTwoSpansIntersectDec(Span[int]{2, 1}, Span[int]{4, 2}))
	require.Error(t, isTwoSpansIntersectDec(Span[int]{4, 2}, Span[int]{2, 1}))
	require.Error(t, isTwoSpansIntersectDec(Span[int]{3, 1}, Span[int]{4, 2}))
	require.Error(t, isTwoSpansIntersectDec(Span[int]{4, 2}, Span[int]{3, 1}))
	require.Error(t, isTwoSpansIntersectDec(Span[int]{4, 1}, Span[int]{4, 2}))
	require.Error(t, isTwoSpansIntersectDec(Span[int]{4, 2}, Span[int]{4, 1}))
	require.Error(t, isTwoSpansIntersectDec(Span[int]{5, 1}, Span[int]{4, 2}))
	require.Error(t, isTwoSpansIntersectDec(Span[int]{4, 2}, Span[int]{5, 1}))
}
