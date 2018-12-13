package calculator

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_Calculator_Add(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	Convey("生成新的 Calculator", t, func() {
		mockMath := NewMockMath(ctrl)
		c := &Calculator{
			math: mockMath,
		}

		Convey("如果， math 运行正常", func() {
			mockMath.
				EXPECT().
				Add(1, 1).
				Return(2)
			Convey("那么， Calculator 显示正确的结果", func() {
				So(c.Add(1, 1), ShouldEqual, 2)
			})
		})

		Convey("如果， math 运行 **不** 正常", func() {
			mockMath.
				EXPECT().
				Add(1, 1).
				Return(3)
			Convey("那么， Calculator 出现错误的结果", func() {
				So(c.Add(1, 1), ShouldNotEqual, 2)
			})
		})
	})

}
