package span

import "golang.org/x/exp/constraints"

type Span[Type constraints.Ordered] struct {
	Begin Type
	End   Type
}
