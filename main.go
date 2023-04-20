package main

import (
	"fmt"
	"newTest/requests"
)

func main() {
	resp, _ := requests.Get("https://httpbin.org/get", requests.GP{})
	fmt.Println(resp.Text())
}
