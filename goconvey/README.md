# GoConvey 测试框架

GoConvey 是一个很有特色的 Go 测试框架。

主页： <http://goconvey.co/>

Wiki： <https://github.com/smartystreets/goconvey/wiki/Documentation>

Github： <https://github.com/smartystreets/goconvey>

GoDoc：<https://godoc.org/github.com/smartystreets/goconvey#pkg-subdirectories>

## 快速开始

```go
package package_name

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
    // 只有顶层的 Convey 需要传递 t
    Convey("Given some integer with a starting value", t, func() {
        x := 1
        Convey("When the integer is incremented", func() {
            x++
            Convey("The value should be greater by one", func() {
                So(x, ShouldEqual, 2)
            })
        })
    })
}
```

注意事项：

1. 导入 convey 的时候，重命名为 "."，可以让后续的程序更通顺。
1. 只有顶层的 Convey 函数需要传入 t

## Web 测试报告

![GoConvey-WEB](GoConvey-WEB.webp)

在命令行输入：

```shell
$GOPATH/bin/goconvey
```

会自动在浏览器中打开以上页面。

GoConvey 会对文件夹进行持续监控，并报告多种测试结果。

## 跳过测试

有的时候，某一处改动会带来多个测试断言的失败。为了一个接一个地处理失败的断言。 GoConvey 提供了跳过某个 Convey 或 So 的方法。

> 在 Convey 和 So 前面添加 Skip 即可。

[点击这里，查看 SkipConvey 和 SkipSo 的说明](https://godoc.org/github.com/smartystreets/goconvey/convey#SkipConvey)

SkipConvey 和 SkipSo 的内容会在 WEB 测试报告中，以 "⚠" 符号标记。

## convey 的隔离执行

当 convey 多层嵌套时，convey 的执行流程会与常识不同。建议仔细阅读 [isolated_execution_test.go](https://github.com/smartystreets/goconvey/blob/master/convey/isolated_execution_test.go) 中的测试代码，好好体会一下。

以下抽离了其中的一个片段来讲解。

```go
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
    })

}
```

最后的断言，是不是有点意想不到。我对 convey 的理解是： 每个 convey 执行时，都会再执行一遍，同级以及上级 convey。另外由于闭包，output 会记录所有步骤的修改。

## 定制断言函数

[这里](https://godoc.org/github.com/smartystreets/goconvey/convey#pkg-variables)罗列了 GoConvey 中的原生断言函数，全部以 Should 开头。

实际上，原生断言函数只是 [assertions](https://github.com/smartystreets/assertions) 库中函数的别名。

[assertions](https://github.com/smartystreets/assertions) 中的函数的类型，全部都是

> func(actual interface{}, expected ...interface{}) string

[assertions/filter.go](https://github.com/smartystreets/assertions/blob/master/filter.go) 中定义了断言函数的使用方法。

于是，我们可以自定义断言函数


```go
func ShouldSummerBeComing(actual interface{}, expected ...interface{}) string {
    if actual == "summer" && expected[0] == "coming" {
        return "" // 返回空字符串表示成功。
    }
    return "Summer is not coming."
}
```

完整代码在[这里](summer/summer_test.go)

## TODO: WebUI Code Gen

## TODO: Reset