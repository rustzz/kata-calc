package main

import (
	"regexp"
	"strconv"
)

func romeToArabic(romeNumber string) string {
	var romeNumberList [10]string = [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	var arabicNumberList [10]string = [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	var arabicNumber string
	for i := 0; i < len(romeNumberList); i++ {
		if romeNumber == romeNumberList[i] {
			arabicNumber += arabicNumberList[i]
			break
		}
	}
	return arabicNumber
}

func arabicToRome(arabicNumber int) string {
	var listRomeNumbers [9]string = [9]string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var listArabicNumbers [9]int = [9]int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	var count int = 0
	var remainder int
	var romeNumber string

	for count < len(listRomeNumbers) {
		for arabicNumber >= listArabicNumbers[count] {
			remainder = arabicNumber / listArabicNumbers[count]
			arabicNumber %= listArabicNumbers[count]
			for i := 0; i < remainder; i++ {
				romeNumber += listRomeNumbers[count]
			}
		}
		count++
	}
	return romeNumber
}

func isArabic(operand string) bool {
	matched, err := regexp.MatchString("\\d", operand)
	if err != nil {
		panic(err)
	}
	if matched {
		return true
	}
	return false
}

func performAnAction(operator1 string, operations string, operator2 string) int {
	switch operations {
	case ("+"):
		op1, err := strconv.Atoi(operator1)
		if err != nil {
			panic(err)
		}
		op2, err := strconv.Atoi(operator2)
		if err != nil {
			panic(err)
		}
		return op1 + op2
	case ("-"):
		op1, err := strconv.Atoi(operator1)
		if err != nil {
			panic(err)
		}
		op2, err := strconv.Atoi(operator2)
		if err != nil {
			panic(err)
		}
		return op1 - op2
	case ("*"):
		op1, err := strconv.Atoi(operator1)
		if err != nil {
			panic(err)
		}
		op2, err := strconv.Atoi(operator2)
		if err != nil {
			panic(err)
		}
		return op1 * op2
	case ("/"):
		op1, err := strconv.Atoi(operator1)
		if err != nil {
			panic(err)
		}
		op2, err := strconv.Atoi(operator2)
		if err != nil {
			panic(err)
		}
		return op1 / op2
	default:
		panic("Неверный знак операции")
	}
}
