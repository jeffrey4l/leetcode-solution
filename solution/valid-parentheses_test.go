package solution

import (
	"testing"
)

func TestValidParentheses(t *testing.T) {
	expected := map[string]bool{
		"()":  true,
		"()]": false,
	}

	for k, v := range expected {
		if v != isValid(k) {
			t.Error()
		}
	}

}
