package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Введите выражение: ")
		line, _ := reader.ReadString('\n')
		splittedLine := strings.Split(line, " ")
		for index, _ := range splittedLine {
			splittedLine[index] = strings.TrimSpace(splittedLine[index])
		}

		if len(splittedLine) != INPUT_DATA {
			fmt.Println("Введенные данные не удовлетворяет условию: 2 числа и знак арифметической операции в одну строку")
			continue
		}

		if isArabic(splittedLine[ORDINAL_FIRST_NUMBER]) && isArabic(splittedLine[ORDINAL_SECOND_NUMBER]) {
			fmt.Println("dawdawdaw")
			fmt.Println(performAnAction(splittedLine[ORDINAL_FIRST_NUMBER], splittedLine[OPERATION_SIGN], splittedLine[ORDINAL_SECOND_NUMBER]))
			continue
		}
		if !isArabic(splittedLine[ORDINAL_FIRST_NUMBER]) && !isArabic(splittedLine[ORDINAL_SECOND_NUMBER]) {
			var resultArabic int = performAnAction(romeToArabic(splittedLine[ORDINAL_FIRST_NUMBER]), splittedLine[OPERATION_SIGN], romeToArabic(splittedLine[ORDINAL_SECOND_NUMBER]))
			if resultArabic < 1 {
				fmt.Println("В римской системе нет чисел меньше 1")
			}
			fmt.Println(arabicToRome(resultArabic))
			continue
		}
		if !(isArabic(splittedLine[ORDINAL_FIRST_NUMBER]) && isArabic(splittedLine[ORDINAL_SECOND_NUMBER])) {
			fmt.Println("Используются одновременно разные системы счисления")
		}
	}
}
