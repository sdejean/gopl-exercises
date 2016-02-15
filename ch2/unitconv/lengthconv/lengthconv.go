// Package that perform conversions between SI and Imperial units of length.
package lengthconv

import "fmt"

// TODO: add furlongs, kilometres and miles (nautical?)
type Metre float64
type Inch float64
type Foot float64
type Yard float64

const (
	MetreToYard Metre = 1.0936
	YardToMetre Yard  = 0.9144
)

func (m Metre) String() string { return fmt.Sprintf("%g m", m) }
func (i Inch) String() string  { return fmt.Sprintf("%g in", i) }
func (f Foot) String() string  { return fmt.Sprintf("%g ft", f) }
func (y Yard) String() string  { return fmt.Sprintf("%g yd", y) }

// MToF converts a length in Metres into Feet
func MToF(m Metre) Foot { return Foot(YToF(MToY(m))) }

// MToI converts a length in Metres into Inches
func MToI(m Metre) Inch { return Inch(YToI(MToY(m))) }

// MToY converts a length in Metres to Yards
func MToY(m Metre) Yard { return Yard(m * MetreToYard) }

// FToI converts a length in Feet into Inches
func FToI(f Foot) Inch { return Inch(f * 12) }

// FToM converts a length in Feet into Metres
func FToM(f Foot) Metre { return Metre(YToM(FToY(f))) }

// FToY converts a length in Feet into Yards
func FToY(f Foot) Yard { return Yard(f / 3) }

// YToF converts a length in Yards into Feet
func YToF(y Yard) Foot { return Foot(y * 3) }

// YToI converts a length in Yards into Inches
func YToI(y Yard) Inch { return Inch(y * 36) }

// YToM converts a length in Yards into Metres
func YToM(y Yard) Metre { return Metre(y * YardToMetre) }
