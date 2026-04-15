package main

import (
	"fmt"
	"os"
)

func main() {
	rev := os.Args[1:]
	
	for v := len(rev)-1; v >= 0; v--{
		fmt.Print(rev[v])
		fmt.Println()
	}
}