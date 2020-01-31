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

	var numberOfTests, _ = strconv.Atoi(inputs[0])

	//var words []string = strings.Split(inputs[1], "-")

	// 3
	// Simon says raise your right hand.
	// Lower your right hand.
	// Simon says raise your left hand.

	for i := 0; i <= numberOfTests; i++ {
		var simon []string = strings.Split(inputs[i], "Simon says ")

		if len(simon) == 2 {
			fmt.Println(simon[1])
		}
	}
}
