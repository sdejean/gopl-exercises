/*
 * Conv, a program to convert between SI and Imperial units of length, mass
 * and temperature.
 */
package main

import "fmt"
import "os"
import "strconv"

import "../unitconv/lengthconv"
import "../unitconv/tempconv"

func main() {

	args := os.Args[1:]
	fmt.Printf("Units of length:\n")
	printLenConv(args)
	fmt.Printf("Units of mass:\n")
	printMassConv(args)
	fmt.Printf("Units of temperature:\n")
	printTempConv(args)
}

// printLenConv
func printLenConv(args []string) {
	for _, arg := range args {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cfk: %v\n", err)
			os.Exit(1)
		}
		m := lengthconv.Metre(t)
		y := lengthconv.Yard(t)
		f := lengthconv.Foot(t)
		// i := lengthconv.Inch(t)
		fmt.Printf("%s = %s, %s = %s, %s = %s\n",
			m, lengthconv.MToY(m), m, lengthconv.MToF(m), m, lengthconv.MToI(m))
		fmt.Printf("%s = %s, %s = %s, %s = %s\n",
			y, lengthconv.YToM(y), y, lengthconv.YToF(y), y, lengthconv.YToI(y))
		fmt.Printf("%s = %s, %s = %s, %s = %s\n",
			f, lengthconv.FToM(f), f, lengthconv.FToY(f), f, lengthconv.FToI(f))
		// fmt.Printf("%s = %s, %s = %s, %s = %s\n",
		//	i, lengthconv.IToM(i), i, lengthconv.IToY(i), i, lengthconv.IToF(i))
	}
}

// printMassConv
func printMassConv(args []string) {
	fmt.Println("TODO: MASS CONVERSIONS")
}

// printTempConv
func printTempConv(args []string) {
	for _, arg := range args {
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
