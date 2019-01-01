package solution

// Probelem: https://leetcode.com/problems/roman-to-integer/

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	Symbol       Value
	I             1
	V             5
	X             10
	L             50
	C             100
	D             500
	M             1000
*/

var m = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

func romanToInt(s string) int {
	sum := 0
	last := 0
	for index, v := range s {

		n := m[string(v)]
		if index > 0 && n > last {
			sum = sum + n - 2*last

		} else {
			sum += n
		}
		last = n
	}
	return sum

}
func TestSuccess(t *testing.T) {
	expectes := map[string]int{
		"III": 3,
		"V":   5,
		"VI":  6,
		"IV":  4,
	}
	for roman, n := range expectes {
		assert.Equal(t, n, romanToInt(roman), "")
	}
}
