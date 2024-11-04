package span

import (
	"golang.org/x/exp/constraints"
)

// Compare function for increasing sequence of spans.
//
// Partially detects spans intersections, but does not guarantee complete verification.
//
// Partially detects the presence of different types of span sequences, but does not
// guarantee complete verification.
func CompareInc[Type constraints.Ordered](first, second Span[Type]) int {
	if first.Begin > first.End || second.Begin > second.End {
		panic(ErrSpansUnexpectedSequencing)
	}

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
// Partially detects spans intersections, but does not guarantee complete verification.
//
// Partially detects the presence of different types of span sequences, but does not
// guarantee complete verification.
func CompareDec[Type constraints.Ordered](first, second Span[Type]) int {
	if first.Begin < first.End || second.Begin < second.End {
		panic(ErrSpansUnexpectedSequencing)
	}

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
// Partially detects spans intersections, but does not guarantee complete verification.
//
// Partially detects the presence of different types of span sequences, but does not
// guarantee complete verification.
func Compare[Type constraints.Ordered](first, second Span[Type]) int {
	switch {
	case first.Begin == first.End && second.Begin == second.End:
		// There is no way to determine whether the sequence is increasing or decreasing
		return 0
	case first.Begin <= first.End && second.Begin <= second.End:
		switch {
		case first.End < second.Begin:
			return -1
		case first.End > second.Begin:
			return 1
		}

		panic(ErrSpansIntersect)
	case first.Begin >= first.End && second.Begin >= second.End:
		switch {
		case first.End > second.Begin:
			return -1
		case first.End < second.Begin:
			return 1
		}

		panic(ErrSpansIntersect)
	}

	panic(ErrSpansDiffSequencing)
}
