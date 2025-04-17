package span_test

import (
	"fmt"

	"github.com/akramarenkov/span"
)

func ExampleLinear() {
	spans, err := span.Linear(1, 7, 2)
	fmt.Println(err)
	fmt.Println(spans)
	// Output:
	// <nil>
	// [{1 2} {3 4} {5 6} {7 7}]
}
