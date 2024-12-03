package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// 循环
func PopCountByFor(x uint64) int {
	idx := []uint{0, 1, 2, 3, 4, 5, 6, 7}
	var r byte
	for _, v := range idx {
		r += pc[byte(x>>(v*8))]
	}
	return int(r)
}
