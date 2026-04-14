package main

func TrimAtoi(s string) int {
	sign := 1           // assume positive unless we find a '-'
	num := 0            // will store the final number
	foundDigit := false // tracks if we've found any digits yet

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '-' && !foundDigit {
			sign = -1 // set sign to negative
		} else if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0') // build the number
			foundDigit = true         // mark that we've found a digit
		}
		// ignore all other characters
	}
	return num * sign
}
