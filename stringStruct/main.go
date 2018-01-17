package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "abcdefg"

	s1 := s[:3]
	s2 := s[1:4]
	s3 := s[2:]

	fmt.Println(s1, s2, s3)

	// 观察s,s1,s2,s3 的内存地址，可以发现，他们都是使用同一个字节数组
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s)))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s1)))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s2)))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s3)))

	// 通常 sbs := string(bs) 会把底层数组进行复制，
	// 但是
	// 利用 unsafe 的方法，使得
	// bs 和 sbs 共有一个底层数组
	// 对 bs 进行修改后，会影响到 sbs
	bs := []byte("abcdefg")
	sbs := toString(bs)
	fmt.Println(sbs)
	bs[0] = 'A'
	fmt.Println(sbs)
}

// toString 返回与 bs 共用底层字节数组的 string
func toString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}
