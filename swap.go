package main

func Swap(a *int, b *int) {
	oldA := *a
	oldB := *b
	*a = oldB
	*b = oldA
}