package main

func IsPrime2(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	// Only check up to sqrt(n)
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}

	// Make sure we start with an odd number
	if nb%2 == 0 {
		nb++
	}


	for !IsPrime2(nb) {
		nb += 2 // skip even numbers
	}
	return nb
}

