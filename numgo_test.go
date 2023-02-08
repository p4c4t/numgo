package numgo

import (
	"testing"
)

func BenchmarkAddParallel20(bench *testing.B) {
	a, b := Make(1e8), Make(1e8)
	Add(&a, &b)
}
