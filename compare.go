package span

import "cmp"

// Compare function for sorting of increasing sequence of spans.
//
// Partially detects spans intersections, but does not guarantee complete verification.
//
// Partially detects the presence of spans that have a different sequence type, but
// does not guarantee complete verification.
func CompareInc[Type cmp.Ordered](first, second Span[Type]) int {
	if first.Begin > first.End || second.Begin > second.End {
		panic(ErrSpanSequenceNotNonDecreasing)
	}

	switch {
	case first.End < second.Begin:
		return -1
	case first.Begin > second.End:
		return 1
	}

	panic(ErrSpansIntersect)
}

// Compare function for sorting of decreasing sequence of spans.
//
// Partially detects spans intersections, but does not guarantee complete verification.
//
// Partially detects the presence of spans that have a different sequence type, but
// does not guarantee complete verification.
func CompareDec[Type cmp.Ordered](first, second Span[Type]) int {
	if first.Begin < first.End || second.Begin < second.End {
		panic(ErrSpanSequenceNotNonIncreasing)
	}

	switch {
	case first.End > second.Begin:
		return -1
	case first.Begin < second.End:
		return 1
	}

	panic(ErrSpansIntersect)
}

// Compare function for searching in increasing sequence of spans.
//
// Partially detects the presence of spans that have a different sequence type, but
// does not guarantee complete verification.
func SearchInc[Type cmp.Ordered](item, target Span[Type]) int {
	if item.Begin > item.End || target.Begin > target.End {
		panic(ErrSpanSequenceNotNonDecreasing)
	}

	switch {
	case target.End > item.End:
		return -1
	case target.Begin < item.Begin:
		return 1
	}

	return 0
}

// Compare function for searching in decreasing sequence of spans.
//
// Partially detects the presence of spans that have a different sequence type, but
// does not guarantee complete verification.
func SearchDec[Type cmp.Ordered](item, target Span[Type]) int {
	if item.Begin < item.End || target.Begin < target.End {
		panic(ErrSpanSequenceNotNonIncreasing)
	}

	switch {
	case target.End < item.End:
		return -1
	case target.Begin > item.Begin:
		return 1
	}

	return 0
}
