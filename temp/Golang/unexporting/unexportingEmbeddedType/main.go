package main

import (
	"ShowYouTheGolangCode/Golang/unexporting/unexportingEmbeddedType/entities"
	"fmt"
)

func main() {
	a := entities.Admin{
		Rights: 10,
	}

	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Println(a)
}
