package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=csH5bu8IiFm0CovWr0oY8UkJRm8SnKJAFyQ4ZPZmN2oK26S-esMZJZye9vwiyl61pvBPjbKFJhJXZ_klP_I3yGmcb4bKbertyiinRPR5CjTCWY1hGkn3ZW-3XAG2RjSrRnxJARfHvQ30cjT1BpNAGOHy2ikpjP39DK3EHvjGFoeDZSOghb6ISHfjIVNUka7XMoi4kQTACELLiVFSU1bzBA&userid=LiuYan"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
