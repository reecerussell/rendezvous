# Factorials

<img src="https://media.giphy.com/media/3o7btPCcdNniyf0ArS/giphy.gif" align="right" width="300" />

Given a number, return true if the input is a factorial of any natural number.

### Example

```
> isFactorial(3)
> false
> isFactorial(6)
> true
```

### Solution

```go
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
```
