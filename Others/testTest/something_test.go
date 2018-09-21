// testTest_test.go
package testTest

import (
	"testing"
)

func Test_Division_1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("触发测试没有通过。")
	} else {
		t.Log("Test_Dvivision_1 passed.")
	}
}

func Test_Division_2(t *testing.T) {
	if _, e := Division(1, 0); e == nil {
		t.Error("除数为0时，没有报错。")
	} else {
		t.Log("Test_Division_2 Passed.")
	}
}
