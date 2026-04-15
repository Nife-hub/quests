package main

func IsPrime(n int) bool{
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	if n == 3 {
		return true
	}

	for i := 3; i <= n; i +=3 {     // loop through multiples of 3 starting from 3 bcos any number that is a multiple of 3 and greater than 3 cannot be prime
		if n%i == 0 {
			return false
		}
	}
	return true
}