# GoMock

[GoDoc](https://godoc.org/github.com/golang/mock/gomock)

GoMock 是由 Go 官方开发维护的测试框架，实现了较为完整的基于 interface 的 Mock 功能，能够与 Go 内置的 testing 包良好集成，也能用于其它的测试环境中。GoMock 测试框架包含了 GoMock 包和 mockgen 工具两部分，其中 GoMock 包完成对桩对象生命周期的管理，mockgen 工具用来生成 interface 对应的 Mock 类源文件。

## 安装

> 注意： 以下命令需要你的 $GOPATH 应该只指定了一个目录

在命令行运行以下命令安装

```shell
$ go get github.com/golang/mock/gomock
$ cd $GOPATH/src/github.com/golang/mock/mockgen
$ go build
$ mv mockgen $GOPATH/bin
$ cd ~
$ mockgen
mockgen has two modes of operation: source and reflect.
...
```

## 参考

- [GoStub框架使用指南](https://www.jianshu.com/p/70a93a9ed186)