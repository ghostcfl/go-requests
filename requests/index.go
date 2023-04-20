package requests

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

type Session struct {
	client  *http.Client
	Headers *KV
	Cookies *KV
	BaseUrl string
	Timeout time.Duration
	Proxies string
}

const defaultUserAgent = "go-requests/0.0.1"

func NewSession() *Session {
	jar, _ := cookiejar.New(nil)
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			CipherSuites: getCipherSuites(),
		},
		ForceAttemptHTTP2: true,
	}

	return &Session{
		client: &http.Client{
			Jar:       jar,
			Transport: transport,
			Timeout:   0 * time.Second,
		},
	}
}

func getCipherSuites() []uint16 {
	CipherSuitesArray := []uint16{
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,
		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
		tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
		tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_RSA_WITH_AES_128_CBC_SHA,
		tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA,
		tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
		tls.TLS_AES_128_GCM_SHA256,
		tls.TLS_AES_256_GCM_SHA384,
		tls.TLS_CHACHA20_POLY1305_SHA256,
	}
	randIndex := rand.Intn(len(CipherSuitesArray))
	return CipherSuitesArray[:randIndex]
}

func dataMethod(method string) bool {
	switch strings.ToUpper(method) {
	case "POST":
		return true
	case "PUT":
		return true
	default:
		return false
	}

}

func (p *P) prepare(r *http.Request) error {
	var err error
	prepareHeaders(p.Headers, r)
	prepareQuery(p.Params, r)
	prepareCookies(p.Cookies, r)
	if dataMethod(p.Method) {
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
	}
	return err
}

func (s *Session) prepare(r *http.Request) error {
	prepareHeaders(s.Headers, r)
	prepareCookies(s.Cookies, r)
	return nil
}

func prepareHeaders(headers *KV, r *http.Request) {
	if headers != nil {
		for key, val := range *headers {
			r.Header.Set(key, val)
		}
		if r.Header.Get("user-agent") == "" {
			r.Header.Set("user-agent", defaultUserAgent)
		}
	} else {
		r.Header.Set("user-agent", defaultUserAgent)
	}
}

func prepareQuery(params *KV, r *http.Request) {
	if params != nil {
		query := r.URL.Query()
		for key, value := range *params {
			query.Set(key, value)
		}
		r.URL.RawQuery = query.Encode()
	}
}

func prepareData(data *KV, r *http.Request) {
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

func prepareJson(data *J, r *http.Request) error {
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

func prepareFiles(files *Files, form *KV, r *http.Request) error {
	var err error
	// Create a new buffer to write the request body
	requestBody := &bytes.Buffer{}
	// Create a new multipart writer with the request body buffer
	multipartWriter := multipart.NewWriter(requestBody)
	defer multipartWriter.Close()
	// 设置请求头
	fmt.Println(r.Header.Get("content-type"))
	if r.Header.Get("content-type") == "" {
		r.Header.Set("content-type", multipartWriter.FormDataContentType())
	}

	if files != nil {
		for fieldName, file := range *files {
			f := io.Reader(bytes.NewReader(file.Buffer))
			part, err := multipartWriter.CreateFormFile(fieldName, file.Filename)
			if err != nil {
				return err
			}
			_, err = io.Copy(part, f)
		}
	}

	if form != nil {
		for key, val := range *form {
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

func prepareCookies(cookies *KV, r *http.Request) {
	if cookies != nil {
		for key, val := range *cookies {
			ck := &http.Cookie{Name: key, Value: val}
			r.AddCookie(ck)
		}
	}
}

func prepareProxy(proxy string, session *Session) error {
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
