package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type sidesAmount int

const (
	Circle   = 0
	Triangle = 3
	Square   = 4
)

func CalcSquare(sideLen float64, sidesNum sidesAmount) float64 {
	switch sidesNum {
	case Circle:
		return math.Pi * math.Pow(sideLen/2, 2)
	case Triangle:
		return (math.Sqrt(3.0) / 4.0 * math.Pow(sideLen, 2))
	case Square:
		return math.Pow(sideLen, 2)
	default:
		return 0
	}
}

var (
	SidesTriangle = sidesAmount(3)
	SidesSquare   = sidesAmount(4)
	SidesCircle   = sidesAmount(0)
)
