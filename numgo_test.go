package numgo_test

import (
	"numgo"
	"testing"
)

func TestRand(test *testing.T) {
	a, b := numgo.Make(5), numgo.Make(5)
	c, d := numgo.SmRandFill(&a), numgo.RandFill(&b)
	numgo.Print(*c)
	numgo.Print(*d)
}

func TestAdd(test *testing.T) {
	a, b := numgo.Make(5), numgo.Make(5)
	c, d := numgo.SmRandFill(&a), numgo.SmRandFill(&b)
	numgo.Print(*c)
	numgo.Print(*d)
	numgo.Print(*numgo.Add(c, d))
}

func TestMult(test *testing.T) {
	a, b := numgo.Make(5), numgo.Make(5)
	c, d := numgo.SmRandFill(&a), numgo.SmRandFill(&b)
	numgo.Print(*c)
	numgo.Print(*d)
	numgo.Print(*numgo.Mult(c, d))
}

func BenchmarkAddParallel20(bench *testing.B) {
	a, b := numgo.Make(1e8), numgo.Make(1e8)
	numgo.Add(&a, &b)
}
