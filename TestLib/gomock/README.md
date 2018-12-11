# GoMock

[GoDoc](https://godoc.org/github.com/golang/mock/gomock)

GoMock 是由 Go 官方开发维护的测试框架，实现了较为完整的基于 interface 的 Mock 功能，能够与 Go 内置的 testing 包良好集成，也能用于其它的测试环境中。GoMock 测试框架包含了 GoMock 包和 mockgen 工具两部分，其中 GoMock 包完成对桩对象生命周期的管理，mockgen 工具用来生成 interface 对应的 Mock 类源文件。

## 安装

> 注意： 以下命令需要你的 $GOPATH 应该只指定了一个目录

 一旦安装完成了 Go 语言，就可以可以运行以下命令安装 `gomock` 库和 `mockgen` 工具了：

```shell
go get github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen
```

## 文档

安装完成后，使用 `go doc` 命令查看 `gomock` 的文档。

```shell
go doc github.com/golang/mock/gomock
```

或者查看这个文档的在线版本 <https://godoc.org/github.com/golang/mock/gomock>。

## 运行 mockgen

`mockgen` 可以自动生成代码，其有两种操作：源码与反射。具体的使用方式与区别，可以直接运行 `mockgen` 查看

### mockgen 选项说明

-aux_files string

```string
source 模式专用
<https://github.com/golang/mock/tree/master/mockgen/internal/tests/aux_imports_embedded_interface> 中讨论了需要这个参数的原因和用法
```

-build_flags string

```string
reflect 模式专用
为 `go build` 准备的额外 flag。即，传递给build工具的参数。
```

-debug_parser

```string
TODO: 弄清楚这个
```

-destination string

```string
指定输出文件的位置。未指定的话，输出到命令行。
```

-exec_only string

```string
反射模式专用，如果设置的话，会执行这个反射程序。
```

-imports string

```string
source 模式专用，利用 “逗号分隔” 的 “name=path” 对，来阐明需要引入的内容。
```

-mock_names string

```string
通常，默认 `Interface` 被 Mock 后的名称是 `MockInterface`。但是，可以通过 `-mock_names` 选项，和 “逗号分隔” 的 “interfaceName=mockName” 对的形式来修改默认名称。
```

-package string

```string
指定生成代码的 package name。默认是使用输入代码的 package name，再加上 `mock_` 前缀完成的。
```

-prog_only

```string
reflect 模式专用
只生成反射程序，输出到标准输出，然后退出。
```

-self_package string

```string
The full package import path for the generated code. The purpose of this flag is to prevent import cycles in the generated code by trying to include its own package. This can happen if the mock's package is set to one of its inputs (usually the main one) and the output is stdio so mockgen cannot detect the final output package. Setting this flag will then tell mockgen which import to exclude.
TODO:
```

-source string

```string
source 模式专用
输入 Go 代码文件，打开 source 模式。
```

-write_package_comment bool

```string
设置为 true 时，生成代码时，会带上 `godoc` 用的上的评论。默认就是 true。
```

## 测试 Demo

编写测试用例有一些基本原则，我们一起回顾一下：

- 每个测试用例只关注一个问题，不要写大而全的测试用例
- 测试用例是黑盒的
- 测试用例之间彼此独立，每个用例要保证自己的前置和后置完备
- 测试用例要对产品代码非入侵

根据基本原则，我们不要在一个测试函数的多个测试用例之间共享mock控制器，于是就有了下面的Demo:

```go
func TestObjDemo(t *testing.T) {
    Convey("test obj demo", t, func() {
        Convey("create obj", func() {
            ctrl := NewController(t)
            defer ctrl.Finish()
            mockRepo := mock_db.NewMockRepository(ctrl)
            mockRepo.EXPECT().Retrieve(Any()).Return(nil, ErrAny)
            mockRepo.EXPECT().Create(Any(), Any()).Return(nil)
            mockRepo.EXPECT().Retrieve(Any()).Return(objBytes, nil)
            stubs := StubFunc(&redisrepo.GetInstance, mockRepo)
            defer stubs.Reset()
            ...
        })

        Convey("bulk create objs", func() {
            ctrl := NewController(t)
            defer ctrl.Finish()
            mockRepo := mock_db.NewMockRepository(ctrl)
            mockRepo.EXPECT().Create(Any(), Any()).Return(nil).Times(5)
            stubs := StubFunc(&redisrepo.GetInstance, mockRepo)
            defer stubs.Reset()
            ...
        })

        Convey("bulk retrieve objs", func() {
            ctrl := NewController(t)
            defer ctrl.Finish()
            mockRepo := mock_db.NewMockRepository(ctrl)
            objBytes1 := ...
            objBytes2 := ...
            objBytes3 := ...
            objBytes4 := ...
            objBytes5 := ...
            mockRepo.EXPECT().Retrieve(Any()).Return(objBytes1, nil)
            mockRepo.EXPECT().Retrieve(Any()).Return(objBytes2, nil)
            mockRepo.EXPECT().Retrieve(Any()).Return(objBytes3, nil)
            mockRepo.EXPECT().Retrieve(Any()).Return(objBytes4, nil)
            mockRepo.EXPECT().Retrieve(Any()).Return(objBytes5, nil)
            stubs := StubFunc(&redisrepo.GetInstance, mockRepo)
            defer stubs.Reset()
            ...
        })
        ...
    })
}
```

## 参考

- [GoMock框架使用指南](https://www.jianshu.com/p/f4e773a1b11f)