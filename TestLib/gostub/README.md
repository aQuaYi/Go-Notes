# gostub

Go 语言的打桩库。

Github： <https://github.com/prashantv/gostub>

GoDoc： <https://godoc.org/github.com/prashantv/gostub>

## 什么是打桩

桩，或称桩代码，是指用来代替关联代码或者未实现代码的代码。如果函数B用B1来代替，那么，B称为原函数，B1称为桩函数。

打桩就是编写或生成桩代码。

打桩的目的主要有：

- 隔离: 将测试任务之外的，并且与测试任务相关的代码，用桩来代替，从而实现分离测试任务。例如函数 A 调用了函数 B，B又调用了 C 和 D，如果函数B用桩来代替，A 就可以完全割断与 C 和 D 的关系。
- 补齐: 用桩来代替未实现的代码，例如，函数 A 调用了函数 B，而 B 由其他程序员编写，且未实现，那么，可以用桩来代替 B，使 A 能够运行并测试。补齐在并行开发中很常用。
- 控制: 在测试时，人为设定相关代码的行为，使之符合测试需求。

## 使用场景

### 为全局变量打桩

```go
stubs := Stub(&num, 150)
defer stubs.Reset()
```

stubs 会在上级函数结束后，把 num 恢复为原值。

### 为函数打桩

想要给函数打桩，需要将其重构成函数变量

把

```go
func Exec(s string) string {
    // ...
}
```

写成

```go
var Exec = func(s string) string {
    // ...
}
```

> 注意： 这样的重构丝毫不会影响 Exec 的功能。

此时，就可以对 Exec 像普通变量那样进行打桩了

```go
stubs := Stub(&Exec, func(s string) string {
            return "This is stub function."
})
defer stubs.Reset()
```

实际上，gostub 提供了 StubFunc 专门用于函数的打桩。

```go
stubs := StubFunc(&Exec, "This is stub function.")
defer stubs.Reset()
```

### 为 import 的库中的函数打桩

很多时候，我们会 import 其他库，并调用其中的函数。却不能重构那些库中的函数。此时，只要把那些函数定义成本库中函数变量的值，即可。如以下代码所示。

```go
var Marshal = json.Marshal
```

再对 `Marshal` 变量进行打桩即可。

### 为环境变量打桩

<example/gostub_env_test.go> 中的测试函数，显示了给系统变量打桩的方式。

### 组合打桩

如果单个测试中，需要打多个桩，则需要进行组合打桩。

Stub 和 StubFunc 函数，都会生成 stubs 对象， 该对象仍然有 Stub 和 StubFunc 方法， 所以在同一个测试用例中，可以对多个变量或函数打桩，stubs 会将初始值都保存在一起，并在 defer stubs.Reset() 统一做回滚处理。
通常的代码结构是：

```go
stubs := Stub(...)
defer stubs.Reset()
stubs.Stub(...)
...
stubs.StubFunc(...)
...
```

## DEMO

结合 GoConvey 的 Convey 语句的嵌套，即一个函数有一个测试函数，测试函数中嵌套两级 Convey 语句，第一级 Convey 语句对应测试函数，第二级 Convey 语句对应测试用例。在第二级的每个 Convey 函数中都会产生一个 stubs 对象，彼此独立，互不影响。

```go
func TestFuncDemo(t *testing.T) {
    Convey("TestFuncDemo", t, func() {
        Convey("for succ", func() {
            stubs := Stub(&num, 150)
            defer stubs.Reset()
            stubs.StubFunc(&Exec,"xxx-vethName100-yyy", nil)
            var liLei = `{"name":"LiLei", "age":"21"}`
            stubs.StubFunc(&adapter.Marshal, []byte(liLei), nil)
            stubs.StubFunc(&DestroyResource)
            //several So assert
        })

        Convey("for fail when num is too small", func() {
            stubs := Stub(&num, 50)
            defer stubs.Reset()
            //several So assert
        })

        Convey("for fail when Exec error", func() {
            stubs := Stub(&num, 150)
            defer stubs.Reset()
            stubs.StubFunc(&Exec, "", ErrAny)
            //several So assert
        })

        Convey("for fail when Marshal error", func() {
            stubs := Stub(&num, 150)
            defer stubs.Reset()
            stubs.StubFunc(&Exec,"xxx-vethName100-yyy", nil)
            stubs.StubFunc(&adapter.Marshal, nil, ErrAny)
            //several So assert
        })

    })
}
```

## 不适用的复杂情况

1. 被测函数中多次调用了数据库读操作函数接口 ReadDb，并且数据库为key-value型。被测函数先是 ReadDb 了一个父目录的值，然后在 for 循环中读了若干个子目录的值。在多个测试用例中都有将ReadDb打桩为在多次调用中呈现不同行为的需求，即父目录的值不同于子目录的值，并且子目录的值也互不相等。
1. 被测函数中有一个循环，用于一个批量操作，当某一次操作失败，则返回失败，并进行错误处理。假设该操作为Apply，则在异常的测试用例中有将Apply打桩为在多次调用中呈现不同行为的需求，即Apply的前几次调用返回成功但最后一次调用却返回失败。
1. 被测函数中多次调用了同一底层操作函数，比如 exec.Command，函数参数既有命令也有命令参数。被测函数先是创建了一个对象，然后查询对象的状态，在对象状态达不到期望时还要删除对象，其中查询对象是一个重要的操作，一般会进行多次重试。在多个测试用例中都有将 exec.Command 打桩为多次调用中呈现不同行为的需求，即创建对象、查询对象状态和删除对象对返回值的期望都不一样。

## 参考

- [GoStub框架使用指南](https://www.jianshu.com/p/70a93a9ed186)