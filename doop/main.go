package main

import (
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		return
	}

	aStr := args[0]
	op := args[1]
	bStr := args[2]

	// Validate operator
	if op != "+" && op != "-" && op != "*" && op != "/" && op != "%" {
		return
	}

	// Convert a and b to ints
	a, ok1 := Atoi(aStr)
	b, ok2 := Atoi(bStr)
	if !ok1 || !ok2 {
		return
	}

	// Division or modulo by zero
	if (op == "/" || op == "%") && b == 0 {
		if op == "/" {
			Print("No division by 0\n")
		} else {
			Print("No modulo by 0\n")
		}
		return
	}

	// Compute result with overflow detection
	result, ok := Compute(a, b, op)
	if !ok {
		return
	}

	Print(IntToString(result) + "\n")
}

func Compute(a, b int, op string) (int, bool) {
	switch op {
	case "+":
		res := a + b
		if (a > 0 && b > 0 && res < 0) || (a < 0 && b < 0 && res > 0) {
			return 0, false
		}
		return res, true

	case "-":
		res := a - b
		if (a > 0 && b < 0 && res < 0) || (a < 0 && b > 0 && res > 0) {
			return 0, false
		}
		return res, true

	case "*":
		res := a * b
		if a != 0 && res/a != b {
			return 0, false
		}
		return res, true

	case "/":
		return a / b, true

	case "%":
		return a % b, true
	}
	return 0, false
}

func Atoi(s string) (int, bool) {
	sign := 1
	i := 0
	n := 0

	if len(s) == 0 {
		return 0, false
	}

	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		n = n*10 + int(s[i]-'0')

		// overflow check (using int32 compatible boundaries)
		if n < 0 {
			return 0, false
		}
	}

	return n * sign, true
}

func IntToString(n int) string {
	if n == 0 {
		return "0"
	}

	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	digits := []byte{}
	for n > 0 {
		d := byte(n % 10)
		digits = append([]byte{d + '0'}, digits...)
		n /= 10
	}

	return sign + string(digits)
}

func Print(s string) {
	os.Stdout.Write([]byte(s))
}
