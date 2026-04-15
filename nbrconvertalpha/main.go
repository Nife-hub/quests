package main

func NbrConvertAlpha(n int) string {
	if n <= 0 || n > 26 {
		return ""
	}
	if !(n >= 'a' && n <= 'z') {
		return ""
	}
	return string(rune('a' + n - 1))
}
