// echoex12 - echo exercise 1.2 from chapter 1 of "The Go Programming Language"
package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, arg := range os.Args[1:] {
		fmt.Println(idx)
		fmt.Println(arg)
	}
}
