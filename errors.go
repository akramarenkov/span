package span

import "errors"

var (
	ErrSpansQuantityNegative = errors.New("spans quantity is negative")
	ErrSpansQuantityZero     = errors.New("spans quantity is zero")
	ErrSpanWidthNegative     = errors.New("span width is negative")
	ErrSpanWidthZero         = errors.New("span width is zero")
)
