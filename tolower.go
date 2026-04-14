package main

func ToLower(s string) string {
	result := ""
	for _, ch := range s{
		if ch >= 'A' && ch <= 'Z' {
			ch += 32
		}
		result = result + string(ch)
	}
	return result
}