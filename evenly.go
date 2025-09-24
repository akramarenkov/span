package span

import (
	"iter"

	"github.com/akramarenkov/safe"
	"golang.org/x/exp/constraints"
)

// Divides a linear sequence of integers evenly from begin to end inclusive into a
// specified quantity of spans.
//
// If begin is greater than end, the sequence of integers will be considered decreasing,
// otherwise - increasing.
//
// Length of the returned slice can be less than the specified quantity of spans, but
// cannot be greater.
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

	capacity := uint64(quantity)

	if distance == 1 && remainder == 0 {
		capacity = safe.IterSize(begin, end)
	}

	spans := make([]Span[Type], 0, capacity)

	if begin > end {
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

// A range iterator used to iterating over a sequence of spans obtained by dividing
// a linear sequence of integers evenly from begin to end inclusive into a specified
// quantity of spans.
//
// If begin is greater than end, the sequence of spans will be decreasing,
// otherwise - increasing.
//
// Quantity of iterations can be less than the specified quantity of spans, but
// cannot be greater.
//
// If a zero or negative quantity of spans is specified, the iterator will panic.
//
// Works like [Evenly] but does not perform memory allocation.
func Even[Type constraints.Integer](begin, end, quantity Type) iter.Seq2[uint64, Span[Type]] {
	if quantity < 0 {
		panic(ErrSpansQuantityNegative)
	}

	if quantity == 0 {
		panic(ErrSpansQuantityZero)
	}

	iterator := func(yield func(uint64, Span[Type]) bool) {
		if quantity == 1 {
			yield(0, Span[Type]{Begin: begin, End: end})
			return
		}

		distance, remainder := evenlyDistance(begin, end, quantity)

		if begin > end {
			evenDec(begin, end, distance, remainder, yield)
			return
		}

		evenInc(begin, end, distance, remainder, yield)
	}

	return iterator
}

// A range iterator used to iterating over a subslices obtained by evenly dividing
// a main slice into a specified quantity of subslices.
//
// Quantity of iterations can be less than the specified quantity of subslices, but
// cannot be greater.
//
// If a zero or negative quantity of subslices is specified, the iterator will panic.
func EvenSlices[Type any, TypeQ constraints.Integer](divisible []Type, quantity TypeQ) iter.Seq[[]Type] {
	if quantity < 0 {
		panic(ErrSpansQuantityNegative)
	}

	if quantity == 0 {
		panic(ErrSpansQuantityZero)
	}

	iterator := func(yield func([]Type) bool) {
		if len(divisible) == 0 {
			return
		}

		if quantity == 1 {
			yield(divisible)
			return
		}

		begin := uint64(0)
		end := uint64(len(divisible)) - 1

		distance, remainder := evenlyDistance(begin, end, uint64(quantity))

		give := func(_ uint64, sp Span[uint64]) bool {
			return yield(divisible[sp.Begin : sp.End+1])
		}

		evenInc(begin, end, distance, remainder, give)
	}

	return iterator
}

func evenlyDistance[Type constraints.Integer](begin, end, quantity Type) (Type, Type) {
	if begin < end {
		// Overflow and other errors is not possible with these operations given the
		// checks on the values of the quantity argument located above in the calling
		// function
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

func evenDec[Type constraints.Integer](
	begin Type,
	end Type,
	distance Type,
	remainder Type,
	yield func(uint64, Span[Type]) bool,
) {
	id := uint64(0)

	for spanBegin, spanEnd := begin, begin-distance+1; ; {
		if remainder != 0 {
			spanEnd--
			remainder--
		}

		span := Span[Type]{
			Begin: spanBegin,
			End:   spanEnd,
		}

		if !yield(id, span) {
			return
		}

		if spanEnd == end {
			return
		}

		id++
		spanBegin = spanEnd - 1
		spanEnd -= distance
	}
}

func evenInc[Type constraints.Integer](
	begin Type,
	end Type,
	distance Type,
	remainder Type,
	yield func(uint64, Span[Type]) bool,
) {
	id := uint64(0)

	for spanBegin, spanEnd := begin, begin+distance-1; ; {
		if remainder != 0 {
			spanEnd++
			remainder--
		}

		span := Span[Type]{
			Begin: spanBegin,
			End:   spanEnd,
		}

		if !yield(id, span) {
			return
		}

		if spanEnd == end {
			return
		}

		id++
		spanBegin = spanEnd + 1
		spanEnd += distance
	}
}
