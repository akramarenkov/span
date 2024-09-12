package span

import "golang.org/x/exp/constraints"

// Describes span.
type Span[Type constraints.Ordered] struct {
	Begin Type
	End   Type
}
