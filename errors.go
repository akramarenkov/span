package span

import "errors"

var (
	ErrSpansDiffSequencing       = errors.New("spans have different types of sequences")
	ErrSpansDiscontinuous        = errors.New("spans is discontinuous")
	ErrSpansIntersect            = errors.New("spans is intersect")
	ErrSpansQuantityNegative     = errors.New("spans quantity is negative")
	ErrSpansQuantityZero         = errors.New("spans quantity is zero")
	ErrSpansUnexpectedSequencing = errors.New("spans sequence type is unexpected")
	ErrSpanWidthNegative         = errors.New("span width is negative")
	ErrSpanWidthZero             = errors.New("span width is zero")
)
