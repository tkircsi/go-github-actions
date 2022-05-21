package math

import "testing"

func TestSum(t *testing.T) {
	c := Sum(10, 20)
	if c != 30 {
		t.Errorf("Sum was incorrect, got: %d, want :%d.", c, 30)
	}
}
