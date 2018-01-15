package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// 目录下的bothOutput文件会同时输出 standard output 和 standard error
	// 执行 ./bothOutput 1>stdout 2>stderr 可分别查看两种输出的结果

	cmd := exec.Command("./bothOutput")
	// CombinedOutput 会 执行 命令，并合并输出结果
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)
}
