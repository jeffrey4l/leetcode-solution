package longest_common_prefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func longestCommonPrefix(strs []string) string {
	s := ""
	for i := 0; len(strs) > 0 && i < len(strs[0]); i++ {
		c := strs[0][i]
		founded := true
		for index, v := range strs {
			if index == 0 {
				continue
			}
			if len(v) < i+1 || v[i] != c {
				founded = false
				break
			}
		}
		if founded {
			s += string(c)
		} else {
			break
		}
	}
	return s

}

func TestSuccess(t *testing.T) {
	s := longestCommonPrefix([]string{"abc", "ab"})
	assert.Equal(t, s, "ab", "")

	s = longestCommonPrefix([]string{"", "ab"})
	assert.Equal(t, s, "", "")

	s = longestCommonPrefix([]string{})
	assert.Equal(t, s, "", "")

	s = longestCommonPrefix([]string{"abc"})
	assert.Equal(t, s, "abc")
}
