package main

import (
	"fmt"
	"github.com/ghostcfl/go-requests/requests"
)

func main() {
	resp, err := requests.Get("https://httpbin.org/get", requests.GP{
		Headers: &requests.KV{
			"token":      "token1",
			"user-agent": "my-user-agent",
		},
	})
	if err != nil {
		panic(err)

	}
	fmt.Println(resp.Text())
}
