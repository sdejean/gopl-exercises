// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

const (
	AbsoluteZeroK Kelvin = 0
	FreezingK     Kelvin = 273.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

// CToF converts a Celsius temperature to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK converts a Celsius temperature to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

// FToC converts a Fahrenheit temperature to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK converts a Fahrenheit temperature to Kelvin
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }

// KToC converts a Kelvin temperature to Celsius
func KToC(k Kelvin) Celsius { return Celsius(k - FreezingK) }

// KToF converts a Kelvin temperature to Fahrenheit
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }
