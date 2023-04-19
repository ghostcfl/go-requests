package requests

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Content    []byte
	StatusCode int
	Header     http.Header
	Cookie     []*http.Cookie
}

func (r *Response) Text() string {
	return string(r.Content)
}

func (r *Response) Json(result any) error {
	/*
		To Struct
		response text like {"status": "1", "state": "success", "data": [{"value": 7396}, {"value": 5018}]}
		type vv struct {
			Value int `json:"value"`
		}
		type rr struct {
			Status string `json:"status"`
			State  string `json:"state"`
			Data   []vv   `json:"data"`
		}
		var res rr
		_ := resp.Json(&res)
		fmt.Println(res)

		To Map
		var res map[string]interface
		_ := resp.Json(&res)
		fmt.Println(res)
	*/
	return json.Unmarshal(r.Content, &result)
}
