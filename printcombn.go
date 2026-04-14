package main

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n < 1 || n > 9 {
		return
	}

	combination := make([]int, n)
	for i := 0; i < n; i++ {
		combination[i] = i
	}

	for {
		printCombination(combination)

		if !nextCombination(combination) {
			break
		}

		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
}

func printCombination(combination []int) {
	for _, digit := range combination {
		z01.PrintRune(rune(digit + '0'))
	}
}

func nextCombination(combination []int) bool {
	n := len(combination)

	for i := n - 1; i >= 0; i-- {
		if combination[i] < 10-n+i {
			combination[i]++

			for j := i + 1; j < n; j++ {
				combination[j] = combination[j-1] + 1
			}

			return true
		}
	}
	z01.PrintRune('\n')
	return false
}
