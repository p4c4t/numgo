package numgo

import (
	"testing"
)

func BenchmarkAddSlow(bench *testing.B) {
	a, b, res := Make(1e8), Make(1e8), Make(1e8)
	AddSlow(&res.data, &a.data, &b.data, 0, 1e8)
}

func BenchmarkAddParallel2(bench *testing.B) {
	a, b, res := Make(1e8), Make(1e8), Make(1e8)
	AddParallel(&res.data, &a.data, &b.data, 1e8, 2)
}

func BenchmarkAddParallel4(bench *testing.B) {
	a, b, res := Make(1e8), Make(1e8), Make(1e8)
	AddParallel(&res.data, &a.data, &b.data, 1e8, 4)
}

func BenchmarkAddParallel8(bench *testing.B) {
	a, b, res := Make(1e8), Make(1e8), Make(1e8)
	AddParallel(&res.data, &a.data, &b.data, 1e8, 8)
}

func BenchmarkAddParallel16(bench *testing.B) {
	a, b, res := Make(1e8), Make(1e8), Make(1e8)
	AddParallel(&res.data, &a.data, &b.data, 1e8, 16)
}

func BenchmarkAddParallel32(bench *testing.B) {
	a, b, res := Make(1e8), Make(1e8), Make(1e8)
	AddParallel(&res.data, &a.data, &b.data, 1e8, 32)
}

func BenchmarkAddParallel64(bench *testing.B) {
	a, b, res := Make(1e8), Make(1e8), Make(1e8)
	AddParallel(&res.data, &a.data, &b.data, 1e8, 64)
}

func BenchmarkAddParallel128(bench *testing.B) {
	a, b, res := Make(1e8), Make(1e8), Make(1e8)
	AddParallel(&res.data, &a.data, &b.data, 1e8, 128)
}
