package span

import "golang.org/x/exp/constraints"

func isTwoSpansIntersectInc[Type constraints.Ordered](first, second Span[Type]) error {
	if first.End >= second.Begin && first.End <= second.End {
		return ErrSpansIntersect
	}

	if first.Begin >= second.Begin && first.Begin <= second.End {
		return ErrSpansIntersect
	}

	if second.End >= first.Begin && second.End <= first.End {
		return ErrSpansIntersect
	}

	return nil
}

func isTwoSpansIntersectDec[Type constraints.Ordered](first, second Span[Type]) error {
	if first.End <= second.Begin && first.End >= second.End {
		return ErrSpansIntersect
	}

	if first.Begin <= second.Begin && first.Begin >= second.End {
		return ErrSpansIntersect
	}

	if second.End <= first.Begin && second.End >= first.End {
		return ErrSpansIntersect
	}

	return nil
}
