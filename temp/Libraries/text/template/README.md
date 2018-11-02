# text/template

模板库，使用数据驱动模板生成文本。

// TODO: 完成模板库的介绍

## 概述

```go
package main

import (
    "os"
    "text/template"
)

type Inventory struct {
    Material string
    Count    uint
}

func main() {
    sweaters := Inventory{"wool", 17}
    tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
    if err != nil {
        panic(err)
    }
    err = tmpl.Execute(os.Stdout, sweaters)
    if err != nil {
        panic(err)
    }
}
```

上述代码中，模板文件内容是 "{{.Count}} items are made of {{.Material}}"。

`{{.Count}}` 是一种标记手段。表示，在按照模板生成文本的时候，获取对象的 Count 属性的值，替换此处的标记。

> tmpl.Execute(os.Stdout, sweaters)

这个语句的含义是，使用 sweaters 变量的值，按照模板的设置，生成文本。并把文本输出到 os.Stdout。
[运行代码](https://play.golang.org/p/ZS0HyvHG6-1)后，可看到以下输出。

> 17 items are made of wool

## Actions 动作
