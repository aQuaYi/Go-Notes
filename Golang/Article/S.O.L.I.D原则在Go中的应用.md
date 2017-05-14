# S.O.L.I.D 原则在 Go 中的应用

原文地址：[【译】S.O.L.I.D 原则在 Go 中的应用（上）](http://yemengying.com/2016/09/11/solid-go-design-1/index.html)，[S.O.L.I.D 原则在 Go 中的应用（下）](http://kevin.doyeden.com/2016/09/21/solid-go-design-2/)


## 臭代码的表现

- Rigid 代码僵硬，由于严格的类型和参数导致修改代码的成本提高。
- Fragile 代码脆弱，一点小的改动就会造成巨大的破坏。
- Immobile 难以重构。
- Complex 过度设计。
- Verbose 条例不清。

## SOLID

2002年，Robert Martin 描述了可重用软件设计的SOLID 原则

- 单一责任：Single Responsibility Principle
- 开放封闭：Open / Closed Principle
- 里氏替换：Liskov Substitution Principle
- 接口分离：Interface Segregation Principle
- 依赖倒置：Dependency Inversion Principle


## 单一责任
A class should have one, and only one, reason to change(修改某个类的原因，有且只有一个)，即，一个类只负责一项职责。

`耦合`是指两个东西需要一起修改—对其中一个的改动会影响到另一个。
`内聚`用来描述一段代码内各个元素彼此结合的紧密程度。

## 开放封闭
Bertrand Mey认为，Software entities should be open for extension,but closed for modification（软件实体应当对扩展开放，对修改关闭）。

```go
package main
import (
	"fmt"
)

type A struct {
	year int
}

func (a A) Greet() {
	fmt.Println("Hello GolangUK", a.year)
}

type B struct {
	A
}

func (b B) Greet() {
	fmt.Println("Welcome to GolangUK", b.year)
}

func main(){
	var a A
	a.year = 2016
	var b B
	b.year = 2016
	a.Greet()//Hello GolangUK 2016
	b.Greet()//Welcome to GolangUK 2016
}
```
上面的代码中，我们有类型A，包含属性 year 和一个方法 Greet。我们还有类型B，B中嵌入(embedding)了类型A，并且B提供了他自己的 Greet 方法，覆盖了A的。

嵌入不仅仅是针对方法，还可以通过嵌入使用被嵌入类型的属性。我们可以看到，在上面的例子中，因为A和B定义在同一个包中，所以B可以像使用自己定义的属性一样使用A中的 private 的属性 year。


```go
package main
import (
	"fmt"
)

type Cat struct{
	Name string
}

func (c Cat) Legs() int {
	return 4
}

func (c Cat) PrintLegs() {
	fmt.Printf("I have %d legs\n", c.Legs())
}

type OctoCat struct {
	Cat
}

func (c OctoCat) Legs() int {
	return 5
}

func main() {
	var octo OctoCat
	fmt.Printf("I have %d legs\n", octo.Legs())// I have 5 legs
	octo.PrintLegs()// I have 4 legs
}
```
在这个例子中，我们有一个 Cat 类型，它拥有一个 Legs 方法可以获得腿的数目。我们将 Cat 类型嵌入到一个新类型 OctoCat 中，然后声明 Octocat 有5条腿。然而，尽管 OctoCat 定义了它自己的 Legs 方法返回5，在调用 PrintLegs 方法时依旧会打印“I have 4 legs”。

这是因为 PrintLegs 方法是定义在 Cat 类型中的，它将 Cat 作为接收者，所以会调用 Cat 类型的 Legs 方法。Cat 类型并不会感知到它被嵌入到其他类型中，所以它的方法也不会被更改。

所以，我们可以说 Go 的类型是对扩展开放，对修改关闭的。

## 里氏替换
里氏替换原则由 Barbara Liskov 提出，如果两个类型表现的行为对于调用者来说没有差别，那么我们就可以认为这两个类型是可互相替换的。

在面向对象的语言中，里氏替换原则通常解释为一个抽象基类拥有继承它的许多具体的子类。但 Go 中并没有类或继承，所以无法通过类的继承来实现替换。

### 接口

但是，我们可以通过接口实现替换。在 Go 中，类型并不需要指定他们实现的接口，只要在类型中提供接口要求的所有方法即可。

所以，Go 中的接口是隐式的(非侵入性的)，而非显式的。这对于我们如何使用这门语言有着深远的影响。

一个好的接口应该是小巧的，比较流行的做法是一个接口只包含一个方法。因为一般情况下，小的接口往往意味着简洁的实现。

io.Reader
```go
type Reader interface {
        // Read reads up to len(buf) bytes into buf.
        Read(buf []byte) (n int, err error)
}
```
下面让我们看下 Go 中我最爱的接口—io.Reader。

io.Reader 接口的功能非常简单；将数据读取到提供的缓冲区中，并返回读取的字节数以及读取过程中遇到的错误。虽然看上去简单，但是确非常有用。

因为io.Reader可以处理任何可以表示为字节流的东东，我们几乎可以为所有东西构造读取器；比如：一个常量字符串，字节数组，标准输入，网络流，tar 文件，通过 ssh 远程执行命令的标准输出，等等。

而由于实现了同样的接口，这些具体实现都是互相可替换的。

我们可以用 Jim Weirich 的一句话来描述里式替换原则在 Go 中的应用。

> Require no more, promise no less。

## 接口分离
Robert C. Martin 解释为：

> 调用者不应该被强制依赖那些他们不需要使用的方法。

在 Go 中，接口隔离原则的应用可以参考如何分离一个函数功能的过程。举个例子，我们需将一份文档持久化到磁盘。函数签名可以设计如下：

```go
// Save 方法将文档的内容写到文件f
func Save(f *os.File, doc *Document) error
```
我们定义的 Save 方法，将 *os.File 作为文档写入的目标，这样的设计会有一些问题。

Save 方法的签名设计排除了将文档内容存储到网络设备上的可能。假设后续有将文档存储到网络存储设备上的需求，Save 方法的签名需要做出相应的改变，导致所有 Save 方法的调用方也需要做出改变。

由于 Save 方法直接在磁盘上操作文件，导致对测试不友好。为了验证 Save 所做的操作，测试需要在文档写入后从磁盘上读取文档内容来做验证。 除此之外，测试还要确保文件被写入到临时空间，之后被删除。

*os.File 定义了许多与 Save 操作不相关的方法，比如读取文件目录，检查一个路径是否是符号链接。如果 Save 方法的签名只描述 *os.File 部分相关的操作会更有帮助。

我们应该如何解决这些问题呢？

```go
// Save 方法将文档文档内容写入到指定的ReadWriterCloser
func Save(rwc io.ReadWriteCloser, doc *Document) error
```
使用 io.ReadWriteCloser 我们可以根据接口隔离原则来重新定义 Save 方法，将更通用的文件描述接口作为参数。

重构后，任何实现了 io.ReadWriteCloser 接口的类型都可以替代之前的 *os.File 接口。这扩大了 Save 方法的应用场景，相比使用 *os.File 接口，Save 方法对调用者开说变得更加透明。

Save 方法的编写者也无需关心 *os.File 包含的那些不相关的方法，因为这些细节都被 io.ReadWriteCloser 接口屏蔽掉了。我们还可以进一步将接口隔离原则发挥一下。

首先，如果 Save 方法遵循单一职责原则，方法不应该读取文件内容来对刚写入的内容做验证，这应该是另一个代码片段应该做的事。因此，我们进一步细化传递给 Save 方法的接口定义，仅保留写入和关闭的功能。

```go
// Save 方法将文档文档内容写入到指定的WriterCloser
func Save(wc io.WriteCloser, doc *Document) error
```
译者注：注意，这里接口名字是io.WriteCloser，而上一个签名的参数是io.ReadWriterCloser

其次，根据我们所期望的通用文件描述所具备的功能，给 Save 方法提供关闭流的机制。但是这会引发一个新的问题： wc 在什么时机关闭。Save 方法可以无条件的调用 Close 方法，或者是 Close 方法在执行成功的条件下才会被调用。

不管哪种关闭流的方式都会产生个问题，因为 Save 方法的调用者可能希望在写入的文档的流后面追加数据，而此时流已经被关闭。

```go
type NopCloser struct {
    io.Writer
}

// Close 重写 Close 方法，提供空实现
func (c *NopCloser) Close() error { return nil }
```
如上示例代码所示，一种粗暴的做法就是重新定一个类型，组合了 io.Writer , 重写 Close 函数，替换为空实现，防止 Save 方法关闭数据流。

但是，这违反了里氏替换原则，因为 NopCloser 并没有真正关闭流。

```go
// Save 方法将文档文档内容写入到指定的Writer
func Save(w io.Writer, doc *Document) error
```
一种更加优雅的解决方案是重新定义 Save 方法的参数，将 io.Writer 作为参数，把 Save 方法的职责进一步细化，除了写入数据，其他不相关的事情都不做。

通过将接口隔离原则应用到 Save 方法，把方法功能更加明确化，它仅需要一种可以写入的东西就可以。方法的定义更具有普适性，现在我们可以使用Save 方法去保存数据到任何实现了 io.Writer 的设备。

> Go 中非常重要的一个原则就是接受interface，返回structs。 – Jack Lindamood

上述引言是 GO 在这些年的发展过程中渗透到 GO 设计思想里中的非常有意思的约定。

Jack 的精悍言论可能跟实际会有细微差别，但是我认为这是 Go 设计中颇具有代表性的声明。

## 依赖倒置

最后一个原则是依赖倒置。可以这样理解：上层模块不应该依赖于下层模块，他们都应该依赖于抽象。

> 抽象不应依赖于实现细节，实现细节应该依赖于抽象。 – Robert C. Martin

那么，对于 Go 语言开发者来讲，依赖倒置具体指的是什么呢？

如果你应用了我们上面讲述的4个原则，那么你的代码已经组织在独立的 package 中，每一个包的职责都定义清晰。你的代码的依赖声明应该通过接口来实现，这些接口仅描述了方法需要的功能行为，换句话说，你不需要为此做太多的改变。

因此我认为,在 Go 语言中，Martin 所讲的依赖倒置是你的依赖导入的结构。

在 Go 语言中，你的依赖导入结构必须是非循环的，不遵守此约定的后果是可能会导致编译错误，但是更为严重的是设计上的错误。

良好的依赖导入结构因该是平坦的，而不是层次结构很深。如果你有一个包，在没有其他包的情况下，不能正常工作，这可能你的代码包的组织结构没有划分好边界。

依赖倒置原则鼓励你将具体实现细节尽可能的放到依赖导入结构的最顶层，在 main package 或者更高层级的处理程序中，让低层级的代码去处理抽象的接口。

## SOLID Go 设计

简要回顾一下，在将 SOLID 应用到 Go 语言时，每一个原则都陈述了其设计思想，但是所有原则都遵循了同一个中心思想。

- 单一职责原则鼓励你组织 function，type 到高内聚的包中，type 或者方法都应该有单一的职责。
- 开闭原则鼓励你通过组合简单类型来表达复杂类型。
- 里氏替换原则鼓励你通过接口来表达包之间的依赖关系，而非具体的实现。通过定义职责明确的接口，我们可以确保其具体实现足以满足接口契约。
- 接口隔离原则将里氏替换原则进一步提升，鼓励你在定义方法或者函数的时候仅包含他所需要的功能。如果仅需要一个 interface 类型的参数的方法就可以满足业务功能，那么我们可以认为这也满足了单一职责原则。
- 依赖倒置原则鼓励将你的 package 的依赖从编译阶段推迟到运行时，这样可以减少 import 的数量。

一句话来总结以上讲述: 善用 interface 可以将 SOLID 原则应用到 Go 编程中。因为interface 让你将关注点放在描述包的接口上，而非具体的实现，这也是实现解耦的另一种方式，实际上这也是我们设计的目标，松耦合的软件更容易对修改开放。