package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("./bothOutput")

	// 获取 cmd 的 standard err 输出接口
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	// 运行 cmd 分别向两个输出端口输出
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// 读取 stderr 的输出内容
	slurp, _ := ioutil.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)

	// Wait 会关闭 stderr
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
