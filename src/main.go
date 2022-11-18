package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение: ")
	text, _ := reader.ReadString('\n')

	fmt.Println(executeExpression(replaceRoman(text)))
}
