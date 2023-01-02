package calc

import "testing"

func TestAdd(t *testing.T) {
	var cal Calculator
	res := cal.Add(1, 1)

	if res != 2 {
		t.Fatal("wrong result")
	}
}
