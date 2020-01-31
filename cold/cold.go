package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	var nums []string = strings.Split(inputs[1], " ")
	var count int = 0

	for _, val := range nums {

		var valor, _ = strconv.Atoi(val)

		if valor < 0 {
			count++
		}
	}

	fmt.Println(count)
}
