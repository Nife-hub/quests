package main

func AlphaCount(s string) int {
	count := 0

	for _, ch := range s {
		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') {
			count++
		}
	} 
	return count
}


// lowercount := 0
	// for c := 'a'; c <= 'z'; c++ {
	// 	lowercount++
	// }

	// uppercount := 0
	// for v := 'A'; v <= 'Z'; v++ {
	// 	uppercount++
	// }
	// fmt.Print(lowercount + uppercount)