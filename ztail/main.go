package main

import (
	"fmt"
	"os"
)

func strToInt(s string) (int, bool) {
	n := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0, false
		}
		n = n*10 + int(r-'0')
	}
	return n, true
}

func main() {
	if len(os.Args) < 4 || os.Args[1] != "-c" {
		os.Exit(1)
	}
	n, ok := strToInt(os.Args[2])
	if !ok || n <= 0 {
		os.Exit(1)
	}
	files := os.Args[3:]
	status := 0
	printedAny := false
	for _, fn := range files {
		f, err := os.Open(fn)
		if err != nil {
			fmt.Println(err)
			status = 1
			printedAny = true
			continue
		}
		info, err := f.Stat()
		if err != nil {
			fmt.Println(err)
			status = 1
			f.Close()
			printedAny = true
			continue
		}
		if printedAny && len(files) > 1 {
			fmt.Println()
		}
		if len(files) > 1 {
			fmt.Printf("==> %s <==\n", fn)
		}
		size := info.Size()
		start := int64(0)
		if int64(n) < size {
			start = size - int64(n)
		}
		f.Seek(start, 0)
		buf := make([]byte, int(size-start))
		_, err = f.Read(buf)
		f.Close()
		if err != nil {
			fmt.Println(err)
			status = 1
			printedAny = true
			continue
		}
		fmt.Print(string(buf))
		printedAny = true
	}
	if status != 0 {
		os.Exit(1)
	}
}
