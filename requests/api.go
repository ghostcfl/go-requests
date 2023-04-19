package requests

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func (session *Session) Request(url string, p P) (*Response, error) {
	var err error
	if p.Method == "" {
		fmt.Println(p.Method)
		p.Method = "GET"
	}

	if !strings.HasPrefix(url, "http") {
		if session.BaseUrl != "" {
			url = session.BaseUrl + url
		} else {
			return nil, errors.New(fmt.Sprintf("%s Scheme Error!", url))
		}
	}
	// 设置超时
	if p.Timeout > 0 {
		session.client.Timeout = p.Timeout * time.Second
	} else if session.Timeout > 0 {
		session.client.Timeout = session.Timeout * time.Second
	}
	// 设置代理
	if p.Proxies != "" {
		err = prepareProxy(p.Proxies, session)
	} else if session.Proxies != "" {
		err = prepareProxy(session.Proxies, session)
	}
	if err != nil {
		return nil, err
	}

	// 设置重定向
	if p.NotAllowRedirects {
		session.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	req, err := http.NewRequest(p.Method, url, nil)
	if err != nil {
		return nil, err
	}

	err = session.prepare(req)
	err = p.prepare(req)

	if err != nil {
		return nil, err
	}

	resp, err := session.client.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &Response{
		Content:    respBody,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Cookie:     resp.Cookies(),
	}, nil
}
