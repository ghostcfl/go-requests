package main

import (
	"fmt"
	"newTest/requests"
)

func main() {
	session := requests.NewSession()

	//resp, err := session.Request("https://kawayiyi.com/tls", requests.P{
	resp, err := session.Request("https://match.yuanrenxue.cn/api/match/19?page=1", requests.P{
		Headers: &requests.KV{
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.48",
		},
	})
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-type"))
	fmt.Println(resp.Cookie)
	fmt.Println(resp.Text())

	if err != nil {
		panic(err)
		return
	}
}
