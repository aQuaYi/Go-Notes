package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("比较两个格式的时间格式字符串")
	earlyStr := time.Now().Format("2006-01-02 15:04:05")
	lateStr := time.Now().Add(time.Hour).Format("2006-01-02 15:04:05")
	inqualitySign := ""
	if earlyStr < lateStr {
		inqualitySign = "<"
	} else {
		inqualitySign = ">"
	}
	fmt.Printf("%s %s %s\n", earlyStr, inqualitySign, lateStr)
	fmt.Println("结论：时间日期的字符串也能比较，判断标准与time.Time一致。")
}
