package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// Output 会直接 执行 命令，并返回 stdout 的输出结果。
	out, err := exec.Command("./bothOutput").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The message is %s\n", out)
}
