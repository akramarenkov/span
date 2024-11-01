package span

import (
	"github.com/akramarenkov/safe"
	"golang.org/x/exp/constraints"
)

// Divides a linear sequence of integers evenly from begin to end inclusive into a
// specified quantity of spans.
//
// If begin is greater than end, the sequence will be considered decreasing,
// otherwise - increasing.
//
// If a zero or negative quantity of spans is specified, an error is returned.
func Evenly[Type constraints.Integer](begin, end, quantity Type) ([]Span[Type], error) {
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

	distance, remainder := evenlyDistance(begin, end, quantity)

	spans := make([]Span[Type], 0, quantity)

	if begin < end {
		for spanBegin, spanEnd := begin, begin+distance-1; ; {
			if remainder != 0 {
				spanEnd++
				remainder--
			}

			span := Span[Type]{
				Begin: spanBegin,
				End:   spanEnd,
			}

			spans = append(spans, span)

			if spanEnd == end {
				return spans, nil
			}

			spanBegin = spanEnd + 1
			spanEnd += distance
		}
	}

	for spanBegin, spanEnd := begin, begin-distance+1; ; {
		if remainder != 0 {
			spanEnd--
			remainder--
		}

		span := Span[Type]{
			Begin: spanBegin,
			End:   spanEnd,
		}

		spans = append(spans, span)

		if spanEnd == end {
			return spans, nil
		}

		spanBegin = spanEnd - 1
		spanEnd -= distance
	}
}

func evenlyDistance[Type constraints.Integer](begin, end, quantity Type) (Type, Type) {
	if begin < end {
		// Overflow is not possible with these operations given the checks on the values
		// ​​of the input arguments located above in the calling function
		distance, _ := safe.SubDiv(end, begin, quantity)
		remainder, _ := safe.SubDivRem(end, begin, quantity)

		// +1 due to the constant presence of begin in the sequence
		//
		// Overflow on this operation is impossible because maximum value of remainder
		// is maximum value of the divisor minus one and at positive divisor is maximum
		// value for given type minus one
		remainder++

		if distance == 0 {
			distance = 1
			remainder = 0
		}

		return distance, remainder
	}

	distance, _ := safe.SubDiv(begin, end, quantity)
	remainder, _ := safe.SubDivRem(begin, end, quantity)

	remainder++

	if distance == 0 {
		distance = 1
		remainder = 0
	}

	return distance, remainder
}
