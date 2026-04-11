package main

func Atoi(s string) int {
	signCount := 0
	negative := false
	
	for i := 0; i < len(s); i++ {
		if s[i] == '+' || s[i] == '-' {
			signCount++

			if i != 0 {
				return 0
			}

			if s[i] == '-' {
				negative = true
			}

			if signCount > 1 {
				return 0
			}
		}
	}

	result := 0
	for _, r := range s {

		if r == '+' || r == '-' {
			continue
		}

		if r < '0' || r > '9' {
			return 0
		}

		result = result*10 + int(r-'0')
	}

	if negative {
		result = -result
	}

	return result
}