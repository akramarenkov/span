package span

import (
	"github.com/akramarenkov/safe"
	"golang.org/x/exp/constraints"
)

const (
	significantSpansQuantity = 2
)

// Checks that a sequence of spans does not contain spans with different sequence types
// (increasing or decreasing).
func IsNotDiffSequencing[Type constraints.Ordered](spans []Span[Type]) error {
	if len(spans) < significantSpansQuantity {
		return nil
	}

	// Since one of the spans being compared may have the same Begin and End,
	// while the other does not, it is necessary to compare all spans with all
	for id, first := range spans[:len(spans)-1] {
		for _, second := range spans[id+1:] {
			if err := isTwoSpansNotDiffSequencing(first, second); err != nil {
				return err
			}
		}
	}

	return nil
}

func isTwoSpansNotDiffSequencing[Type constraints.Ordered](first, second Span[Type]) error {
	if first.Begin <= first.End && second.Begin <= second.End {
		return nil
	}

	if first.Begin >= first.End && second.Begin >= second.End {
		return nil
	}

	return ErrSpansDiffSequencing
}

// Checks that a spans sequence consists of only increasing spans.
func IsIncreasing[Type constraints.Ordered](spans []Span[Type]) error {
	if len(spans) == 0 {
		return ErrSpansUnexpectedSequencing
	}

	for _, span := range spans {
		if span.Begin > span.End {
			return ErrSpansUnexpectedSequencing
		}
	}

	return nil
}

// Checks that a spans sequence consists of only decreasing spans.
func IsDecreasing[Type constraints.Ordered](spans []Span[Type]) error {
	if len(spans) == 0 {
		return ErrSpansUnexpectedSequencing
	}

	for _, span := range spans {
		if span.Begin < span.End {
			return ErrSpansUnexpectedSequencing
		}
	}

	return nil
}

// Checks that a sequence of spans does not contain intersect spans.
func IsNotIntersect[Type constraints.Ordered](spans []Span[Type]) error {
	if len(spans) < significantSpansQuantity {
		return nil
	}

	for id, first := range spans[:len(spans)-1] {
		for _, second := range spans[id+1:] {
			if err := isTwoSpansNotIntersect(first, second); err != nil {
				return err
			}
		}
	}

	return nil
}

func isTwoSpansNotIntersect[Type constraints.Ordered](first, second Span[Type]) error {
	if err := isTwoSpansNotDiffSequencing(first, second); err != nil {
		return err
	}

	if first.End >= second.Begin && first.End <= second.End {
		return ErrSpansIntersect
	}

	if first.Begin >= second.Begin && first.Begin <= second.End {
		return ErrSpansIntersect
	}

	if second.End >= first.Begin && second.End <= first.End {
		return ErrSpansIntersect
	}

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

// Checks that a sequence of spans is continuous, monotone and does not contain
// intersecting spans.
func IsContinuous[Type constraints.Integer](spans []Span[Type]) error {
	if len(spans) < significantSpansQuantity {
		return nil
	}

	for id, first := range spans[:len(spans)-1] {
		second := spans[id+1]

		if err := isTwoSpansContinuous(first, second); err != nil {
			return err
		}
	}

	return nil
}

func isTwoSpansContinuous[Type constraints.Integer](first, second Span[Type]) error {
	if err := isTwoSpansNotIntersect(first, second); err != nil {
		return err
	}

	if diff, err := safe.Sub(second.Begin, first.End); err == nil {
		if diff == 1 {
			return nil
		}
	}

	if diff, err := safe.Sub(first.Begin, second.End); err == nil {
		if diff == 1 {
			return nil
		}
	}

	if diff, err := safe.Sub(first.End, second.Begin); err == nil {
		if diff == 1 {
			return nil
		}
	}

	if diff, err := safe.Sub(second.End, first.Begin); err == nil {
		if diff == 1 {
			return nil
		}
	}

	return ErrSpansDiscontinuous
}
