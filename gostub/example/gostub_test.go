package example

import (
	"testing"

	"github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
)

var num = 100

func Test_StubGlobalVar(t *testing.T) {
	Convey("stub num to -100", t, func() {
		newNum := -100
		stubs := gostub.Stub(&num, newNum)
		defer stubs.Reset()
		So(num, ShouldEqual, newNum)
	})
}
