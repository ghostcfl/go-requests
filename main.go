package main

import (
	"fmt"
	"newTest/requests"
)

func main() {
	session := requests.NewSession()

	//session.BaseUrl = "https://kawayiyi.com"
	session.BaseUrl = "https://httpbin.org"
	session.Headers = &requests.KV{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.48",
	}
	session.Cookies = &requests.KV{
		"a": "b",
	}

	resp, err := session.Request("/get", requests.P{
		Cookies: &requests.KV{
			"c": "d",
		},
		Headers: &requests.KV{
			"referer": "/post",
		},
	})
	if err != nil {
		panic(err)

	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-type"))
	fmt.Println(resp.Cookie)
	fmt.Println(resp.Text())

	if err != nil {
		panic(err)
		return
	}
}
