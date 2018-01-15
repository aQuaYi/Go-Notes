# testify/assert

github.com/testify/assert弥补了golang没有assert关键字的缺憾。

## 关键点
1. 每一个assert函数，都使用testing.T对象作为第一个参数，以此来输出错误信息。这样可以保证对`go test`的兼容性
1. 每一个assert函数，都会根据断言成功与否来返回一个bool值，这样有利于做进一步的断言。

## 更多要点

请参考[assert的GoDoc](https://godoc.org/github.com/stretchr/testify/assert)