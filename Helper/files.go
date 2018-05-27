package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sync"

	"github.com/dixonwille/skywalker"
)

type worker struct {
	*sync.Mutex
	found []string
}

func (ew *worker) Work(path string) {
	ew.Lock()
	defer ew.Unlock()
	ew.found = append(ew.found, path)
}

// 扫描 path 路径下所有的 ext 后缀的文件
func scan(path, ext string) []string {
	ew := new(worker)
	ew.Mutex = new(sync.Mutex)

	sw := skywalker.New(path, ew)
	sw.ExtListType = skywalker.LTWhitelist
	sw.ExtList = []string{ext}
	err := sw.Walk()
	if err != nil {
		panic(err)
	}

	log.Printf("已经完成 %s 中 %s 后缀文件的扫描", sw.Root, ext)

	return ew.found

}

func getInformation(path string) (title, abstraction string) {
	f, err := os.Open(path) //打开文件
	if err != nil {
		log.Fatalf("打开 %s 时，出错: %s", path, err)
	}
	defer f.Close()

	buff := bufio.NewReader(f) //读入缓存

	tmp := [3]string{}

	for i := range tmp {
		tmp[i], err = buff.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || err == io.EOF {
			log.Fatalf("读取 %s 的第一行时，出错: %s", path, err)
		}
	}

	if len(tmp[0]) <= 2 {
		log.Fatalf("%s 中的标题长度不够", path)
	}
	title = tmp[0][2 : len(tmp[0])-1]

	if tmp[2][0] != '#' {
		abstraction = tmp[2][:len(tmp[2])-1]
	}

	return
}
