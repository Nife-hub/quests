package main

func ToUpper(s string) string{
	result := ""
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			ch -= 32
		}
		result = result + string(ch)
	}
	return result
}