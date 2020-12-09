package tempconv

import "fmt"

// 基底型(underlying-type)に名前をつける
// type name underlying-type
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC             = 0
	BoilingC              = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
