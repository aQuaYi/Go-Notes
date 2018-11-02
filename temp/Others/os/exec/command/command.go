package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	// cat example | tr a-z A-Z > EXAMPLE
	// 以上命令，可以把example文件的小写字母，全部变成大写的，并保存到 EXAMPLE中
	cmd := exec.Command("tr", "a-z", "A-Z")

	// 提供输入项
	cmd.Stdin = strings.NewReader("some input")
	// 提供输出项
	var out bytes.Buffer
	cmd.Stdout = &out

	// 运行命令
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// 查看命令执行的结果
	fmt.Printf("in all caps: %q\n", out.String())
}
