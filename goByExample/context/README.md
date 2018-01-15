# Go语言context模块小结

## 简介
在 Go http包的Server中，每一个请求在都有一个对应的 goroutine 去处理，用G表示这个goroutine。
1. G通常**又**会启动别的 goroutine 用来访问后端服务，比如数据库和RPC服务。
1. G通常还会需要访问一些与请求特定的数据，比如终端用户的身份认证信息、验证相关的token、请求的截止时间。
1. 当一个请求被取消或超时时，所有用来处理该请求的 goroutine 都应该迅速退出，然后系统才能释放这些 goroutine 占用的资源。

所以，在处理一个请求的时候，要管理这个请求的过程，需要以下功能
1. 跨越多个goroutine的控制工具。
1. 跨越多个goroutine的公共变量。
1. 跨越多个goroutine的统一开关。

golang的context模块可以提供以上功能。

## 源代码

### context接口
context 包的核心是 struct Context，声明如下：
```go
// 一个 Context 对象可以携带截止时间信号
// 取消信号、和跨多个API的请求作用域数据
// 它的方法是线程安全的。
type Context interface {
    // Done 返回一个 channel，该channel在Context对象被取消时关闭
    // 因此，从Context.Done()中接收到了数据，说明需要关闭所在的goroutine了。
    Done() <-chan struct{}
    
    // Done channel 关闭后，Err()会返回Context 对象被取消的原因
    // Done channel 关闭前，Context.Err()的返回值为 nil
    // 错误值为以下两种：
    //     context deadline exceeded
    //     context canceled
    Err() error
    
    // Deadline 方法返回 Context 被取消的时间点。
    // 如果Context没有设置超时时间，则返回值ok为false    
    Deadline() (deadline time.Time, ok bool)
 
    // Value 方法返回与key对应的value，或 nil
    //**注意**，由于context存在于多个goroutine中，返回值value需要是线程安全的。
    Value(key interface{}) interface{}
}


```
更详细的描述查看 context 的 godoc。

Context对象是线程安全的，你可以把一个 Context对象传递给任意个数的 gorotuine， 对它执行取消 操作时，所有 goroutine 都会接收到取消信号。

Deadline方法允许函数确定它们是否应该开始工作。如果剩下的时间太少，也许这些函数就不值得启动。代码中，我们也可以使用 Deadline对象为 I/O 操作设置截止时间。

Value方法允许 Context对象携带request作用域的数据，该数据必须是线程安全的。

### Context模块的函数

context 包提供了一些函数，协助用户从现有的Context对象创建新的 Context对象。这些 Context对象形成一棵树：`当一个Context对象被取消时，继承自它的所有 Context都会被取消。`

Background是所有 Context对象树的根，它不能被取消。它的声明如下：
```go
// Background 返回一个空 Context，不能被取消，也没有截止时间，也不能存储数据。
// Background 通常用在 main, init和测试代码里，作为入口请求的顶层 Context 
func Background() Context
```

WithCancel和 WithTimeout函数会返回子Context对象，这些对象可以比它们的父Context更早地取消。

当请求处理函数返回时，与该请求关联的 Context会被取消。当使用多个副本发送请求时，可以使用WithCancel取消多余的请求。

WithTimeout在设置对后端服务器请求截止时间时非常有用。 下面是这三个函数的声明：

```go
// WithCancel 返回父 Context(parent) 的拷贝
// 一旦parent.Done channel关闭，或
// 该拷贝被cancel，它的Done channel也被关闭
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
 
// CancelFunc 可以取消一个 Context
type CancelFunc func()

// WithTimeout返回父Context(parent)的拷贝
// 一旦parent.Done Channel被关闭、该拷贝的cancel函数被调用或 timeout 到期，它的Done channel随之关闭。
// 新 Context 如果有 Deadline，必然早于now+timeout，也早于父 Context 的Deadline。
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
```
WithValue函数能够将请求作用域的数据与 Context对象建立关系。声明如下：
```go 
// WithValue返回父 Context 的一个子context
// 子context.Value(key)
//
// WithValue returns a copy of parent in which the value associated with key is
// val.
//
// Use context Values only for request-scoped data that transits processes and
// APIs, not for passing optional parameters to functions.
//
// The provided key must be comparable and should not be of type
// string or any other built-in type to avoid collisions between
// packages using context. Users of WithValue should define their own
// types for keys. To avoid allocating when assigning to an
// interface{}, context keys often have concrete type
// struct{}. Alternatively, exported context key variables' static
// type should be a pointer or interface.
func WithValue(parent Context, key interface{}, val interface{}) Context
```

## 典型用法

### 通过cancel关闭子goroutine
```go
func fathter() {
    //father获取了context变量ctx，并拥有了ctx的关闭按钮cancel。
	ctx, cancel := context.WithCancel(context.Background())

    //father把ctx变量传递给son
    go son(ctx)

    //一秒钟后，father执行了ctx的cancel函数，关闭了ctx.Done() Channel。
	time.Sleep(time.Second)
	cancel()
}

func son(ctx context.Context) {
	//do something
	for {
    //do something
    select {
    	case <-ctx.Done():
            //do closing things
            //接收到数据，说明ctx.Done()已经关闭。
            //可以着手结束son函数。
    		return
    	default:
	    }
    }
}
```
以上代码说明
1. father goroutine创建ctx变量，并拥有关闭ctx的函数cancel
1. son goroutine获取ctx后，通过能否从`<- ctx.Done()`获取数据来判断是否需要执行关闭操作。

### 通过timeout时间，来关闭子goroutine
```go
func fathter() {
    //father获取了context变量ctx，并拥有了ctx的关闭按钮cancel。
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    //father把ctx变量传递给son
    go son(ctx)

    //Sleep 10秒是为了避免father结束，触发defer cancel()
    time.Sleep(10 * time.Second)
}

func son(ctx context.Context) {
	//do something
	for {
    //do something
    select {
    	case <-ctx.Done():
            //do closing things
            //接收到数据，说明ctx.Done()已经关闭。
            //可以着手结束son函数。
    		return
    	default:
	    }
    }
}
```
以上代码说明
1. father goroutine创建ctx变量，并拥有关闭ctx的函数cancel
1. son goroutine获取ctx后，通过能否从`<- ctx.Done()`获取数据来判断是否需要执行关闭操作。
1. 由于设定好了Timeout，deadline一到，自动关闭ctx.Done() channel。
1. father()中`defer cancel()`的用意是，为了保证ctx一定会被cancel掉，这是一个好习惯。



## 总结
context包通过构建树型关系的Context，来达到上一层Goroutine能对传递给下一层Goroutine的控制。对于处理一个Request请求操作，需要采用context来层层控制Goroutine，以及传递一些变量来共享。

- Context对象的生存周期一般仅为一个请求的处理周期。即针对一个请求创建一个Context变量（它为Context树结构的根）；在请求处理结束后，撤销此ctx变量，释放资源。
- 每次创建一个Goroutine，要么将原有的Context传递给Goroutine，要么创建一个子Context并传递给Goroutine。
- Context能灵活地存储不同类型、不同数目的值，并且使多个Goroutine安全地读写其中的值。
- 当通过父Context对象创建子Context对象时，可同时获得子Context的一个撤销函数，这样父Context对象的创建环境就获得了对子Context将要被传递到的Goroutine的撤销权。
- 在子Context被传递到的goroutine中，应该对该子Context的Done信道（channel）进行监控，一旦该信道被关闭（即上层运行环境撤销了本goroutine的执行），应主动终止对当前请求信息的处理，释放资源并返回。

## 参考文档
1. [Go语言并发模型：使用 context](https://mp.weixin.qq.com/s?__biz=MzIzODUwMzMzNg%25253D%25253D&idx=1&mid=2247483693&scene=0&sn=6f31dd036f4eef7d39c44eb7034814a7)
1. [go程序包源码解读——golang.org/x/net/contex](http://studygolang.com/articles/5131)
1. [golang中context包解读](http://www.01happy.com/golang-context-reading/)
1. [理解Go Context机制](http://lanlingzi.cn/post/technical/2016/0802_go_context/)
