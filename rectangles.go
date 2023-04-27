package main

import (
	"math"
)

func rectLeft(n int, left, right, accuracy float64, eqt equation) (float64, int, int) {
	var sum, prevSum float64
	n /= 2
	sum, prevSum = 100, 2
	it := 0
	for runge(sum, prevSum, 1) > math.Pow(10, -accuracy) {
		n *= 2
		prevSum = sum
		sum = 0
		h := (right - left) / float64(n)
		for i := left; i < right; i += h {
			sum += countEqt(i, eqt)
		}
		sum *= h
		it++
	}

	return round(sum, accuracy), n, it
}

func rectRight(n int, left, right, accuracy float64, eqt equation) (float64, int, int) {
	var sum, prevSum float64
	n /= 2
	sum, prevSum = 100, 2
	it := 0
	for runge(sum, prevSum, 1) > math.Pow(10, -accuracy) {
		n *= 2
		prevSum = sum
		sum = 0
		h := (right - left) / float64(n)
		for i := left + h; i <= right; i += h {
			sum += countEqt(i, eqt)
		}
		sum *= h
		it++
	}

	return round(sum, accuracy), n, it
}

func rectMid(n int, left, right, accuracy float64, eqt equation) (float64, int, int) {
	var sum, prevSum float64
	n /= 2
	sum, prevSum = 100, 2
	it := 0
	for runge(sum, prevSum, 2) > math.Pow(10, -accuracy) {

		n *= 2
		prevSum = sum
		sum = 0
		h := (right - left) / float64(n)
		for i := left; i < right; i += h {
			sum += countEqt(i+h/2, eqt)
		}
		sum *= h
		it++
	}

	return round(sum, accuracy), n, it
}
