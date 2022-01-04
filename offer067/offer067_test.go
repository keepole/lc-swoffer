// offer067_test.go kee > 2021/11/14

package swoffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrToInt(t *testing.T) {
	assert := assert.New(t)

	var UnitTesting = map[string]int{
		"42":              42,
		"   -42":          -42,
		"4193 with words": 4193,
		"words and 987":   0,
		"-91283472332":    -2147483648,
	}

	for in, out := range UnitTesting {
		assert.Equal(out, StrToInt(in))
	}
}
