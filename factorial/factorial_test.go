package factorial

import "testing"

var factorials = map[int]bool{
	1: true,
	2: true,
	3: false,
	4: false,
	5: false,
}

func TestIsFactorial(t *testing.T) {
	for v, f := range factorials {
		if ok := IsFactorial(v); ok != f {
			t.Fatalf("the value '%d' is not a factorial of a number", v)
		}
	}
}
