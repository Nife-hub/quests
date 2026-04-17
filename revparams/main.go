// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	rev := os.Args[1:]
	
// 	for v := len(rev)-1; v >= 0; v--{
// 		fmt.Print(rev[v])
// 		fmt.Println()
// 	}
// }

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	rev := os.Args[1:]

	for v := len(rev) - 1; v >= 0; v-- {
		for _, r := range rev[v] {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
