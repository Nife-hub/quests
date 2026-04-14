package main

func BasicJoin(elem []string) string {
	result := ""

	for _, str := range elem {
		result += str
	}
	return result
}