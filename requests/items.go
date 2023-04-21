package requests

import "time"

type (
	KV map[string]string
	J  map[string]interface{}
	L  []interface{}
	F  struct {
		Filename string
		Buffer   []byte
	}
	Files map[string]F

	P struct {
		Method            string
		Params            KV
		Data              KV
		DataString        string
		Json              J
		JsonString        string
		Headers           KV
		Cookies           KV
		Files             Files
		Form              KV
		Proxies           string
		NotAllowRedirects bool
		Timeout           time.Duration
	}

	GP struct {
		Params            KV
		Headers           KV
		Cookies           KV
		Proxies           string
		NotAllowRedirects bool
		Timeout           time.Duration
	}

	PP struct {
		Params            KV
		Data              KV
		DataString        string
		Json              J
		JsonString        string
		Headers           KV
		Cookies           KV
		Files             Files
		Form              KV
		Proxies           string
		NotAllowRedirects bool
		Timeout           time.Duration
	}
)

func (k KV) Set(name string, value string) {
	k[name] = value
}

func (k KV) Get(name string) string {
	return k[name]
}
