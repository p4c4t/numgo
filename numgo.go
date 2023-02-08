package numgo

import (
	"math/rand"
	"sync"
)

func assert(b bool, errmsg string) {
	if !b {
		panic(errmsg)
	}
}

type Int64Array struct {
	length uint
	data   []int64
}

func Make(length uint) Int64Array {
	return Int64Array{
		length: length,
		data:   make([]int64, length),
	}
}

func Print(a Int64Array) {
	for _, x := range a.data {
		print(x, " ")
	}
	println()
}

type BinaryByElemFunc func(int64, int64) int64

func BinOperateSlow(res, a, b *[]int64, start, end uint, op BinaryByElemFunc, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := uint(start); i < uint(end); i++ {
		(*res)[i] = op((*a)[i], (*b)[i])
	}
}

func BinOperateParallel(a, b *Int64Array, op BinaryByElemFunc, threads uint) *Int64Array {
	assert(
		a.length == b.length,
		"Arrays must have the same length",
	)
	assert(
		threads != 0,
		"we can't just do nothing!",
	)
	n := a.length
	res := Make(n)
	var wg sync.WaitGroup
	for i := uint(0); i < threads; i++ {
		wg.Add(1)
		go BinOperateSlow(
			&res.data,
			&a.data, &b.data,
			(n*i)/threads, (n*(i+1))/threads,
			op,
			&wg,
		)
	}
	wg.Wait()
	return &res
}

func Add(a, b *Int64Array) *Int64Array {
	return BinOperateParallel(
		a, b,
		func(a, b int64) int64 {
			return a + b
		},
		20,
	)
}

func Mult(a, b *Int64Array) *Int64Array {
	return BinOperateParallel(
		a, b,
		func(a, b int64) int64 {
			return a * b
		},
		20,
	)
}

type UnaryByElemFunc func(int64) int64

func UnOperateSlow(res, a *[]int64, start, end uint, op UnaryByElemFunc, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := uint(start); i < uint(end); i++ {
		(*res)[i] = op((*a)[i])
	}
}

func UnOperateParallel(a *Int64Array, op UnaryByElemFunc, threads uint) *Int64Array {
	assert(
		threads != 0,
		"we can't just do nothing!",
	)
	n := a.length
	res := Make(n)
	var wg sync.WaitGroup
	for i := uint(0); i < threads; i++ {
		wg.Add(1)
		go UnOperateSlow(
			&res.data,
			&a.data,
			(n*i)/threads, (n*(i+1))/threads,
			op,
			&wg,
		)
	}
	wg.Wait()
	return &res
}

func RandFill(a *Int64Array) *Int64Array {
	return UnOperateParallel(
		a,
		func(x int64) int64 {
			return rand.Int63()
		},
		20,
	)
}

func SmRandFill(a *Int64Array) *Int64Array {
	return UnOperateParallel(
		a,
		func(x int64) int64 {
			return rand.Int63n(10)
		},
		20,
	)
}
