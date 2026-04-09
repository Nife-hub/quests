package main

import "github.com/01-edu/z01"

func main() {
	for beta := 'z'; beta >= 'a'; beta-- {
		z01.PrintRune(beta)
	}
}