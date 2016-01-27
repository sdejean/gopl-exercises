// echoex11 - echo exercise 1.1 from chapter 1 of "The Go Programming Language"
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, ` `))
}
