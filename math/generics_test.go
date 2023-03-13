package math

import (
	"testing"

	"github.com/bmizerany/assert"
)

func Test_Abs(t *testing.T) {
	assert.Equal(t, Abs(-1), 1)
	assert.Equal(t, Abs(-1.), 1.)
}
