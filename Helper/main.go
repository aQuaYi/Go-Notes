package main

import (
	"fmt"
	"sort"
)

func main() {

	// folders := []string{
	// 	"Go",
	// 	"Libraries",
	// 	"Algorithms",
	// 	"DesignPattern",
	// 	"Others",
	// }

	tmpls := scan(".", ".markdown")

	sort.Slice(tmpls, func(i int, j int) bool {
		return len(tmpls[i]) > len(tmpls[j])
	})

	for i := range tmpls {
		fmt.Println(tmpls[i])
	}

	path := "./Algorithms/sort/quick/README.md"
	a, b := getInformation(path)
	fmt.Printf("%q %q", a, b)

	path = "./Algorithms/sort/merge/README.md"
	a, b = getInformation(path)
	fmt.Printf("%q %q", a, b)
}
