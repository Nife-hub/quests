// package main

// import ( 
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Println(os.Args[0])
// }

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[0]
	lastSlash := -1

	// Find last '/'
	for i, r := range arg {
		if r == '/' {
			lastSlash = i
		}
	}

	// Print from lastSlash+1 using rune iteration
	for i, r := range arg {
		if i > lastSlash {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}
