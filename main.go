package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var eq1, eq2, eq3, eq equation
	eq1.cube, eq1.sqr, eq1.fst, eq1.zer = 2, -4, 6, -40
	eq2.cube, eq2.sqr, eq2.fst, eq2.zer = 1, -3, 6, -19
	eq3.cube, eq3.sqr, eq3.fst, eq3.zer = -3, -2, 5, 39

	fNumber, left, right, accuracy := GetValues(eq1, eq2, eq3)

	switch fNumber {
	case 1:
		eq = eq1
	case 2:
		eq = eq2
	case 3:
		eq = eq3
	default:
		fmt.Println("Непредвиденная ошибка!")
		os.Exit(0)
	}

	fmt.Printf("Для функции %d:\nЛевый предел: %f, правый предел:%f, точность %d знака\n", fNumber, left, right, int(accuracy))
	answL, iterL := rectLeft(4, left, right, accuracy, eq)
	answR, iterR := rectRight(4, left, right, accuracy, eq)
	answM, iterM := rectMid(4, left, right, accuracy, eq)
	answT, iterT := trapeze(4, left, right, accuracy, eq)
	answS, iterS := simpson(4, left, right, accuracy, eq)
	fmt.Printf("Прямоугольники:\n-левые: %s, %d итераций\n-правые: %s, %d итераций\n-средние: %s, %d итераций\n",
		strconv.FormatFloat(answL, 'f', int(accuracy), 64), int(iterL),
		strconv.FormatFloat(answR, 'f', int(accuracy), 64), int(iterR),
		strconv.FormatFloat(answM, 'f', int(accuracy), 64), int(iterM))
	fmt.Printf("Трапеции: %s, %d итераций\n", strconv.FormatFloat(answT, 'f', int(accuracy), 64), int(iterT))
	fmt.Printf("Метод Симпсона: %s, %d итераций\n", strconv.FormatFloat(answS, 'f', int(accuracy), 64), int(iterS))

}
