package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput() []string {
	var inputs []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var input string = scanner.Text()

		inputs = append(inputs, input)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error encountered:", err)
	}

	return (inputs)
}

func main() {
	var inputs []string = readInput()
	var result string = ""

	var words []string = strings.Split(inputs[0], "-")

	for _, val := range words {

		result = result + strings.Split(val, "")[0]
	}

	fmt.Println(result)
}
