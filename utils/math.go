package utils

// Modulo Returns the mathematical modulo, always non-negative unlike Go's %.
// Usage: Modulo(-1, 10) returns 9, whereas -1 % 10 returns -1.
func Modulo(a, b int) int {
	return ((a % b) + b) % b
}

// Absolute returns the absolute value of an integer.
// Usage: Abs(-5) returns 5, Abs(5) returns 5.
func Absolute(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
