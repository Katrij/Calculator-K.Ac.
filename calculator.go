package main

import (
	"fmt"
	"os"
)

var N1, N2, to_do, t1, t2 string
var n1, n2 int
var superfluous = ""

func main() {

	fmt.Print("Пример: ")
	fmt.Fscanln(os.Stdin, &N1, &to_do, &N2, &superfluous)

	if superfluous != "" {
		panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор")
	}

	n1, t1 = Translate(N1)
	n2, t2 = Translate(N2)

	if t1 == t2 {
		switch to_do {
		case "+":
			n1 = n1 + n2
			fmt.Print("Ответ: ", AnserTranslate(n1, t1))
		case "-":
			n1 = n1 - n2
			fmt.Print("Ответ: ", AnserTranslate(n1, t1))
		case "*":
			n2 = n1 * n2
			fmt.Print("Ответ: ", AnserTranslate(n2, t2))
		case "/", ":":
			n2 = n1 / n2
			fmt.Print("Ответ: ", AnserTranslate(n2, t2))
		default:
			panic("Неизвестная арифметическая операция")
		}
	} else {
		panic("Разные системы счисления")
	}
}

func Translate(x string) (int, string) {
	var y int
	var T string
	switch x {
	case "I":
		y = 1
		T = "Rim"
	case "II":
		y = 2
		T = "Rim"
	case "III":
		y = 3
		T = "Rim"
	case "IV":
		y = 4
		T = "Rim"
	case "V":
		y = 5
		T = "Rim"
	case "VI":
		y = 6
		T = "Rim"
	case "VII":
		y = 7
		T = "Rim"
	case "VIII":
		y = 8
		T = "Rim"
	case "IX":
		y = 9
		T = "Rim"
	case "X":
		y = 10
		T = "Rim"
	default:
		switch x {
		case "1":
			y = 1
			T = "Arab"
		case "2":
			y = 2
			T = "Arab"
		case "3":
			y = 3
			T = "Arab"
		case "4":
			y = 4
			T = "Arab"
		case "5":
			y = 5
			T = "Arab"
		case "6":
			y = 6
			T = "Arab"
		case "7":
			y = 7
			T = "Arab"
		case "8":
			y = 8
			T = "Arab"
		case "9":
			y = 9
			T = "Arab"
		case "10":
			y = 10
			T = "Arab"
		default:
			panic("Не является математической операцией")
		}
	}
	return y, T
}

func AnserTranslate(x int, y string) int {
	if x > 0 && y == "Rim" {
		fmt.Print("Ответ: ")
		for x > 0 {
			if x/100 == 1 {
				fmt.Print("C")
				x = x - 100
			} else if x/90 >= 1 {
				fmt.Print("XC")
				x = x - 90
			} else if x/50 >= 1 {
				fmt.Print("L")
				x = x - 50
			} else if x/40 >= 1 {
				fmt.Print("XL")
				x = x - 40
			} else if x/10 >= 1 {
				fmt.Print("X")
				x = x - 10
			} else if x/9 >= 1 {
				fmt.Print("IX")
				x = x - 9
			} else if x/5 >= 1 {
				fmt.Print("V")
				x = x - 5
			} else if x/4 >= 1 {
				fmt.Print("IV")
				x = x - 4
			} else if x/1 >= 1 {
				fmt.Print("I")
				x = x - 1
			} else {
				fmt.Println("Ответ не переводится!")
			}
		}
		os.Exit(0)

	} else if y == "Arab" {

	} else {
		panic("В римской системе нет отрицательных чисел")
	}
	return x
}
