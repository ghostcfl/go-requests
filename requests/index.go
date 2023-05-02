package requests

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

func (p *P) prepare(r *http.Request) error {
	var err error
	prepareHeaders(p.Headers, r)
	prepareQuery(p.Params, r)
	prepareCookies(p.Cookies, r)
	if p.Data != nil || p.DataString != "" {
		if r.Header.Get("content-type") == "" {
			r.Header.Set("content-type", "application/x-www-form-urlencoded")
		}
		prepareData(p.Data, r)
		prepareDataString(p.DataString, r)
	} else if p.Json != nil || p.JsonString != "" {
		if r.Header.Get("content-type") == "" {
			r.Header.Set("content-type", "application/json")
		}
		prepareJsonString(p.JsonString, r)
		err = prepareJson(p.Json, r)
		if err != nil {
			return err
		}
	} else if p.Files != nil || p.Form != nil {
		err = prepareFiles(p.Files, p.Form, r)
	}
	return err
}

func (s *session) prepare(r *http.Request) error {
	prepareHeaders(s.Headers, r)
	prepareCookies(s.Cookies, r)
	return nil
}

func prepareHeaders(headers KV, r *http.Request) {
	if headers != nil {
		for key, val := range headers {
			r.Header.Set(key, val)
		}
		if r.Header.Get("user-agent") == "" {
			r.Header.Set("user-agent", defaultUserAgent)
		}
	} else {
		if r.Header.Get("user-agent") == "" {
			r.Header.Set("user-agent", defaultUserAgent)
		}
	}
}

func prepareQuery(params KV, r *http.Request) {
	if params != nil {
		query := r.URL.Query()
		for key, value := range params {
			query.Set(key, value)
		}
		r.URL.RawQuery = query.Encode()
	}
}

func prepareData(data KV, r *http.Request) {
	if data != nil {
		d := MapToQueryString(data)
		r.Body = stringToReadCloser(d)
	}
}

func prepareDataString(s string, r *http.Request) {
	if s != "" {
		r.Body = stringToReadCloser(s)
	}
}

func prepareJson(data J, r *http.Request) error {
	if data != nil {
		j, err := MapToJsonString(data)
		if err != nil {
			return err
		}
		r.Body = stringToReadCloser(j)
	}
	return nil
}

func prepareJsonString(s string, r *http.Request) {
	if s != "" {
		r.Body = stringToReadCloser(s)
	}
}

func prepareFiles(files Files, form KV, r *http.Request) error {
	var err error
	// Create a new buffer to write the request body
	requestBody := &bytes.Buffer{}
	// Create a new multipart writer with the request body buffer
	multipartWriter := multipart.NewWriter(requestBody)
	defer multipartWriter.Close()
	// 设置请求头
	if r.Header.Get("content-type") == "" {
		r.Header.Set("content-type", multipartWriter.FormDataContentType())
	}

	if files != nil {
		for fieldName, file := range files {
			f := io.Reader(bytes.NewReader(file.Buffer))
			part, err := multipartWriter.CreateFormFile(fieldName, file.Filename)
			if err != nil {
				return err
			}
			_, err = io.Copy(part, f)
		}
	}

	if form != nil {
		for key, val := range form {
			formFieldWriter, err := multipartWriter.CreateFormField(key)
			if err != nil {
				return err
			}

			_, err = formFieldWriter.Write([]byte(val))
			if err != nil {
				return err
			}
		}
	}

	r.Body = io.NopCloser(requestBody)

	return err
}

func prepareCookies(cookies KV, r *http.Request) {
	if cookies != nil {
		var ckList []string
		for key, val := range cookies {
			ckList = append(ckList, fmt.Sprintf("%s=%s", key, val))
		}
		r.Header.Add("Cookie", strings.Join(ckList, "; "))
	}
}

func prepareProxy(proxy string, session *session) error {
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		return err
	}

	// create a new transport with the proxy URL
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	session.client.Transport = transport

	return nil
}
