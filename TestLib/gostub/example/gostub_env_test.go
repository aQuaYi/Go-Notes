package example

import (
	"os"
	"testing"

	"github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStubSettedEnv(t *testing.T) {
	stubs := gostub.New()

	Convey("设置环境变量 SSE", t, func() {
		os.Setenv("SSE", "SSE")

		Convey("如果，没有对其打桩的话", func() {

			Convey("其中的变量值，应该是原先的值", func() {
				So(os.Getenv("SSE"), ShouldEqual, "SSE")
			})
		})

		Convey("如果，对其打桩的话", func() {
			stubs.SetEnv("SSE", "SSE Stubed")
			defer stubs.Reset()

			Convey("其中的变量值，应该是 Stub 的值", func() {
				So(os.Getenv("SSE"), ShouldEqual, "SSE Stubed")
			})
		})

		Convey("如果，打桩成 unset 的话", func() {
			stubs.UnsetEnv("SSE")
			defer stubs.Reset()

			Convey("应该查询不到这个变量的值", func() {
				_, ok := os.LookupEnv("SSE")
				So(ok, ShouldBeFalse)
			})
		})
	})

}

func TestStubUnsettedEnv(t *testing.T) {
	stubs := gostub.New()

	Convey("unset 环境变量 USE", t, func() {
		os.Unsetenv("USE")

		Convey("如果，没有对其打桩的话", func() {

			Convey("应该查询不到这个变量的值", func() {
				_, ok := os.LookupEnv("USE")
				So(ok, ShouldBeFalse)
			})
		})

		Convey("如果，对其打桩的话", func() {
			stubs.SetEnv("USE", "USE Stubed")

			Convey("其中的变量值，应该是 Stub 的值", func() {
				So(os.Getenv("USE"), ShouldEqual, "USE Stubed")
			})
			stubs.Reset()

			Convey("Stubs.Reset() 后，应该依然查询不到该值", func() {
				_, ok := os.LookupEnv("USE")
				So(ok, ShouldBeFalse)
			})
		})
	})
}
