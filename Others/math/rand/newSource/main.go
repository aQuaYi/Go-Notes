package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("rand.Intn生成的随机数，会每次都一样")
	fmt.Println(rand.Intn(100))

	fmt.Println("添加了NewSource的随机数，每次都不一样")
	// NewSource 不一定要是time，每次不一样的内容都可以。
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		fmt.Println(r.Intn(100))
	}
	fmt.Println("请多运行几次，对比结果")
}
