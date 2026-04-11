package main

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	for n > 0 && n < 1 {
		for i := '0'; i <= n; i++ {
			z01.PrintRune(i)
		}
	}
}