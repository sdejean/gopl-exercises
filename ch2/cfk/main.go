// Cfk converts its numeric argument to Celsius, Fahrenheit and Kelvin.
package main

import (
	"fmt"
	"os"
	"strconv"

	"../tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cfk: %v\n", err)
			os.Exit(1)
		}
		c := tempconv.Celsius(t)
		f := tempconv.Fahrenheit(t)
		k := tempconv.Kelvin(t)
		fmt.Printf("%s = %s, %s = %s\n",
			c, tempconv.CToF(c), c, tempconv.CToK(c))
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), f, tempconv.FToK(f))
		fmt.Printf("%s = %s, %s = %s\n",
			k, tempconv.KToC(k), k, tempconv.KToF(k))
	}
}
