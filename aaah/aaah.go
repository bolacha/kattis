package main

import (
	"bufio"
	"fmt"
	"os"
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

	var pacient string = inputs[0]
	var doctor string = inputs[1]

	// aaah     pacient
	// aaaaah	doctor
	// no

	// aaah		pacient
	// ah		doctor
	// go

	if len(doctor) > len(pacient) {
		fmt.Println("no")
	} else if len(doctor) == len(pacient) {
		fmt.Println("go")
	} else {
		fmt.Println("go")
	}
}
