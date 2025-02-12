package span

import "cmp"

// Describes span.
type Span[Type cmp.Ordered] struct {
	Begin Type
	End   Type
}
