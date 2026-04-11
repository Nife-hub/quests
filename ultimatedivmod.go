package main

func UltimateDivMod(a *int, b *int){
	oldA := *a
	*a = oldA / *b
	*b = oldA % *b
}