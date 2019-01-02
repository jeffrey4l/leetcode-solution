package solution

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var expected = func() map[int]bool {
	return map[int]bool{
		20:            false,
		11:            true,
		1:             true,
		1221:          true,
		313:           true,
		31:            false,
		1001:          true,
		math.MaxInt64: false,
	}
}()

func TestPalindromeNumberSuccess(t *testing.T) {
	for k, v := range expected {

		a := isPalindrome(k)
		assert.Equal(t, a, v, fmt.Sprintf("Expected %d is %v", k, v))
	}
	for k, v := range expected {

		a := isPalindrome2(k)
		assert.Equal(t, a, v, fmt.Sprintf("Expected %d is %v", k, v))
	}
}

func BenchmarkPalindromeNumber(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for k, v := range expected {

			a := isPalindrome(k)
			if v != a {
				b.Errorf("Expected %d is %v", k, v)
			}
		}
	}
}

func BenchmarkPalindromeNumber2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for k, v := range expected {

			a := isPalindrome2(k)
			if v != a {
				b.Errorf("Expected %d is %v", k, v)
			}
		}
	}
}
