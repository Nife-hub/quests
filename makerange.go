package main

func MakeRange(min, max int) []int {
	if min >= max {
		return []int{}
	}

	size := max - min
	result := make([]int, size)
	for i := 0; i < size; i++{
		result[i] = min + i
	}
	 return result
}
