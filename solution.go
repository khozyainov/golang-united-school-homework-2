package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type sidesNumber int

const (
	SidesCircle   sidesNumber = 0
	SidesTriangle sidesNumber = 3
	SidesSquare   sidesNumber = 4
)

func CalcSquare(sideLen float64, sidesNum sidesNumber) float64 {
	var square float64
	switch sidesNum {
	case SidesCircle:
		square = math.Pi * sideLen * sideLen
	case SidesTriangle:
		square = sideLen * sideLen / 2
	case SidesSquare:
		square = sideLen * sideLen
	}
	return square
}
