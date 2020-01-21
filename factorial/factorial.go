package factorial

// IsFactorial will return a true if num is a factorial
// of a natural number, otherwise false will be returned.
func IsFactorial(num int) bool {
	f := 1

	for true {
		// Check if num is a factorial.
		if f == num {
			return true
		}

		// If f is greater than num, return false
		// as it is not and won't be a factorial.
		if f > num {
			return false
		}

		// Find next factorial value, using
		// f(n+1) = f(n) * n+1, where f(0) = 1
		f = f * (f + 1)
	}

	return false
}
