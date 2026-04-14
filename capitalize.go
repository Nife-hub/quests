package main

func Capitalize(s string) string {
	result := ""
	newWord := true

	for i := 0; i < len(s); i++ {
		ch := s[i]

		// Check if alphanumeric
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') {

			if newWord {
				// Capitalize if lowercase
				if ch >= 'a' && ch <= 'z' {
					ch -= 32
				}
				newWord = false
			} else {
				// Lowercase if uppercase
				if ch >= 'A' && ch <= 'Z' {
					ch += 32
				}
			}
		} else {
			// Not alphanumeric → next is new word
			newWord = true
		}

		result += string(ch)
	}

	return result
}