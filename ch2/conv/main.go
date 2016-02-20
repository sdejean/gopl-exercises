/*
 * Conv, a program to convert between SI and Imperial units of length, mass
 * and temperature.
 */
package main

import "fmt"
import "os"
import "strconv"

import "github.com/sdejean/gopl-exercises/ch2/unitconv/lengthconv"
import "github.com/sdejean/gopl-exercises/ch2/unitconv/massconv"
import "github.com/sdejean/gopl-exercises/ch2/unitconv/tempconv"

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
	for _, arg := range args {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cfk: %v\n", err)
			os.Exit(1)
		}
		k := massconv.Kilogram(t)
		o := massconv.Ounce(t)
		p := massconv.Pound(t)
		s := massconv.Stone(t)
		fmt.Printf("%s = %s, %s = %s, %s = %s\n",
			k, massconv.KToO(k), k, massconv.KToP(k), k, massconv.KToS(k))
		fmt.Printf("%s = %s, %s = %s, %s = %s\n",
			o, massconv.OToK(o), o, massconv.OToP(o), o, massconv.OToS(o))
		fmt.Printf("%s = %s, %s = %s, %s = %s\n",
			p, massconv.PToK(p), p, massconv.PToO(p), p, massconv.PToS(p))
		fmt.Printf("%s = %s, %s = %s, %s = %s\n",
			s, massconv.SToK(s), s, massconv.SToO(s), s, massconv.SToP(s))
	}
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
