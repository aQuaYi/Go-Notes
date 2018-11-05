# math/big

int 和 float 的范围有限， float 数还因为其 [IEEE754](https://zh.wikipedia.org/zh-cn/IEEE_754) 标准，存在精度问题。
为了实现任意精简的数字运算， Go 提供了 math/big 标准库。代价就是更大的内存和更慢的运行速度。

关于 math/big 库的完整说明，可以参考 <https://golang.org/pkg/math/big/>

## 运算方式

math/big 库中使用方法提供运算，返回值同时也会保存在 receiver 当中。

```golang
d := c.Add(1,2) // c = d = 3
```

```golang
package main

import (
    "fmt"
    "math/big"
)

func main() {
    a := big.NewInt(1)
    b := big.NewInt(2)
    c := big.NewInt(4)
    d := big.NewInt(8)

    d.Add(a, b).Mul(d, c)
    fmt.Println(d)

    d = c.Add(a, b)
    fmt.Println(c, d)
}

/* Output:
12
3 3
*/
```

[在 Go playground 运行](https://play.golang.org/p/KF-DDyWn8a5)
