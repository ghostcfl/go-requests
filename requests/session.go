package requests

import (
	"crypto/tls"
	"math/rand"
	"net/http"
	"time"
)

type Session struct {
	client  *http.Client
	Headers KV
	Cookies KV

	BaseUrl  string
	Timeout  time.Duration
	Proxies  string
	MaxRetry int

	beforeRequestHookFunctions []BeforeRequestHookFunction
	afterResponseHookFunctions []AfterResponseHookFunction
}

func NewSession() *Session {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			CipherSuites: getCipherSuites(),
		},
		ForceAttemptHTTP2: true,
	}

	return &Session{
		client: &http.Client{
			Transport: transport,
			Timeout:   0 * time.Second,
		},
		Headers: KV{},
		Cookies: KV{},

		beforeRequestHookFunctions: []BeforeRequestHookFunction{},
		afterResponseHookFunctions: []AfterResponseHookFunction{},
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
