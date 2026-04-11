package main

func BasicAtoi(s string) int{
	result := 0
	for _, r := range s {
		digit := int(r - '0')
		result = result*10 + digit 
	}
	return result
}

// '0' = 48