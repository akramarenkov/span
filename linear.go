package span

import (
	"github.com/akramarenkov/safe"
	"golang.org/x/exp/constraints"
)

// Divides a linear sequence of integers from begin to end inclusive into spans of the
// specified width.
//
// If begin is greater than end, the sequence will be considered decreasing,
// otherwise - increasing.
//
// If a zero or negative width of span is specified, an error is returned.
func Linear[Type constraints.Integer](begin, end, width Type) ([]Span[Type], error) {
	if width < 0 {
		return nil, ErrSpanWidthNegative
	}

	if width == 0 {
		return nil, ErrSpanWidthZero
	}

	spans := make([]Span[Type], safe.StepSize(begin, end, width))

	if begin < end {
		for id, spanBegin := range safe.Step(begin, end, width) {
			spanEnd, err := safe.Add(spanBegin, width-1)
			if err != nil {
				spanEnd = end
			}

			if spanEnd > end {
				spanEnd = end
			}

			span := Span[Type]{
				Begin: spanBegin,
				End:   spanEnd,
			}

			spans[id] = span
		}

		return spans, nil
	}

	for id, spanBegin := range safe.Step(begin, end, width) {
		spanEnd, err := safe.Sub(spanBegin, width-1)
		if err != nil {
			spanEnd = end
		}

		if spanEnd < end {
			spanEnd = end
		}

		span := Span[Type]{
			Begin: spanBegin,
			End:   spanEnd,
		}

		spans[id] = span
	}

	return spans, nil
}
