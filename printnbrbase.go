package main

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	// Validate base
	if !isValidBase(base) {
		printStr("NV")
		return
	}

	// Handle negative numbers
	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}

	// Convert & print
	printInBase(nbr, base)
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)

	for _, r := range base {
		if r == '+' || r == '-' {
			return false
		}
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

// Recursively print number in the given base
func printInBase(n int, base string) {
	b := len(base)
	if n >= b {
		printInBase(n/b, base)
	}
	z01.PrintRune(rune(base[n%b]))
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
