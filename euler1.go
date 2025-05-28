package eulergo

func Euler1(n uint64) uint64 {
	var x uint64 = 0
	var i uint64 = 0
	for i = 0; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			x += i
		}
	}
	return x
}
