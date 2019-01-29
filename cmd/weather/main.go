package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net/http"
)

/*
	api.openweathermap.org/data/2.5/forecast?q={city name},{country code}
	{city name},{country code} は、以下参考
	http://www.asahi-net.or.jp/~ax2s-kmtn/ref/iso3166-1.html
*/

const endpoint = "http://api.openweathermap.org/data/2.5/forecast"

func main() {
	target := endpoint + "?q=%s,%s&APPID=%s"

	resp, err := http.Get(fmt.Sprintf(target, "London", "uk", "appid"))
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatal(errors.New("failed to request"))
	}

	sc := bufio.NewScanner(resp.Body)
	for sc.Scan() {
		fmt.Println(sc.Text())
	}

}
