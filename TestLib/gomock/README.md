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

## 测试Demo

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