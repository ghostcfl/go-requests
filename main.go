package main

import (
	"fmt"
	"github.com/ghostcfl/go-requests/requests"
	"os"
)

type KaWaYiYiTls struct {
	Sni             string   `json:"sni"`
	TlsVersion      string   `json:"tlsVersion"`
	TcpConnectionId string   `json:"tcpConnectionId"`
	TlsHashOrigin   string   `json:"tlsHashOrigin"`
	TlsHashMd5      string   `json:"tlsHashMd5"`
	CipherList      []string `json:"cipherList"`
	Extentions      []string `json:"extentions"`
	Supportedgroups []string `json:"supportedgroups"`
	EcPointFormats  []string `json:"ecPointFormats"`
	Proto           string   `json:"proto"`
	H2              struct {
		SETTINGS struct {
			Field1 string `json:"2"`
			Field2 string `json:"4"`
			Field3 string `json:"6"`
		} `json:"SETTINGS"`
		WINDOWUPDATE string   `json:"WINDOW_UPDATE"`
		HEADERS      []string `json:"HEADERS"`
	} `json:"h2"`
	UserAgent string `json:"user_agent"`
	ClientIp  string `json:"clientIp"`
}

func ja3_check(name string) {
	resp, err := requests.Get("https://kawayiyi.com/tls", requests.GP{})
	if err != nil {
		panic(err)
	}
	var data KaWaYiYiTls
	err = resp.Json(&data)
	if err != nil {
		return
	}

	fmt.Printf("%s请求的TlsHashMd5:%s\t%s请求的Proto:%s\n", name, data.TlsHashMd5, name, data.Proto)

}

func postFormAndFiles() {
	file, err := os.ReadFile("1.txt")
	if err != nil {
		return
	}
	resp, err := requests.Post("https://httpbin.org/post", requests.PP{
		Files: requests.Files{
			"file1": requests.F{
				Filename: "hello.txt",
				Buffer:   file,
			},
		},
		Form: requests.KV{
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
		Json: requests.J{
			"a": "b",
			"b": []string{"1", "2", "3"},
			"c": requests.KV{
				"c1":  "val c1",
				"c22": "val c2",
			},
		},
	})
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(resp.Text())
	// use JsonString
	resp, err = requests.Post("https://httpbin.org/post", requests.PP{
		JsonString: `{"a":"b","b":["1","2","3"],"c":{"c1":"val c1","c22":"val c2"}}`,
	})
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(resp.Text())
}

func postUrlencoded() {
	// use KV struct
	resp, err := requests.Post("https://httpbin.org/post", requests.PP{
		Data: requests.KV{
			"a":    "b",
			"name": "caifuliang",
		},
	})
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(resp.Text())
	// use urlencoded string
	resp, err = requests.Post("https://httpbin.org/post", requests.PP{
		DataString: "a=b&name=caifuliang",
	})
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(resp.Text())
}

func useSession() {
	session := requests.NewSession()
	session.Headers = requests.KV{
		"user-agent": "my-ua",
	}
	session.Cookies = requests.KV{
		"token": "my-cookies-token",
	}
	session.BaseUrl = "https://httpbin.org"

	resp, err := session.Get("/cookies/set", requests.GP{
		Params: requests.KV{
			"free": "true",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Text())
	resp, err = session.Get("/get", requests.GP{})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Text())
	fmt.Println(session.Cookies)
}

func main() {
	postUrlencoded()
}
