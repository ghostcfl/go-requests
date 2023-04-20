## 这是一个类似于python requests 的请求库

```text
go version 1.20.1
```

### (一)特点

1. 和python requests一样的使用方法
2. 拥有随机TLS JA3指纹，可以通过一些JA3反爬的网站
3. 支持HTTP2请求，直接请求，无需额外参数
4. 其他还没想好

### (二)安装

```shell
go get github.com/ghostcfl/go-requests
```

### (三)使用
```go
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
```

### (四)JA3验证和H2请求
```go
package main

import (
	"fmt"
	"github.com/ghostcfl/go-requests/requests"
	"sync"
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

func main() {
	var wg sync.WaitGroup
	{
	}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ja3_check(fmt.Sprintf("协程%d", i))
		}(i)

	}
	wg.Wait()
}
```
```text
协程3请求的TlsHashMd5:87ce5f1229ae88c379f7c3dd01447677	协程3请求的Proto:HTTP/2
协程4请求的TlsHashMd5:39260f8e997e2452a118aa31930887c4	协程4请求的Proto:HTTP/2
协程1请求的TlsHashMd5:334373ba2e41842cde57f275cd2c6ad7	协程1请求的Proto:HTTP/2
协程0请求的TlsHashMd5:907ddc4b40855e3e820900940c1d539d	协程0请求的Proto:HTTP/2
协程2请求的TlsHashMd5:1e6ce23b75d203e2880b6c35a6039aef	协程2请求的Proto:HTTP/2
```
