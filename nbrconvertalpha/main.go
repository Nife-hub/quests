package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return -1
		}
		n = n*10 + int(s[i]-'0')
	}
	return n
}

func PrintAlpha(n int, upper bool) {
	if n <= 0 || n > 26 {
		z01.PrintRune(' ')
		return
	}

	if upper {
		z01.PrintRune(rune('A' + n - 1))
	} else {
		z01.PrintRune(rune('a' + n - 1))
	}
}

func main() {
	if len(os.Args) == 1 {
		return
	}

	upper := false
	start := 1

	if os.Args[1] == "--upper" {
		upper = true
		start = 2
	}

	if start >= len(os.Args) {
		return
	}

	for i := start; i < len(os.Args); i++ {
		n := Atoi(os.Args[i])
		PrintAlpha(n, upper)
	}

	z01.PrintRune('\n')
}