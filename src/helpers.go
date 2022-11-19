package main

import (
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
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
		logger.Error(err)
		os.Exit(1)
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
			logger.Error(err)
			os.Exit(1)
		}
		op2, err := strconv.Atoi(operator2)
		if err != nil {
			logger.Error(err)
			os.Exit(1)
		}
		return op1 + op2
	case ("-"):
		op1, err := strconv.Atoi(operator1)
		if err != nil {
			logger.Error(err)
			os.Exit(1)
		}
		op2, err := strconv.Atoi(operator2)
		if err != nil {
			logger.Error(err)
			os.Exit(1)
		}
		return op1 - op2
	case ("*"):
		op1, err := strconv.Atoi(operator1)
		if err != nil {
			logger.Error(err)
			os.Exit(1)
		}
		op2, err := strconv.Atoi(operator2)
		if err != nil {
			logger.Error(err)
			os.Exit(1)
		}
		return op1 * op2
	case ("/"):
		op1, err := strconv.Atoi(operator1)
		if err != nil {
			logger.Error(err)
			os.Exit(1)
		}
		op2, err := strconv.Atoi(operator2)
		if err != nil {
			logger.Error(err)
			os.Exit(1)
		}
		return op1 / op2
	default:
		panic("Неверный знак операции")
	}
}

func calculate(splittedLine []string) (interface{}, error) {
	for index := range splittedLine {
		splittedLine[index] = strings.TrimSpace(splittedLine[index])
	}
	if len(splittedLine) != INPUT_DATA {
		return 0, errors.New("введенные данные не удовлетворяет условию: 2 числа и знак арифметической операции в одну строку")
	}
	if isArabic(splittedLine[ORDINAL_FIRST_NUMBER]) && isArabic(splittedLine[ORDINAL_SECOND_NUMBER]) {
		return performAnAction(splittedLine[ORDINAL_FIRST_NUMBER], splittedLine[OPERATION_SIGN], splittedLine[ORDINAL_SECOND_NUMBER]), nil
	}
	if !isArabic(splittedLine[ORDINAL_FIRST_NUMBER]) && !isArabic(splittedLine[ORDINAL_SECOND_NUMBER]) {
		var resultArabic int = performAnAction(romeToArabic(splittedLine[ORDINAL_FIRST_NUMBER]), splittedLine[OPERATION_SIGN], romeToArabic(splittedLine[ORDINAL_SECOND_NUMBER]))
		if resultArabic < 1 {
			return 0, errors.New("в римской системе нет чисел меньше 1")
		}
		return arabicToRome(resultArabic), nil
	}
	if !(isArabic(splittedLine[ORDINAL_FIRST_NUMBER]) && isArabic(splittedLine[ORDINAL_SECOND_NUMBER])) {
		return 0, errors.New("используются одновременно разные системы счисления")
	}
	return 0, errors.New("ни одно из условий проверки не подошло")
}
