package summer

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func ShouldSummerBeComming(actual interface{}, expected ...interface{}) string {
	if actual == "summer" && expected[0] == "coming" {
		return "" // 返回空字符串表示成功。
	}
	return "Summer is not coming."
}

func Test_summer(t *testing.T) {
	Convey("Test_summer", t, func() {
		So("summer", ShouldSummerBeComming, "coming")
		So("winter", ShouldSummerBeComming, "coming")
	})
}
