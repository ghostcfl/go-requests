package requests

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
}
