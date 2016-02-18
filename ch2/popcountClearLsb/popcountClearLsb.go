// gopl, ch2, exercise 2.5 - popcountClearLsb
package popcountClearLsb

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var sum byte
	for x != 0 {
		x = x & (x - 1)
		sum += 1
	}
	return int(sum)
}
