package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ddddddO/disaster/lib"
)

/*
	api.openweathermap.org/data/2.5/forecast?q={city name},{country code}
	{city name},{country code} は、以下参考
	http://www.asahi-net.or.jp/~ax2s-kmtn/ref/iso3166-1.html
*/

var apiKey string

const endpoint = "http://api.openweathermap.org/data/2.5/forecast"

func main() {
	flag.StringVar(&apiKey, "key", "", "OpenWeatherMap API Key")
	flag.Parse()

	target := endpoint + "?q=%s,%s&APPID=%s"
	url := fmt.Sprintf(target, "London", "uk", apiKey)

	c := lib.NewClient(url)
	resp, err := c.Get()
	if err != nil {
		log.Fatalf("stack trace: %+v\n", err)
	}

	w := Weather{}
	err = lib.Unmarshal(resp.Body, &w)
	if err != nil {
		log.Fatalf("stack trace: %+v\n", err)
	}

	fmt.Println(w.City)
	fmt.Println(w.Message)
	fmt.Println(w.List[0])
}
