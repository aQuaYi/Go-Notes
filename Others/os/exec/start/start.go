package main

import (
	"log"
	"os/exec"
)

func main() {
	// sleep 5秒
	cmd := exec.Command("sleep", "5")

	// Start 会在cmd启动后，立刻返回
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")

	// Wait 会阻塞地等待程序执行的结果。
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}
