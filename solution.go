package square

import (
	"math"
)

type CountSides int

const (
	SidesCircle CountSides = iota
	_
	_
	SidesTriangle = iota
	SidesSquare = iota
)

func CalcSquare(sideLen float64, sidesNum CountSides) float64 {
	switch sidesNum {
		case 0:
			return math.Pi * math.Pow(sideLen, 2)
		case 3:
			return (math.Pow(sideLen, 2) * math.Sqrt(3)) / 4
		case 4:
			return math.Pow(sideLen, 2)
		default:
			return 0
		}
}