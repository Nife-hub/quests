package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		fmt.Print(arg)
		fmt.Println()
	}
}