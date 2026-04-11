package main

func BasicAtoi2(s string) int{
	result := 0
	for _, r := range s {
		digit := int(r - '0')
		result = result*10 + digit 
		
		if r < '0' || r > '9' {
			return 0
		}
	}
	return result
}

	// if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r == ' ' {
	// 	return 0
	// }