package main

import "math"

func trapeze(n int, left, right, accuracy float64, eqt equation) (float64, int, int) {
	var sum, prevSum float64
	n /= 2
	sum, prevSum = 100, 2
	it := 0
	for runge(sum, prevSum, 2) > math.Pow(10, -accuracy) {
		n *= 2
		prevSum = sum
		sum = 0
		h := (right - left) / float64(n)
		for i := left + h; i < right; i += h {
			sum += countEqt(i, eqt)
		}
		sum += (countEqt(left, eqt) + countEqt(right, eqt)) / 2
		sum *= h
		it++
	}

	return round(sum, accuracy), n, it
}
