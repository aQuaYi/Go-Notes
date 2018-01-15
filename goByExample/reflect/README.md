# reflect

反射（reflect）可以查看接口的动态类型和动态值，并对其进行操作。
可以先阅读官方文档的中文翻译版[Go语言反射规则](https://github.com/williamhng/The-Laws-of-Reflection)。
## 接口值
反射的操作对象是接口，所以在讲解反射前，需要先说明接口的相关知识。

像int，string，struct等这样的，既规定了具体实现方式，又定义了相关的操作方法的类型，称为*具体类型*。

接口类型是一种*抽象类型*。定义接口类型时，只规定了所需的方法，没有指定具体的实现方式。凡是实现了接口类型所需的方法的具体类型的变量，都能够赋值给这个接口类型的变量，称为其接口值。

接口类型`interface{}`所包含的方法是空的，所以，任何具体类型的值，都能成为接口类型`interface{}`的值。

```go
	var i interface{}
	i = 1
	fmt.Printf("%T\n", i)
	fmt.Printf("%v\n", i)
    // output:
	// int
	// 1
```
在上面的代码中，先定义了i为interface{}接口的变量。然后，把数值1赋值给了i。最后，通过输出函数，查看了interface{}接口变量i的动态类型为int，动态值为1。

通过i的说明可以看到，接口类型的值由两个部分组成：
1. 具体类型，称为接口的动态类型
1. 具体类型的一个值，称为接口的动态值


## reflect.Type 和 reflect.Value


```go
func TypeOf(i interface{}) Type 
func ValueOf(i interface{}) Value 
```
> 注意： 传入上面两个函数的参数，都会被隐式的转换成interface{}接口。