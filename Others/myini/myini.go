// ini.go
package main

import (
	"fmt"

	"github.com/go-ini/ini"
)

func main() {
	cfg, err := ini.Load("raw.ini")
	if err == nil {
		fmt.Println(cfg)
	}
	//	fmt.Println(cfg.)
	secOkcn, err := cfg.GetSection("okcoin.cn")
	if err == nil {
		fmt.Println(secOkcn)
	}
	apikeyOkcn, err := secOkcn.GetKey("apikey")
	if err == nil {
		fmt.Println((apikeyOkcn))
		fmt.Printf("%T\n", apikeyOkcn)
	}

	fmt.Println(cfg.Section("okcoin.cn").Keys())
	fmt.Println(cfg.Section("okcoin.cn").KeyStrings())
	fmt.Printf("%T,%v\n", secOkcn.Key("apikey").Value(), secOkcn.Key("apikey").Value())

	apikeyOkcn.SetValue("ccccccccccccccccccccccc")

	err = cfg.SaveTo("cooked.ini")
	if err != nil {
		fmt.Println(err)
	}

}
