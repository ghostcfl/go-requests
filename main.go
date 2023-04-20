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
