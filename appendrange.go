package main

func AppendRange(min, max int) []int {
	if min >= max {
		return []int{}
	}
	var result []int

	for i := min; i < max; i++{
		result = append(result, i)
	}
	return result
}
