package main

import (
	"fmt"
	"os"
)

// printHelp prints the help message with exact spacing requirements using fmt.
func printHelp() {
	// Short flag: two spaces before the (-)
	// Explanation: tab followed by a space before the beginning of the sentence
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")

	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

// getFlagValue extracts the value after the '=' sign in a flag string.
func getFlagValue(arg string) string {
	for i, r := range arg {
		if r == '=' {
			// Return the substring after the '='
			return arg[i+1:]
		}
	}
	return ""
}

// sortString implements a simple bubble sort to order runes in a string (ASCII order).
func sortString(s string) string {
	runes := []rune(s)
	n := len(runes)
	// Bubble Sort
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
	return string(runes)
}

func main() {
	args := os.Args[1:]

	// 1. Check for no arguments
	if len(args) == 0 {
		printHelp()
		return
	}

	insertValue := ""
	isOrder := false
	mainString := ""
	nonFlagArgs := []string{}

	for _, arg := range args {
		// 2. Check for help flags
		if arg == "--help" || arg == "-h" {
			printHelp()
			return
		}

		// 3. Check for order flags
		if arg == "--order" || arg == "-o" {
			isOrder = true
			continue
		}

		// 4. Check for insert flags
		if len(arg) >= 9 && arg[:9] == "--insert=" {
			insertValue = getFlagValue(arg)
			continue
		}
		if len(arg) >= 3 && arg[:3] == "-i=" {
			insertValue = getFlagValue(arg)
			continue
		}

		// 5. Everything else is a non-flag argument.
		nonFlagArgs = append(nonFlagArgs, arg)
	}

	// The string to be operated on is the last non-flag argument provided.
	if len(nonFlagArgs) > 0 {
		mainString = nonFlagArgs[len(nonFlagArgs)-1]
	}

	// --- Execution Logic ---

	// 1. Apply insertion: Insert value is concatenated to the end of the main string
	result := mainString + insertValue

	// 2. Apply ordering
	if isOrder {
		result = sortString(result)
	}

	// 3. Print the result using fmt.Println
	if len(result) > 0 {
		fmt.Println(result)
	}
}
