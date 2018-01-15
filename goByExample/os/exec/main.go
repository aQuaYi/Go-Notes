package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// LookPath 会在PATH中查询参数所表示的可执行文件的路径
	goBinaryPath, err := exec.LookPath("go")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Go binary path is", goBinaryPath)

	// 查询不存在的可执行文件
	noExistFilePath, err := exec.LookPath("noExistFile")
	if err != nil {
		// 能够执行，就不会报错
		log.Println(err)
	}
	// 但是返回的结果，会是一段报错信息。
	fmt.Println("noExistFilePath is", noExistFilePath)

	zshCmd := exec.Command("bash", "-c", "git push && git checkout master && git merge develop && git push && git checkout develop")
	// err = zshCmd.Run()
	// if err != nil {
	// 	log.Println(err)
	// }
	str, err := zshCmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(str))
}
