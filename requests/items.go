package requests

import "time"

type KV map[string]string

type J map[string]interface{}

type L []interface{}

type F struct {
	Filename string
	Buffer   []byte
}

type Files map[string]F

type P struct {
	Method            string
	Params            *KV
	Data              *KV
	DataString        string
	Json              *J
	JsonString        string
	Headers           *KV
	Cookies           *KV
	Files             *Files
	Form              *KV
	Proxies           string
	NotAllowRedirects bool
	Timeout           time.Duration
}

type GP struct {
	Params            *KV
	Headers           *KV
	Cookies           *KV
	Proxies           string
	NotAllowRedirects bool
	Timeout           time.Duration
}

type PP struct {
	Params            *KV
	Data              *KV
	DataString        string
	Json              *J
	JsonString        string
	Headers           *KV
	Cookies           *KV
	Files             *Files
	Form              *KV
	Proxies           string
	NotAllowRedirects bool
	Timeout           time.Duration
}

func (k KV) Set(name string, value string) {
	k[name] = value
}

func (k KV) Get(name string) string {
	return k[name]
}
