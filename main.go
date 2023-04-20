package main

import (
	"fmt"
	"github.com/ghostcfl/go-requests/requests"
	"os"
)

func postFormAndFiles() {
	file, err := os.ReadFile("1.txt")
	if err != nil {
		return
	}
	resp, err := requests.Post("https://httpbin.org/post", requests.PP{
		Files: &requests.Files{
			"file1": requests.F{
				Filename: "hello.txt",
				Buffer:   file,
			},
		},
		Form: &requests.KV{
			"username": "caifuliang",
		},
	})
	if err != nil {
		return
	}

	fmt.Println(resp.Text())
}

func postJson() {
	// use J struct
	resp, err := requests.Post("https://httpbin.org/post", requests.PP{
		Json: &requests.J{
			"a": "b",
			"b": []string{"1", "2", "3"},
			"c": requests.KV{
				"c1":  "val c1",
				"c22": "val c2",
			},
		},
	})
	if err != nil {
		return
	}
	fmt.Println(resp.Text())
	// use JsonString
	resp, err = requests.Post("https://httpbin.org/post", requests.PP{
		JsonString: `{"a":"b","b":["1","2","3"],"c":{"c1":"val c1","c22":"val c2"}}`,
	})
	if err != nil {
		return
	}

	fmt.Println(resp.Text())

}

func postUrlencoded() {
	// use KV struct
	resp, err := requests.Post("https://httpbin.org/post", requests.PP{
		Data: &requests.KV{
			"a":    "b",
			"name": "caifuliang",
		},
	})
	if err != nil {
		return
	}
	fmt.Println(resp.Text())
	// use urlencoded string
	resp, err = requests.Post("https://httpbin.org/post", requests.PP{
		DataString: "a=b&name=caifuliang",
	})
	if err != nil {
		return
	}
	fmt.Println(resp.Text())
}

func main() {
	postUrlencoded()
}
