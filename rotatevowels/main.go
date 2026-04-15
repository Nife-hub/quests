package main

import (
	"os"

	"github.com/01-edu/z01"
)

// check if rune is a string
func Contains(str string, ch rune) bool {
	for _, c := range str {
		if c == ch {
			return true
		}
	}
	return false
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		z01.PrintRune('\n')
		return
	}

	// combine arguments into a single slice of runes, separating with spaces except after last arg
	var args_s []rune
	for i, arg := range args {
		for _, c := range arg {
			args_s = append(args_s, c)
		}
		if i != len(args)-1 {
			args_s = append(args_s, ' ')
		}
	}

	vowels := "aeiouAEIOU"
	var vowel_pos []int

	// collect positions of vowels
	for i, c := range args_s {
		if Contains(vowels, c) {
			vowel_pos = append(vowel_pos, i)
		}
	}

	// if there are vowels, mirror them
	for i := 0; i < len(vowel_pos)/2; i++ {
		j := len(vowel_pos) - 1 - i
		args_s[vowel_pos[i]], args_s[vowel_pos[j]] = args_s[vowel_pos[j]], args_s[vowel_pos[i]]
	}

	// print result
	for _, c := range args_s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
