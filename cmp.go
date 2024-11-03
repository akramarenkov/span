package span

import (
	"golang.org/x/exp/constraints"
)

// Compare function for increasing sequence of spans.
//
// Partially detects spans intersections, but does not guarantee full verification.
func CompareInc[Type constraints.Ordered](first, second Span[Type]) int {
	switch {
	case first.End < second.Begin:
		return -1
	case first.End > second.Begin:
		return 1
	}

	panic(ErrSpansIntersect)
}

// Compare function for decreasing sequence of spans.
//
// Partially detects spans intersections, but does not guarantee full verification.
func CompareDec[Type constraints.Ordered](first, second Span[Type]) int {
	switch {
	case first.End > second.Begin:
		return -1
	case first.End < second.Begin:
		return 1
	}

	panic(ErrSpansIntersect)
}

// Compare function for sequence of spans.
//
// Partially detects spans intersections, but does not guarantee full verification.
//
// Partially detects mixing of span sequence types, but does not guarantee full
// verification.
func Compare[Type constraints.Ordered](first, second Span[Type]) int {
	switch {
	case first.Begin <= first.End && second.Begin <= second.End:
		switch {
		case first.End < second.Begin:
			return -1
		case first.End > second.Begin:
			return 1
		}

		panic(ErrSpansIntersect)
	case first.Begin > first.End && second.Begin > second.End:
		switch {
		case first.End > second.Begin:
			return -1
		case first.End < second.Begin:
			return 1
		}
	}

	panic(ErrSpansDiffSequenceType)
}
