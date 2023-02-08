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

func AddSlow(res, a, b *[]int64, start, end uint) {
	for i := uint(start); i < uint(end); i++ {
		(*res)[i] = (*a)[i] + (*b)[i]
	}
}

func AddParallel(res, a, b *[]int64, n uint, threads uint) {
	assert(threads != 0, "we can't just do nothing!")
	for i := uint(0); i < threads; i++ {
		go AddSlow(
			res,
			a, b,
			(n*i)/threads, (n*(i+1))/threads,
		)
	}
}

func Add(a, b *Int64Array) Int64Array {
	assert(
		a.length == b.length,
		"Arrays must have the same length",
	)
	n := a.length
	c := Int64Array{
		length: n,
		data:   make([]int64, n),
	}
	AddParallel(&c.data, &a.data, &b.data, n, 8)
	return c
}
