package main

import "fmt"

func main() {
	var pc [256]byte
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		if i%10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%b = %b + %b\n", pc[i], i/2, byte(i&1))
	}

	var v, x uint64
	for x = 0; x < 16; x++ {
		for v = 0; v < 8; v++ {
			fmt.Printf("x = %v = %b; x>>(%v*8) = %b; pc[val] = %b\n",
				x, x, v, byte(x>>(v*8)), pc[byte(x>>(0*8))])
		}
	}
}
