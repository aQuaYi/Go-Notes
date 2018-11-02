# gexpect库
现代的Shell对程序提供了最小限度的控制(开始，停止，等等)，而把交互的特性留给了用户。这意味着有些程序，你不能非交互的运行，比如说 passwd。有一些程序可以非交互的运行，但在很大程度上丧失了灵活性，比如说fsck。这表明Unix的工具构造逻辑开始出现问题。Expect恰恰填补了其中的一些裂痕，解决了在Unix环境中长期存在着的一些问题。Expect使用Tcl作为语言核心。不仅如此，不管程序是交互和还是非交互的，Expect都能运用。

gexpect是蠢go语言的expect库，为了更好地控制子进程。

## 基本用法
```go
child, err := gexpect.Spawn("ping -c8 127.0.0.1")
if err != nil {
    panic(err)
}
defer child.Close()
```
生成执行`ping -c8 127.0.0.1`的子进程。

```go
child.Expect("Passwd:")
child.SendLine("123456")
```
child.Expect("Passwd:")会监控命令的输出结果，在发现了`Passwd:`前，`一直阻塞`。如果，需要匹配正则表达式，请使用child.ExpectRegex()

child.SendLine("123456")会发送新的命令，在此则是输入密码"123456"。

child.SendLine() = child.Send() + 按下`Enter`键

```go
child.Interact()
```
把交互功能在命令行还给用户


