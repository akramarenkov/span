package span

import "errors"

var (
	ErrSpansDiffSequenceType       = errors.New("spans have different types of sequences")
	ErrSpansIntersect              = errors.New("spans is intersect")
	ErrSpansQuantityNegative       = errors.New("spans quantity is negative")
	ErrSpansQuantityZero           = errors.New("spans quantity is zero")
	ErrSpansSequenceTypeUnexpected = errors.New("spans sequence type is unexpected")
	ErrSpanWidthNegative           = errors.New("span width is negative")
	ErrSpanWidthZero               = errors.New("span width is zero")
)
