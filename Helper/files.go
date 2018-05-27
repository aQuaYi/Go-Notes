package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
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

func j(path, ext string) []string {
	ew := new(worker)
	ew.Mutex = new(sync.Mutex)

	sw := skywalker.New(path, ew)
	sw.ExtListType = skywalker.LTWhitelist
	sw.ExtList = []string{ext}
	err := sw.Walk()
	if err != nil {
		panic(err)
	}

	log.Printf("已经完成 %s 的扫描", sw.Root)

	// sort.Sort(sort.StringSlice(ew.found))
	sort.Slice(ew.found, func(i int, j int) bool {
		return len(ew.found[i]) > len(ew.found[j])
	})

	for _, f := range ew.found {
		fmt.Println(strings.Replace(f, sw.Root, "", 1))
	}
}
