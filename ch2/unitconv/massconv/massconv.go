// Packaget that performs conversions between SI and Imperial units of mass.
package massconv

import "fmt"

type Kilogram float64
type Ounce float64
type Pound float64
type Stone float64

const (
	KilogramToPound Kilogram = 2.2046
	PoundToKilogram Pound    = 0.4536
)

func (k Kilogram) String() string { return fmt.Sprintf("%g kg", k) }
func (o Ounce) String() string    { return fmt.Sprintf("%g oz", o) }
func (p Pound) String() string    { return fmt.Sprintf("%g lb", p) }
func (s Stone) String() string    { return fmt.Sprintf("%g st", s) }

// KToO converts a mass in Kilograms to Ounces
func KToO(k Kilogram) Ounce { return Ounce(PToO(KToP(k))) }

// KToP converts a mass in Kilograms to Pounds
func KToP(k Kilogram) Pound { return Pound(k * KilogramToPound) }

// KToS converts a mass in Kilograms to Stone
func KToS(k Kilogram) Stone { return Stone(PToS(KToP(k))) }

// OToK converts a mass in Ounces to Kilograms
func OToK(o Ounce) Kilogram { return Kilogram(PToK(OToP(o))) }

// OToP converts a mass in Ounces to Pounds
func OToP(o Ounce) Pound { return Pound(o / 16) }

// OToS converts a mass in Ounces to Stone
func OToS(o Ounce) Stone { return Stone(PToS(OToP(o))) }

// PToK converts a mass in Pounds to Kilograms
func PToK(p Pound) Kilogram { return Kilogram(p * PoundToKilogram) }

// PToO converts a mass in Pounds to Ounces
func PToO(p Pound) Ounce { return Ounce(p * 16) }

// PToS converts a mass in Pounds to Stone
func PToS(p Pound) Stone { return Stone(p / 14) }

// SToK converts a mass in Stone to Kilograms
func SToK(s Stone) Kilogram { return Kilogram(PToK(SToP(s))) }

// SToO converts a mass in Stone to Ounces
func SToO(s Stone) Ounce { return Ounce(PToO(SToP(s))) }

// SToP converts a mass in Stone to Pounds
func SToP(s Stone) Pound { return Pound(s * 14) }
