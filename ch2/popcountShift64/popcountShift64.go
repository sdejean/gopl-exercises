// gopl, ch2, exercise 2.4 - popcountShift64
package popcountShift64

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var i uint64
	var sum byte
	for i = 0; i < 64; i++ {
		sum += pc[byte((x >> i & 1))]
	}
	return int(sum)
}
