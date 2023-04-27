package main

import (
	"math"
)

type equation struct {
	cube, sqr, fst, zer float64
}

func round(x, acc float64) float64 {
	x = math.Round(x*math.Pow(10, acc)) / math.Pow(10, acc)
	return x
}

func countEqt(x float64, eqt equation) float64 {
	return (eqt.cube * math.Pow(x, 3)) + (eqt.sqr * math.Pow(x, 2)) + eqt.fst*x + eqt.zer
}

func runge(current, previous, accuracy float64) float64 {
	return math.Abs(current-previous) / (math.Pow(2, accuracy) - 1)
}
