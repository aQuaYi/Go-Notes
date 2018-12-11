package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Cal_Add(t *testing.T) {
	ast := assert.New(t)
	//
	expected := 1
	actual := 1
	ast.Equal(expected, actual)
}
