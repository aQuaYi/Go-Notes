// panicDeferRecover.go
/*
Function可以看做是一层套一层使用的的。
main(){
	//do something...in main
	A(){
		//do something...in A
		defer func(){//A}()
		B(){
			//do something...in B
			defer func(){//B}()
			C(){
				//do something...in C
				defer func(){//C}()
				panic()
				//do OTHER something...in C
			}
			//do OTHER something...in B
		}
		//do OTHER something...in A
	}
	//do OTHER something...in main
}
在上面的代码中， 肯定能被执行的代码有
* //do something...in main
* //do something...in A
* //do something...in B
* //do something...in C

肯定不能被执行的代码是
* //do OTHER something...in C

想要执行`//do OTHER something...in B`中的代码话，需要`defer func(){//C}()`中使用了recover（）来处理panic。

如果所有的defer中，都没有recover()的话，所有的`//do OTHER something...`都不会被执行。整个程序在panic()后，会按层级由里到外执行defer的内容，然后非正常结束。


如果上面的main（）中，只有使用了一次recover（），而没有执行
* //do OTHER something...in B
* //do OTHER something...in A
却执行了
* //do OTHER something...in main
recover（）该出现谁的defer中呢？
答案是
* defer func(){//A}()

*/

package main

import (
	"fmt"
	"time"
)

func panicFunc(i int) {
	fmt.Println("\tpanicFunc: Begin...")
	time.Sleep(time.Second)

	//可以注释掉defer块，观察输出的异同
	defer func() {
		fmt.Println("\tpanicFunc: defer: Before recover")
		time.Sleep(time.Second)
		if err := recover(); err != nil {
			fmt.Printf("\tpanicFunc: err's type is %T\n", err)
			fmt.Println("\tpanicFunc: Recover: Catch Panic, it is", err)
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second)
	fmt.Println("\tpanicFunc: Panic, Afer one second")
	time.Sleep(time.Second)
	panic(i) //可以注释掉这一行，观察输出的异同

	//panic以后的语句都不会被执行
	time.Sleep(time.Second)
	fmt.Println("\t\tpanicFunc: After panic")
	time.Sleep(time.Second)
	defer func() {
		fmt.Println("\t\tpanicFunc: defer 2, After panic")
	}()
}

func main() {
	fmt.Println("main: Begin...")
	time.Sleep(time.Second)
	defer func() {
		fmt.Println("main: defer 1")
		time.Sleep(time.Second)
	}()
	defer func() {
		fmt.Println("main: defer 2")
		time.Sleep(time.Second)
		fmt.Println("main: defer 2: Before recover")
		time.Sleep(time.Second)

		//同时注释掉这个recover块， 观察输出的不同
		if err := recover(); err != nil {
			fmt.Println("main: defer 2: Catch panic, it is", err)
			time.Sleep(time.Second)
		}

		fmt.Println("main: defer 2: After recover")
		time.Sleep(time.Second)
	}()
	defer func() {
		fmt.Println("main: defer 3")
		time.Sleep(time.Second)
	}()
	fmt.Println("main: Before a panic")
	time.Sleep(time.Second)
	panicFunc(5)
	fmt.Println("main: After a panic")
	time.Sleep(time.Second)

	defer func() {
		fmt.Println("main: defer 4")
		time.Sleep(time.Second)
	}()

	fmt.Println("main: The End")
	time.Sleep(time.Second)
}
