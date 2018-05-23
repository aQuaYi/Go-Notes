// 本程序展示了 StdoutPipe 和 StderrPipe 的标准用法
// 使用Wait观察程序执行结果会导致pipe关闭。
// 所以一定要在 读取结果之后，使用 wait

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	// NOTICE: 使用 Start()
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	var person struct {
		Name string
		Age  int
	}

	if err := json.NewDecoder(stdout).Decode(&person); err != nil {
		log.Fatal(err)
	}

	// NOTICE: 读取完结果后，才使用 Wait()
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s is %d years old\n", person.Name, person.Age)
}
