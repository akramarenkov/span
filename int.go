package span

import (
	"github.com/akramarenkov/safe"
	"golang.org/x/exp/constraints"
)

// Divides a sequence of integers from begin to end inclusive into spans.
func Int[Type constraints.Integer](begin Type, end Type, quantity Type) ([]Span[Type], error) {
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
		spans := []Span[Type]{
			{
				Begin: begin,
				End:   end,
			},
		}

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

	for scopeBegin, scopeEnd := begin, begin+distance-1; ; {
		if remainder != 0 {
			scopeEnd++
			remainder--
		}

		item := Span[Type]{
			Begin: scopeBegin,
			End:   scopeEnd,
		}

		spans = append(spans, item)

		if scopeEnd == end {
			return spans, nil
		}

		scopeBegin = scopeEnd + 1
		scopeEnd += distance
	}
}
