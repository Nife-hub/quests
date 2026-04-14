package main

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// collect all digits into a slice
	var digits []int							
	for n > 0 {									
		digits = append(digits, n%10)
		n = n / 10
	}

	// sort the digits (simple ascending order)
	for i := 0; i < len(digits); i++ {
		for j := i + 1; j < len(digits); j++ {
			if digits[i] > digits[j] {
				digits[i], digits[j] = digits[j], digits[i]
			}
		}
	}

	// print each digit using z01.PrintRune
	for _, d := range digits {
		z01.PrintRune(rune(d + '0'))
	}
}
