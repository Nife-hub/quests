package main

func ConvertBase(nbr, baseFrom, baseTo string) string {
	n := fromBase(nbr, baseFrom)
	return toBase(n, baseTo)
}

func fromBase(nbr, base string) int {
	baseLen := len(base)
	value := 0
	for i := 0; i < len(nbr); i++ {
		value = value*baseLen + indexOf(base, nbr[i])
	}
	return value
}

func toBase(n int, base string) string {
	if n == 0 {
		return string(base[0])
	}
	baseLen := len(base)
	var result []rune
	for n > 0 {
		r := rune(base[n%baseLen])
		result = append([]rune{r}, result...)
		n = n / baseLen
	}
	return string(result)
}

func indexOf(base string, c byte) int {
	for i := 0; i < len(base); i++ {
		if base[i] == c {
			return i
		}
	}
	return -1
}
