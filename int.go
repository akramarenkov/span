package span

import (
	"github.com/akramarenkov/safe"
	"golang.org/x/exp/constraints"
)

// Divides a sequence of integers evenly from begin to end inclusive into a specified
// quantity of spans.
func Int[Type constraints.Integer](begin, end, quantity Type) ([]Span[Type], error) {
	if begin > end {
		return nil, ErrBeginGreaterEnd
	}

	if quantity < 0 {
		return nil, ErrSpansQuantityNegative
	}

	if quantity == 0 {
		return nil, ErrSpansQuantityZero
	}

	if quantity == 1 {
		spans := []Span[Type]{{Begin: begin, End: end}}
		return spans, nil
	}

	// Overflow is not possible with these operations given the checks on the values
	// ​​of the input arguments located above
	distance, _ := safe.SubDiv(end, begin, quantity)
	remainder, _ := safe.SubDivRem(end, begin, quantity)

	// +1 due to the constant presence of begin in the sequence
	//
	// Overflow on this operation is impossible because maximum value of remainder is
	// maximum value of the divisor minus one and at positive divisor is maximum value
	// for given type minus one
	remainder++

	if distance == 0 {
		distance = 1
		remainder = 0
	}

	spans := make([]Span[Type], 0, quantity)

	for actualBegin, actualEnd := begin, begin+distance-1; ; {
		if remainder != 0 {
			actualEnd++
			remainder--
		}

		item := Span[Type]{
			Begin: actualBegin,
			End:   actualEnd,
		}

		spans = append(spans, item)

		if actualEnd == end {
			return spans, nil
		}

		actualBegin = actualEnd + 1
		actualEnd += distance
	}
}

// Divides a sequence of integers from begin to end inclusive into spans of the
// specified width.
func IntWidth[Type constraints.Integer](begin, end, width Type) ([]Span[Type], error) {
	if begin > end {
		return nil, ErrBeginGreaterEnd
	}

	if width < 0 {
		return nil, ErrSpanWidthNegative
	}

	if width == 0 {
		return nil, ErrSpanWidthZero
	}

	spans := make([]Span[Type], safe.IncStepSize(begin, end, width))

	for id, actualBegin := range safe.IncStep(begin, end, width) {
		actualEnd, err := safe.Add(actualBegin, width-1)
		if err != nil {
			actualEnd = end
		}

		if actualEnd > end {
			actualEnd = end
		}

		span := Span[Type]{
			Begin: actualBegin,
			End:   actualEnd,
		}

		spans[id] = span
	}

	return spans, nil
}
