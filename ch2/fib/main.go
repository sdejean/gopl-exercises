// Fib - Based on an example in gopl/ch2/2.4.1 Tuple Assignment
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	i, _ := strconv.ParseUint(os.Args[1], 10, 64)
	fmt.Println(fib(i))
}

func fib(n uint64) uint64 {
	var i, x, y uint64
	x, y = 0, 1
	for i = 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
