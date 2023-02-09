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

type constr_number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

func random[T constr_number]() T {
	return T(rand.Int())
}

type ngarray[T constr_number] struct {
	length uint
	data   []T
}

func Make[T constr_number](length uint) ngarray[T] {
	return ngarray[T]{
		length: length,
		data:   make([]T, length),
	}
}

func Print[T constr_number](a ngarray[T]) {
	for _, x := range a.data {
		print(x, " ")
	}
	println()
}

type BinaryByElemFunc[T constr_number] func(T, T) T

func BinOperateSlow[T constr_number](res, a, b *[]T, start, end uint, op BinaryByElemFunc[T], wg *sync.WaitGroup) {
	defer wg.Done()

	for i := uint(start); i < uint(end); i++ {
		(*res)[i] = op((*a)[i], (*b)[i])
	}
}
func BinOperateParallel[T constr_number](a, b *ngarray[T], op BinaryByElemFunc[T], threads uint) *ngarray[T] {
	assert(
		a.length == b.length,
		"Arrays must have the same length",
	)
	assert(
		threads != 0,
		"we can't just do nothing!",
	)
	n := a.length
	res := Make[T](n)
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
func Add[T constr_number](a, b *ngarray[T]) *ngarray[T] {
	return BinOperateParallel(
		a, b,
		func(a, b T) T {
			return a + b
		},
		20,
	)
}
func Mult[T constr_number](a, b *ngarray[T]) *ngarray[T] {
	return BinOperateParallel(
		a, b,
		func(a, b T) T {
			return a * b
		},
		20,
	)
}

type UnaryByElemFunc[T constr_number] func(T) T

func UnOperateSlow[T constr_number](res, a *[]T, start, end uint, op UnaryByElemFunc[T], wg *sync.WaitGroup) {
	defer wg.Done()

	for i := uint(start); i < uint(end); i++ {
		(*res)[i] = op((*a)[i])
	}
}
func UnOperateParallel[T constr_number](a *ngarray[T], op UnaryByElemFunc[T], threads uint) *ngarray[T] {
	assert(
		threads != 0,
		"we can't just do nothing!",
	)
	n := a.length
	res := Make[T](n)
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

func RandFill[T constr_number](a *ngarray[T]) *ngarray[T] {
	return UnOperateParallel(
		a,
		func(x T) T {
			return random[T]()
		},
		20,
	)
}
func SmRandFill[T constr_number](a *ngarray[T]) *ngarray[T] {
	return UnOperateParallel(
		a,
		func(x T) T {
			return random[T]()
		},
		20,
	)
}
func Neg[T constr_number](a *ngarray[T]) *ngarray[T] {
	return UnOperateParallel(
		a,
		func(x T) T {
			return -x
		},
		20,
	)
}
func Scale[T constr_number](a *ngarray[T], x T) *ngarray[T] {
	return UnOperateParallel(
		a,
		func(y T) T {
			return y * x
		},
		20,
	)
}

func Get[T constr_number](a *ngarray[T], i uint) T {
	return a.data[i]
}
func Set[T constr_number](a *ngarray[T], i uint, val T) {
	a.data[i] = val
}

func SumSlow[T constr_number](a *[]T, start, end uint, ch chan T) {
	sum := T(0)
	for i := uint(start); i < uint(end); i++ {
		sum += (*a)[i]
	}
	ch <- sum
}

func Sum[T constr_number](a *ngarray[T]) T {
	n := a.length
	threads := uint(20)
	sum_accum := make(chan T)
	sum := T(0)
	for i := uint(0); i < threads; i++ {
		go SumSlow(
			&a.data,
			(n*i)/threads, (n*(i+1))/threads,
			sum_accum,
		)
	}
	for i := uint(0); i < threads; i++ {
		sum += <-sum_accum
	}
	return sum
}
