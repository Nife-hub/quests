package main

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}

	if n/10 != 0 {
		z01.PrintRune(n / 10)
	}

	digit := n % 10
	if digit < 0 {
		digit = -digit
	}

	z01.PrintRune(rune(digit + '0'))
}