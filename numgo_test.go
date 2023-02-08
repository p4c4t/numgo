package numgo_test

import (
	"numgo"
	"testing"
)

func TestRand(test *testing.T) {
	a, b := numgo.Make[int64](5), numgo.Make[int64](5)
	c, d := numgo.SmRandFill(&a), numgo.RandFill(&b)
	numgo.Print(*c)
	numgo.Print(*d)
}

func TestAdd(test *testing.T) {
	a, b := numgo.Make[int64](5), numgo.Make[int64](5)
	c, d := numgo.SmRandFill(&a), numgo.SmRandFill(&b)
	numgo.Print(*c)
	numgo.Print(*d)
	numgo.Print(*numgo.Add(c, d))
}

func TestMult(test *testing.T) {
	a, b := numgo.Make[int64](5), numgo.Make[int64](5)
	c, d := numgo.SmRandFill(&a), numgo.SmRandFill(&b)
	numgo.Print(*c)
	numgo.Print(*d)
	numgo.Print(*numgo.Mult(c, d))
}

func TestNeg(test *testing.T) {
	a := numgo.Make[int64](5)
	c := numgo.SmRandFill(&a)
	d := numgo.Neg(c)
	numgo.Print(*c)
	numgo.Print(*d)
	numgo.Print(*numgo.Add(c, d))
}

func TestScal(test *testing.T) {
	a := numgo.Make[int64](5)
	c := numgo.SmRandFill(&a)
	d := numgo.Scale(c, -2)
	numgo.Print(*c)
	numgo.Print(*d)
	numgo.Print(*numgo.Add(c, d))
}

func BenchmarkAddParallel20(bench *testing.B) {
	a, b := numgo.Make[int64](1e8), numgo.Make[int64](1e8)
	numgo.Add(&a, &b)
}
