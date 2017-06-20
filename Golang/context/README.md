# Go语言context模块小结

## 简介
在 Go http包的Server中，每一个请求在都有一个对应的 goroutine 去处理，用G表示这个goroutine。

G通常**又**会启动别的 goroutine 用来访问后端服务，比如数据库和RPC服务。

G通常还会需要访问一些与请求特定的数据，比如终端用户的身份认证信息、验证相关的token、请求的截止时间。

当一个请求被取消或超时时，所有用来处理该请求的 goroutine 都应该迅速退出，然后系统才能释放这些 goroutine 占用的资源。

所以，在处理一个请求的时候，要管理这个请求的过程，需要以下功能
1. 跨越多个goroutine的控制工具。
1. 跨越多个goroutine的公共变量。
1. 跨越多个goroutine的统一开关。
你看我写了三遍“跨越多个goroutine”。

golang的context模块就是为了这个目的而生的。

## 源代码

context 包的核心是 struct Context，声明如下：
```go
// 一个 Context 对象可以携带截止时间信号
// 取消信号、和跨多个API的请求作用域数据
// 它的方法是线程安全的。
type Context interface {
    // Done 返回一个 channel，该channel在Context对象被取消时关闭
    // 因此，从Context.Done()中接收到了数据，说明需要关闭所在的goroutine了。
    Done() <-chan struct{}
    
    // Err 方法返回 Done channel 关闭后，Context 对象被取消的原因
    // context被关闭前，Context.Err()的返回值为 nil
    // 错误值为以下两种：
    // context deadline exceeded
    // context canceled
    Err() error
    
    // Deadline 方法返回 Context 被取消的时间点。
    // 如果Context没有设置超时时间，则返回值ok为false    
    Deadline() (deadline time.Time, ok bool)
 
    // Value 方法返回与key对应的value，或 nil
    //**注意**，由于context存在于多个goroutine中，返回值value需要是线程安全的。
    Value(key interface{}) interface{}
}


```
注意: 这里我们对描述进行了简化，更详细的描述查看 context 的 godoc。

Done方法返回一个 channel，这个 channel 对于以 Context方式运行的函数而言，是一个取消信号。

当这个 channel 关闭时，上面提到的这些函数应该终止手头的工作并立即返回。

之后，Err方法会返回一个错误，告知为什么 Context被取消。

一个 Context不能拥有Cancel方法，同时我们也只能 Done channel接收数据。

背后的原因是一致的：接收取消信号的函数 和发送信号的函数通常不是一个。

一个典型的场景是：父操作为子操作启动 goroutine，子操作也就不能取消父操作。

作为一个折中，WithCancel函数 (后面会细说) 提供了一种取消子 Context 对象的方法。

Context对象是线程安全的，你可以把一个 Context对象传递给任意个数的 gorotuine， 对它执行取消 操作时，所有 goroutine 都会接收到取消信号。

Deadline方法允许函数确定它们是否应该开始工作。如果剩下的时间太少，也许这些函数就不值得启动。代码中，我们也可以使用 Deadline对象为 I/O 操作设置截止时间。

Value方法允许 Context对象携带request作用域的数据，该数据必须是线程安全的。
## Context模块的函数

context 包提供了一些函数，协助用户从现有的Context对象创建新的 Context对象。这些 Context对象形成一棵树：当一个Context对象被取消时，继承自它的所有 Context都会被取消。

Background是所有 Context对象树的根，它不能被取消。它的声明如下：
```go
// Background 返回一个空 Context
// 空 Context 不能被取消，也没有
// 截止时间，也不能存储数据。
// Background 通常用在 main, init和
// 和测试代码里，作为入口请求的顶层 Context 
func Background() Context
```

WithCancel和 WithTimeout函数会返回继承的 Context对象，这些对象可以比它们的父 Context更早地取消。

当请求处理函数返回时，与该请求关联的 Context会被取消。当使用多个副本发送请求时，可以使用 WithCancel取消多余的请求。

WithTimeout在设置对后端服务器请求截止时间时非常有用。 下面是这三个函数的声明：
```go
// WithCancel 返回父 Context(parent) 的拷贝
// 一旦parent.Done channel关闭，或
// 该拷贝被cancel，它的Done channel也被关闭
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
 
// CancelFunc 可以取消一个 Context
type CancelFunc func()

// WithTimeout返回父Context(parent)的拷贝
// 一旦parent.Done 被关闭、该拷贝的cancel函数
// 被调用或 timeout 到期，它的Done channel随之关闭。
// 新 Context 如果有 Deadline，必然早于 
// now+timeout，也早于父 Context 的Deadline。
// 此时如果计时器还在运行，cancel函数会释放资源。
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
```

WithValue函数能够将请求作用域的数据与 Context对象建立关系。声明如下：
```go 
// WithValue返回父 Context 的一个拷贝
// 它的 Value 方法返回对应key 的value
func WithValue(parent Context, key interface{}, val interface{}) Context
```

当然，想要知道 Context包是如何工作的，最好的方法是看一个例子。

## 一个例子： Google Web Search

我们的例子是一个 HTTP 服务，它能够将类似于/search?q=golang&timeout=1s的请求 转发给Google Web Search API (见“相关链接”)，然后渲染返回的结果。timeout参数用来告诉 server 取消请求的时间。

这个例子的代码存放在三个包里：
server：它提供 main 函数和 处理 /search 的 http     handler
userip：它能够从请求解析用户的IP，并将请求绑定到一个 Context 对象。
google：它包含了 Search 函数，用来向 Google 发送请求。
深入 server 

server 程序处理类似于 /search?q=golang的请求，返回 Google API 的搜索结果。它将handleSearch函数注册到 /search路由。

处理函数创建一个Context ctx，并对其进行初始化， 以保证Context取消时，处理函数返回。

如果请求的 URL 参数中包含 timeout，那么当 timeout 到期时， Context会被自动取消。 handleSearch 的代码如下：

func handleSearch(w http.ResponseWriter, req *http.Request) {
    // ctx 是该处理函数的 Context 
    // 调用 cancel 函数会关闭 ctx.Done，然后
    // 该函数发起的请求都会收到一个取消信号
    var (
        ctx    context.Context
        cancel context.CancelFunc
    )
    timeout, err := time.ParseDuration(req.FormValue("timeout"))
    if err == nil {
        // 该请求有一个 timeout，所以创建的
        // Context 在timeout 到期时会自动取消。
        ctx, cancel = context.WithTimeout(context.Background(), timeout)
    } else {
        ctx, cancel = context.WithCancel(context.Background())
    }
    defer cancel() // handleSearch一旦返回，ctx 就会被cancel
处理函数(handleSearch) 将query 参数从请求中解析出来，然后通过 userip 包将client IP解析出来。

这里 Client IP 在 后端发送请求时要用到，所以 handleSearch 函数将它 attach 到Context对象 ctx 上。代码如下：
// 检查搜索参数
query := req.FormValue("q")
if query == "" {
    http.Error(w, "no query", http.StatusBadRequest)
    return
}
 
// 将 user IP存储在 ctx 变量中方便其他package使用
userIP, err := userip.FromRequest(req)
if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
}
ctx = userip.NewContext(ctx, userIP)
处理函数带着 Context对象ctx和 query调用google.Search，代码如下：
// 运行 Google Search 并打印结果
start := time.Now()
results, err := google.Search(ctx, query)
elapsed := time.Since(start)
如果搜索成功，处理函数会渲染搜索结果，代码如下：
if err := resultsTemplate.Execute(w, struct {
    Results          google.Results
    Timeout, Elapsed time.Duration
}{
    Results: results,
    Timeout: timeout,
    Elapsed: elapsed,
}); err != nil {
    log.Print(err)
    return
}
深入 userip 包

userip 包提供了两个功能：
从请求解析出 Client IP；
将 Client IP 关联到一个 Context对象。

一个Context对象提供一个key-value 映射，key 和 value 的类型都是 interface{}，但是 key 必须满足等价性（可以比较），value必须是线程安全的。

类似于 userip的包隐藏了映射的细节，提供的是对特定 Context类型值得强类型访问。

为了避免 key 冲突，userip定义了一个非输出类型key，并使用该类型的值作为Context的key。代码如下：
// 为了避免与其他包中的 Context key 冲突
// 这里不输出 key 类型 (首字母小写)
type key int
 
// userIPKey 是 user IP 的 Context key
// 它的值是随意写的。如果这个包中定义了其他
// `Context` key，这些 key 必须不同
const userIPKey key = 0
函数 FromRequest用来从一个 http.Request 对象中解析出 userIP：
func FromRequest(req *http.Request) (net.IP, error) {
    ip, _, err := net.SplitHostPort(req.RemoteAddr)
    if err != nil {
        return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
    }
函数 NewContext返回一个新的Context对象，它携带者 userIP：
func NewContext(ctx context.Context, userIP net.IP) context.Context {
    return context.WithValue(ctx, userIPKey, userIP)
}
函数 FromContext从一个Context对象中解析 userIP：
func FromContext(ctx context.Context) (net.IP, bool) {
    // 如果ctx 中找不到这个 key，ctx.Value 返回 nil
    // 此时net.IP 类型断言返回 ok=false
    userIP, ok := ctx.Value(userIPKey).(net.IP)
    return userIP, ok
}
深入 google 包

函数 google.Search想 Google Web Search API 发送一个 HTTP 请求，并解析返回的 JSON数据。 

该函数接收一个Context对象 ctx 作为第一参数，在请求还没有返回时，一旦ctx.Done关闭，该函数也会立即返回。

Google Web Search API 请求包含 query 关键字和 user IP 两个参数。具体实现如下：
func Search(ctx context.Context, query string) (Results, error) {
    // 准备 Google Search API 请求
    req, err := http.NewRequest("GET", "https://ajax.googleapis.com/ajax/services/search/web?v=1.0", nil)
    if err != nil {
        return nil, err
    }
    q := req.URL.Query()
    q.Set("q", query)
 
    // 如果 ctx 中包含 user IP，则发送给 server
    // Google API 使用 user IP 辨别终端用户从
    // 服务器出实话的请求
    if userIP, ok := userip.FromContext(ctx); ok {
        q.Set("userip", userIP.String())
    }
    req.URL.RawQuery = q.Encode()

函数 Search使用一个辅助函数httpDo发送 HTTP 请求，并在ctx.Done关闭时取消请求 (如果还在处理请求或返回)。 

函数Search传递给 httpDo一个闭包处理 HTTP 结果。下面是具体实现：
var results Results
err = httpDo(ctx, req, func(resp *http.Response, err error) error {
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    
    // 解析 JSON 格式的结果
    // https://developers.google.com/web-search/docs/#fonje
    var data struct {
        ResponseData struct {
            Results []struct {
                TitleNoFormatting string
                URL               string
            }
        }
    }
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        return err
    }
    for _, res := range data.ResponseData.Results {
        results = append(results, Result{Title: res.TitleNoFormatting, URL: res.URL})
    }
    return nil
})
// httpDo 等待闭包函数返回
// 所以可以安全地在这里读取结果
return results, err

函数 httpDo在一个新的 goroutine 中发送 HTTP 请求和处理结果。

如果ctx.Done已经关闭， 而处理请求的 goroutine 还存在，那么取消请求。下面是具体实现：

func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
    // 在一个 goroutine 中发送 HTTP 请求
    // 并将结果交给f 处理
    tr := &http.Transport{}
    client := &http.Client{Transport: tr}
    c := make(chan error, 1)
    go func() { c <- f(client.Do(req)) }()
    select {
    case <-ctx.Done():
        tr.CancelRequest(req)
        <-c // Wait for f to return.
        return ctx.Err()
    case err := <-c:
        return err
    }
}
在自己的代码中使用 Context

许多服务器框架都提供了管理请求作用域数据的包和类型。我们可以定义一个Context接口的实现， 将已有代码和期望Context参数的代码粘合起来。

Gorilla 框架的github.com/gorilla/context 包允许处理函数(handlers) 将数据和请求结合起来，他通过HTTP 请求 到 key-value对 的映射来实现。

在 gorilla.go中，我们提供了一个 Context的具体实现，这个实现的 Value 方法返回的值已经与 gorilla 包中特定的 HTTP 请求关联起来。

还有一些包实现了类似于 Context的取消机制。比如Tomb中有一个 Kill 方法，该方法通过关闭 名为Dying的 channel 发送取消信号。

Tomb也提供了等待 goroutine 退出的方法，类似于 sync.WaitGroup。在 tomb.go 中，我们提供了一个Context的实现，当它的父 Context被取消或 一个Tomb对象被 kill 时，该Context对象也会被取消

## 总结

在 Google，我们要求 Go 程序员把Context作为第一个参数传递给 入口请求和出口请求链路上的每一个函数。

这样机制一方面保证了多个团队开发的 Go 项目能够良好地协作，另一方面它是一种简单的超时和取消机制， 保证了临界区数据 (比如安全凭证) 在不同的 Go 项目中顺利传递。

如果你要在 Context之上构建服务器框架，需要一个自己的Context实现，在框架与期望 Context参数的代码之间建立一座桥梁。

当然，Client 库也需要接收一个Context对象。在请求作用域数据与取消之间建立了通用的接口以后，开发者使用 Context 分享代码、创建可扩展的服务都会非常方便。

## 参考文档
1. [Go语言并发模型：使用 context](https://mp.weixin.qq.com/s?__biz=MzIzODUwMzMzNg%25253D%25253D&idx=1&mid=2247483693&scene=0&sn=6f31dd036f4eef7d39c44eb7034814a7)