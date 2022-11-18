package main

import (
	"go/token"
	"go/types"
	"strconv"
	"strings"
	"unicode"
)

func romanToDecimal(romanNumber string) int {
	var decimal int
	var lastNumber int
	var romanNumeral string = strings.ToUpper(romanNumber)

	for x := len([]rune(romanNumeral)) - 1; x >= 0; x-- {
		var convertToDecimal rune = []rune(romanNumeral)[x]

		switch convertToDecimal {
		case 'M':
			decimal = processDecimal(1000, lastNumber, decimal)
			lastNumber = 1000

		case 'D':
			decimal = processDecimal(500, lastNumber, decimal)
			lastNumber = 500

		case 'C':
			decimal = processDecimal(100, lastNumber, decimal)
			lastNumber = 100

		case 'L':
			decimal = processDecimal(50, lastNumber, decimal)
			lastNumber = 50

		case 'X':
			decimal = processDecimal(10, lastNumber, decimal)
			lastNumber = 10

		case 'V':
			decimal = processDecimal(5, lastNumber, decimal)
			lastNumber = 5

		case 'I':
			decimal = processDecimal(1, lastNumber, decimal)
			lastNumber = 1
		}
	}
	return decimal
}

func processDecimal(decimal int, lastNumber int, lastDecimal int) int {
	if lastNumber > decimal {
		return lastDecimal - decimal
	}
	return lastDecimal + decimal
}

func replaceRoman(exp string) string {
	var collectRoman string
	for _, symbol := range exp {
		if unicode.IsLetter(symbol) {
			collectRoman += string(symbol)
		}
		if len([]rune(collectRoman)) != 0 && !unicode.IsLetter(symbol) {
			exp = strings.Replace(exp, collectRoman, strconv.Itoa(romanToDecimal(collectRoman)), -1)
		}
	}
	return exp
}

func executeExpression(exp string) string {
	fs := token.NewFileSet()
	tv, err := types.Eval(fs, nil, token.NoPos, exp)
	if err != nil {
		panic(err)
	}
	return tv.Value.String()
}
