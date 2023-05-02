package requests

import (
	"net/http"
	"time"
)

type (
	KV map[string]string
	J  map[string]interface{}
	// 列表结构体
	L []interface{}

	F struct {
		Filename string
		Buffer   []byte
	}
	// 文件类型结构体
	Files map[string]F
	// 请求参数
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

	BeforeRequestHookFunction func(*http.Request) error
	AfterResponseHookFunction func(*http.Response) error
)

func (k KV) Set(name string, value string) {
	k[name] = value
}

func (k KV) Get(name string) string {
	return k[name]
}

// 注册请求前的HOOK函数,返回注册成功的 index
func (s *session) RegisterBeforeRequestHook(fn BeforeRequestHookFunction) (int, error) {
	if s.beforeRequestHookFunctions == nil {
		s.beforeRequestHookFunctions = make([]BeforeRequestHookFunction, 0, 8)
	}
	if len(s.beforeRequestHookFunctions) > 7 {
		return -1, ErrHookFuncMaxLimit
	}
	s.beforeRequestHookFunctions = append(s.beforeRequestHookFunctions, fn)
	return len(s.beforeRequestHookFunctions) - 1, nil
}

// 注销请求前的HOOK函数
func (s *session) UnregisterBeforeRequestHook(index int) error {
	if index >= len(s.beforeRequestHookFunctions) {
		return ErrIndexOutOfBound
	}
	s.beforeRequestHookFunctions = append(s.beforeRequestHookFunctions[:index], s.beforeRequestHookFunctions[index+1:]...)
	return nil
}

// 重置请求前的HOOK函数
func (s *session) ResetBeforeRequestHook() {
	s.beforeRequestHookFunctions = []BeforeRequestHookFunction{}
}

// 注册响应后的HOOK函数,返回注册成功的 index
func (s *session) RegisterAfterResponseHook(fn AfterResponseHookFunction) (int, error) {
	if s.afterResponseHookFunctions == nil {
		s.afterResponseHookFunctions = make([]AfterResponseHookFunction, 0, 8)
	}
	if len(s.beforeRequestHookFunctions) > 7 {
		return -1, ErrHookFuncMaxLimit
	}
	s.afterResponseHookFunctions = append(s.afterResponseHookFunctions, fn)
	return len(s.afterResponseHookFunctions) - 1, nil
}

// 注销响应后的HOOK函数
func (s *session) UnregisterAfterResponseHook(index int) error {
	if index >= len(s.afterResponseHookFunctions) {
		return ErrIndexOutOfBound
	}
	s.afterResponseHookFunctions = append(s.afterResponseHookFunctions[:index], s.afterResponseHookFunctions[index+1:]...)
	return nil
}

// 重置响应后的HOOK函数
func (s *session) ResetAfterResponseHook() {
	s.afterResponseHookFunctions = []AfterResponseHookFunction{}
}
