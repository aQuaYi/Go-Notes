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
