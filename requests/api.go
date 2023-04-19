package requests

import (
	"fmt"
	"io"
	"net/http"
)

func (session *Session) Request(url string, p P) (*Response, error) {
	if p.Method == "" {
		fmt.Println(p.Method)
		p.Method = "GET"
	}
	req, err := http.NewRequest(p.Method, url, nil)
	if err != nil {
		return nil, err
	}

	err = p.prepare(req)
	if p.Proxies != "" {
		err = prepareProxy(p.Proxies, session)
	}
	if p.NotAllowRedirects {
		session.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	if err != nil {
		return nil, err
	}

	resp, err := session.client.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)
	return &Response{
		Content:    respBody,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Cookie:     resp.Cookies(),
	}, nil
}
