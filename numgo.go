package numgo

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

type BinaryByElemFunc func(int64, int64) int64

func BinOperateSlow(res, a, b *[]int64, start, end uint, op BinaryByElemFunc) {
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
	for i := uint(0); i < threads; i++ {
		go BinOperateSlow(
			&res.data,
			&a.data, &b.data,
			(n*i)/threads, (n*(i+1))/threads,
			op,
		)
	}
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

func Dot(a, b *Int64Array) *Int64Array {
	return BinOperateParallel(
		a, b,
		func(a, b int64) int64 {
			return a * b
		},
		20,
	)
}
