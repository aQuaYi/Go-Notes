package main

import (
	"fmt"

	sn "github.com/aQuaYi/Go-Notes/temp/article/sameName"
	// . "github.com/aQuaYi/Go-Notes/temp/article/sameName" // 取消这一行的注释会报错
)

// Name is
var Name = "Name"

func main() {
	fmt.Println(Name)
	fmt.Println(sn.Name)
}
