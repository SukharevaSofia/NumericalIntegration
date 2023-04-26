package main

import "math"

func simpson(n int, left, right, accuracy float64, eqt equation) (float64, float64) {
	var sum, prevSum float64
	if n%2 == 1 {
		n += 1
	}
	n /= 2
	sum, prevSum = 100, 2

	for math.Abs(sum-prevSum) > math.Pow(10, -accuracy) {
		n *= 2
		prevSum = sum
		sum = 0
		h := (right - left) / float64(n)
		for i := left + h; i < right; i += 2 * h {
			sum += 4 * countEqt(i, eqt)
		}

		for i := left + 2*h; i < right-h; i += 2 * h {
			sum += 2 * countEqt(i, eqt)
		}
		sum += countEqt(left, eqt) + countEqt(right, eqt)
		sum *= h / 3
	}

	return round(sum, accuracy), math.Log2(float64(n))

}
