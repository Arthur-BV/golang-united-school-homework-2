package square

import "math"

type myInt int

const (
	SidesTriangle = 3
	SidesSquare   = 4
	SidesCircle   = 0
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum myInt) float64 {
	var fSquare float64

	if sidesNum == 0 {
		fSquare = math.Pi * math.Pow(sideLen, 2)
	} else if sidesNum == 3 {
		h := (math.Sqrt(3) / 2) * sideLen
		fSquare = (h * float64(sideLen)) / 2
	} else if sidesNum == 4 {
		fSquare = math.Pow(sideLen, 2)
	}

	return fSquare
}
