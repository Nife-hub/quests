package main

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for i := 0; i < len(a); i++ {
		for _, r := range a[i] {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
