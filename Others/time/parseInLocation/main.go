package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	//利用time.ParseInLocation把时间的字符串转换成time.Time格式

	strTime := "2017-06-18 15:36:12"
	t, err := time.ParseInLocation("2006-01-02 15:04:05", strTime, time.Local)
	if err != nil {
		msg := fmt.Sprintf("无法把%s转换成time.time格式: %s", strTime, err)
		log.Println(msg)
	}

	fmt.Println("时间的字符串为      ", strTime)
	fmt.Println("转换后的time.Time为", t)
}
