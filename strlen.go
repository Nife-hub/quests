package main

func StrLen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
	// r := []rune(s)
	// return len(r)
}
