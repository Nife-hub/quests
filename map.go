package main

func Map (f func(int)bool, a []int) []bool {
	result := make([]bool, len(a))
	for i, c := range a{
		result[i] = f(c)
	}
	return result
}