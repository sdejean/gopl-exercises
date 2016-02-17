// gopl, ch2, exercise 2.3 popcount-loop
package popcountLoop

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var t byte
	for i := uint64(0); i < 8; i++ {
		t += pc[byte(x>>(i*8))]
	}
	return int(t)
}
