package span

import "errors"

var (
	ErrBeginGreaterEnd       = errors.New("begin is greater than end")
	ErrSpansQuantityNegative = errors.New("spans quantity is negative")
	ErrSpansQuantityZero     = errors.New("spans quantity is zero")
	ErrSpanWidthNegative     = errors.New("span width is negative")
	ErrSpanWidthZero         = errors.New("span width is zero")
)
