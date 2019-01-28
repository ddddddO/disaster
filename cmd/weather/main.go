package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
	api.openweathermap.org/data/2.5/forecast?q={city name},{country code}
	{city name},{country code} は、以下参考
	http://www.asahi-net.or.jp/~ax2s-kmtn/ref/iso3166-1.html

	NOTE:APIkey使えるようになるまで時間かかるっぽい
*/

const endpoint = "api.openweathermap.org/data/2.5/forecast"

func main() {
	target := endpoint + "?q=%s,%s&APPID=%s"

	resp, err := http.Get(fmt.Sprintf(target, "London", "uk", "<apiid>"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
