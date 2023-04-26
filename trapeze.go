package main

import "math"

func trapeze(n int, left, right, accuracy float64, eqt equation) (float64, float64) {
	var sum, prevSum float64
	n /= 2
	sum, prevSum = 100, 2

	for math.Abs(sum-prevSum) > math.Pow(10, -accuracy) {
		n *= 2
		prevSum = sum
		sum = 0
		h := (right - left) / float64(n)
		for i := left + h; i < right; i += h {
			sum += countEqt(i, eqt)
		}
		sum += (countEqt(left, eqt) + countEqt(right, eqt)) / 2
		sum *= h
	}

	return round(sum, accuracy), math.Log2(float64(n))
}
