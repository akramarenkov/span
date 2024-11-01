package span_test

import (
	"fmt"

	"github.com/akramarenkov/span"
)

func ExampleEvenly() {
	spans, err := span.Evenly(1, 8, 3)
	fmt.Println(err)
	fmt.Println(spans)
	// Output:
	// <nil>
	// [{1 3} {4 6} {7 8}]
}
