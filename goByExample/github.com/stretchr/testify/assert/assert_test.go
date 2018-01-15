package assert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func Test_WithT(t *testing.T) {
// 	t.Log("每一项测试都是失败的，好观察出错时的输出")

// 	actual := 1234
// 	assert.Equal(t, 123, actual, "actual should be 123, not %d", actual)

// 	assert.NotEqual(t, 1234, actual, "actual shoudl not be 1234")

// 	var N []int
// 	N = []int{1, 2, 3}
// 	assert.Nil(t, N, "N Should be nil")

// 	if assert.NotNil(t, N) {
// 		t.Log("马上判断切片是否相等")
// 		assert.Equal(t, []int{1, 2, 3, 4}, N, "assert果然可以判断切片是否相等")
// 	}
// }

func Test_ADD(t *testing.T) {
	//NOTICE: 为了少写一个参数，可以先生成一个*assert.Assertion对象
	a := assert.New(t)
	a.Equal(2, Add(1, 1), "Add(1, 1)")

	a.NotEqual(3, Add(1, 1), "1+1 居然等于3")
}

func Test_Sub(t *testing.T) {
	a := assert.New(t)
	a.Equal(1, Sub(2, 1), "2-1!=1")
}
