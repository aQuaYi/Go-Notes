// fasthttp.go
package main

import (
	"fmt"

	fh "github.com/valyala/fasthttp"
)

func main() {
	url := "https://www.okcoin.cn/api/v1/ticker.do?symbol=btc_cny"

	statusCode, body, err := fh.Get(nil, url)

	if err != nil {
		fmt.Println(err)
	}
	if statusCode != 200 {
		fmt.Println("Status Code is NOT 200.")
	}
	fmt.Println("body:", string(body))

}
