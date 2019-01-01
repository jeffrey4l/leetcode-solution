package solution

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindromeNumberSuccess(t *testing.T) {
	expected := map[int]bool{
		20:            false,
		11:            true,
		1:             true,
		1221:          true,
		313:           true,
		31:            false,
		1001:          true,
		math.MaxInt64: false,
	}
	timeit(func() {
		for k, v := range expected {

			a := isPalindrome(k)
			assert.Equal(t, a, v, fmt.Sprintf("Expected %d is %v", k, v))
		}
	}, "string", 1000)
	timeit(func() {
		for k, v := range expected {

			a := isPalindrome2(k)
			assert.Equal(t, a, v, fmt.Sprintf("Expected %d is %v", k, v))
		}
	}, "better", 1000)
}
