package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	// 提供一个100毫秒到期的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	fmt.Println("Begin: ", time.Now())
	// #sleep 5
	// 会在命令行sleep 5秒
	// 但是ctx会在100毫秒后到期
	// 所以，注定了err！= nil
	if err := exec.CommandContext(ctx, "sleep", "5").Run(); err != nil {
		// This will fail after 100 milliseconds. The 5 second sleep
		// will be interrupted.
		fmt.Println("End:   ", time.Now())
		fmt.Println(err)
	}
}
