// Exercise 2.1 - Extend package tempconv to also perform Kelvin temperature
// conversions.
package tempconv

import "fmt"

type Kelvin float64

const (
	AbsoluteZeroK Kelvin = 0
	FreezingK     Kelvin = 273.15
)

func (k Kelvin) String() string { return fmt.Sprintf("%g°K", k) }

// CToK converts a Celsius temperature to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

// FToK converts a Fahrenheit temperature to Kelvin
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }

// KToC converts a Kelvin temperature to Celsius
func KToC(k Kelvin) Celsius { return Celsius(k - FreezingK) }

// KToF converts a Kelvin temperature to Fahrenheit
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }
