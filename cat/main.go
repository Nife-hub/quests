package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	args := os.Args

	// No arguments → read from stdin
	if len(args) == 1 {
		buf := make([]byte, 1)
		for {
			n, err := os.Stdin.Read(buf)
			if n > 0 {
				z01.PrintRune(rune(buf[0]))
			}
			if err != nil {
				return
			}
		}
	}

	// Arguments provided → read each file
	for i := 1; i < len(args); i++ {
		filename := args[i]

		content, err := os.ReadFile(filename)
		if err != nil {
			printStr("ERROR: ")
			printStr(err.Error()) // must print EXACT error message
			z01.PrintRune('\n')
			os.Exit(1) // must exit with non-zero status
		}

		for _, b := range content {
			z01.PrintRune(rune(b))
		}
	}
}
