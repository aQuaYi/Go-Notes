package isolation

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNestedScopesWithIsolatedExecution(t *testing.T) {
	output := ""

	Convey("a", t, func() {
		output += "a "

		Convey("aa", func() {
			output += "aa "

			Convey("aaa", func() {
				output += "aaa | "
			})

			Convey("aab", func() {
				output += "aab | "
			})
		})

		Convey("ab", func() {
			output += "ab "

			Convey("aba", func() {
				output += "aba | "
				So(output, ShouldEqual, "a aa aaa | a aa aab | a ab aba | ")
			})
		})

		So([]int{}, ShouldBeEmpty)
		So("asdf", ShouldNotEndWith, "cdf")

		So("asdf", ShouldContainSubstring, "sd") // optional 'expected occurences' arguments?
		So("asdf", ShouldNotContainSubstring, "er")
		So("", ShouldBeBlank)
		So("asdf", ShouldNotBeBlank)

		So(1, ShouldHaveSameTypeAs, 0)
		So(1, ShouldNotHaveSameTypeAs, "asdf")

		So(func() { panic("just panic") }, ShouldPanic)
		So(func() {}, ShouldNotPanic)
		So(func() { panic("just panic") }, ShouldPanicWith, "just panic") // or errors.New("something")
		So(func() { panic("just panic") }, ShouldNotPanicWith, "")        // or errors.New("something")

		So(time.Now(), ShouldHappenBefore, time.Now())
		So(time.Now(), ShouldHappenOnOrBefore, time.Now())
		So(time.Now().Add(time.Nanosecond*1000), ShouldHappenAfter, time.Now())
		So(time.Now().Add(time.Nanosecond*1000), ShouldHappenOnOrAfter, time.Now())
		So(time.Now().Add(time.Nanosecond*100), ShouldHappenBetween, time.Now(), time.Now())
		So(time.Now().Add(time.Nanosecond*100), ShouldHappenOnOrBetween, time.Now(), time.Now())
		So(time.Now(), ShouldNotHappenOnOrBetween, time.Now(), time.Now())
		So(time.Now(), ShouldHappenWithin, time.Nanosecond*100, time.Now())
		So(time.Now(), ShouldNotHappenWithin, time.Nanosecond*10, time.Now())
	})

}
