# Go 语言笔记

GoMock 是由 Go 官方开发维护的测试框架，实现了较为完整的基于 interface 的 Mock 功能，能够与 Go 内置的 testing 包良好集成，也能用于其它的测试环境中。GoMock 测试框架包含了 GoMock 包和 mockgen 工具两部分，其中 GoMock 包完成对桩对象生命周期的管理，mockgen 工具用来生成 interface 对应的 Mock 类源文件。

## 参考

- [GoStub框架使用指南](https://www.jianshu.com/p/70a93a9ed186)