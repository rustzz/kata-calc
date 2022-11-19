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
		output, err := calculate(splittedLine)

		if err != nil {
			logger.Warn(err)
			continue
		}
		fmt.Println("Результат: ", output)
	}
}
