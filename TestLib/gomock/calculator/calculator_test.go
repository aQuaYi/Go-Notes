package calculator

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_Calculator_Add(t *testing.T) {

	Convey("生成新的 Calculator", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockMath := NewMockMath(ctrl)
		c := &Calculator{
			math: mockMath,
		}

		Convey("如果， math 运行正常", func() {
			mockMath.EXPECT().
				Add(1, 1).
				Return(2)
			Convey("那么， Calculator 显示正确的结果", func() {
				So(c.Add(1, 1), ShouldEqual, 2)
			})
		})

		Convey("如果， math 运行 **不** 正常", func() {
			mockMath.EXPECT().
				Add(1, 1).
				Do(func(x, y int) {
					Printf("Add(%d,%d)", x, y)
				}).
				Return(3)
			Convey("那么， Calculator 出现错误的结果", func() {
				So(c.Add(1, 1), ShouldNotEqual, 2)
			})
		})

		Convey("如果你愿意，还可以给 Mock 对象处理代码", func() {
			mockMath.EXPECT().
				Add(gomock.Any(), gomock.Any()).
				DoAndReturn(func(x, y int) int {
					Printf("Add(%d,%d)", x, y)
					if x > 65535 || y > 65535 {
						panic("number is too big")
					}
					return x + y
				}).
				AnyTimes()
			Convey("那么， Calculator 显示正确的结果", func() {
				So(c.Add(1, 1), ShouldEqual, 2)
				So(func() { c.Add(1, 65536) }, ShouldPanicWith, "number is too big")
				So(func() { c.Add(65536, 1) }, ShouldPanicWith, "number is too big")
			})
		})

	})

}
