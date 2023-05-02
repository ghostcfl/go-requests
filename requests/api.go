package requests

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func setCookies(cookies KV, cookiesJar []*http.Cookie) {
	for _, cookie := range cookiesJar {
		if cookies.Get(cookie.Name) == "" {
			cookies.Set(cookie.Name, cookie.Value)
		}
	}
}

func (session *session) Request(url string, p P) (*Response, error) {
	var err error
	if p.Method == "" {
		p.Method = GET
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
			setCookies(session.Cookies, req.Response.Cookies())
			return http.ErrUseLastResponse
		}
	} else {
		session.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			setCookies(session.Cookies, req.Response.Cookies())
			return nil
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

	for _, fn := range session.beforeRequestHookFunctions {
		err = fn(req)
		if err != nil {
			break
		}
	}

	resp, err := session.client.Do(req)

	for _, fn := range session.afterResponseHookFunctions {
		err = fn(resp)
		if err != nil {
			break
		}
	}

	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ck := KV{}
	setCookies(ck, resp.Cookies())

	return &Response{
		Content:    respBody,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Cookie:     ck,
	}, nil
}

func (s *session) Get(url string, params P) (*Response, error) {
	params.Method = GET
	return s.Request(url, params)
}

func (s *session) Delete(url string, params P) (*Response, error) {
	params.Method = DELETE
	return s.Request(url, params)
}

func (s *session) Post(url string, params P) (*Response, error) {
	params.Method = POST
	return s.Request(url, params)
}

func (s *session) Put(url string, params P) (*Response, error) {
	params.Method = PUT
	return s.Request(url, params)
}

func (s *session) Head(url string, params P) (*Response, error) {
	params.Method = HEAD
	return s.Request(url, params)
}

func (s *session) Options(url string, params P) (*Response, error) {
	params.Method = OPTIONS
	return s.Request(url, params)
}

func (s *session) Patch(url string, params P) (*Response, error) {
	params.Method = PATCH
	return s.Request(url, params)
}

func Request(url string, params P) (*Response, error) {
	s := NewSession()
	return s.Request(url, params)
}

func Get(url string, params P) (*Response, error) {
	params.Method = GET
	return Request(url, params)
}

func Delete(url string, params P) (*Response, error) {
	params.Method = DELETE
	return Request(url, params)
}

func Post(url string, params P) (*Response, error) {
	params.Method = POST
	return Request(url, params)
}

func Put(url string, params P) (*Response, error) {
	params.Method = PUT
	return Request(url, params)
}

func Head(url string, params P) (*Response, error) {
	params.Method = HEAD
	return Request(url, params)
}

func Options(url string, params P) (*Response, error) {
	params.Method = OPTIONS
	return Request(url, params)
}
