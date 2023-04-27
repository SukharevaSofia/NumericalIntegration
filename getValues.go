package main

import (
	"fmt"
	"math"
	"strconv"
)

func GetValues(eq1, eq2, eq3 equation) (int, float64, float64, float64) {
	s := "Лабораторная работа №3\n" +
		"Численное интегрирование\n" +
		"Подготовила Софья Сухарева\n" +
		"P32131, ISU:334641, вариант №19"
	fmt.Println(s)

	fmt.Printf("Выберите функцию для интегрирования. Введите 1, 2 или 3\n"+
		"1: %.3fx^3 + (%.3fx^2) + (%.3fx) + (%.3f)\n"+
		"2: %.3fx^3 + (%.3fx^2) + (%.3fx) + (%.3f)\n"+
		"3: %.3fx^3 + (%.3fx^2) + (%.3fx) + (%.3f)", eq1.cube, eq1.sqr, eq1.fst, eq1.zer,
		eq2.cube, eq2.sqr, eq2.fst, eq2.zer,
		eq3.cube, eq3.sqr, eq3.fst, eq3.zer)
	fmt.Print("\nВыбрать функцию: ")

	fNumber := 0
	for {
		fmt.Scan(&fNumber)
		if (fNumber == 1) || (fNumber == 2) || (fNumber == 3) {
			break
		}
		fmt.Println("Введите 1, 2 или 3")
		fmt.Print("Выбрать функцию: ")
	}

	var left, right string
	fmt.Print("Введите левый предел интегрирования: ")
	for {
		fmt.Scan(&left)
		if _, err := strconv.ParseFloat(left, 64); err == nil {
			break
		}
		fmt.Print("Предел должен являться числом: ")
	}
	fmt.Print("Введите правый предел интегрирования: ")
	for {
		fmt.Scan(&right)
		if a, err := strconv.ParseFloat(right, 64); err == nil {
			if b, _ := strconv.ParseFloat(left, 64); a > b {
				break
			}

		}
		fmt.Print("Предел должен являться числом, большим, чем левый предел: ")
	}

	var accuracy string
	fmt.Print("Введите точность вычислений: ")
	for {
		fmt.Scan(&accuracy)
		if _, err := strconv.ParseFloat(accuracy, 64); err == nil {
			break
		}
		fmt.Print("Точность должна являться числом: ")
	}
	c, _ := strconv.ParseFloat(accuracy, 64)
	var acc float64
	for math.Mod(c, 1) != 0 {
		acc++
		c *= 10
	}

	a, _ := strconv.ParseFloat(left, 64)
	b, _ := strconv.ParseFloat(right, 64)
	return fNumber, a, b, acc
}
