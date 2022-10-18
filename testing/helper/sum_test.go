package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	test := assert.New(t)
	test.Equal(Sum(1, 2), 3, "harusnya sih 3")
	test.Equal(Sum(0, 2), 3, "harusnya sih 3")
	test.Equal(Sum(5, 2), 3, "harusnya sih 3")
	test.NotEqualf(Sum(1, 2), 5, "harusnya sih bukan 5")
}
