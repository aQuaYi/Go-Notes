package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer
	//bytes.Buffer实现了io.Writer接口
	//所以，可以向b中写入数据
	b.WriteString("Hello, ")

	//bytes.Buffer实现了io.Writer接口
	//所以，可以向b中添加数据
	fmt.Fprint(&b, "World!")

	//os.Stdout同样实现了io.Writer接口
	//所以，b中的数据，可以写入os.Stdout
	//os.Stdout就是命令行，于是在命令行中看到了Hello，world！
	b.WriteTo(os.Stdout)
}
