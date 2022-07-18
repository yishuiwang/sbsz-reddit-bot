package basic

import (
	"errors"
	"io"
	"net/http"
	"time"
)

const (
	ConfigPath          = "config/reddit.json"
	OAuthEndpoint       = "https://oauth.reddit.com"
	TokenAccessEndpoint = "https://www.reddit.com/api/v1/access_token"
	UserAgent           = "<reddit>:<version 1.0.0> (by /u/sbsznmsl)"
)

type HttpClient struct {
	HttpOptions
	client  *http.Client
	header  map[string]string
	cookies []*http.Cookie
}

type HttpOptions struct {
	TryTime int
	Timeout time.Duration
}

func (c *HttpClient) SetHeader(key string, val string) {
	c.header[key] = val
}

// NewHttpRequest 创建http请求
func NewHttpRequest(method string, api string, body io.Reader) (*http.Request, error) {
	token, err := TokenAccess(ConfigPath)
	if err != nil {
		return nil, err
	}
	Request, err := http.NewRequest(method, OAuthEndpoint+api, body)
	Request.Header.Add("User-Agent", UserAgent)
	Request.Header.Add("Authorization", "Bearer "+token)
	return Request, err
}

// NewHttpClient 创建新Http请求器
func NewHttpClient(option *HttpOptions) *HttpClient {
	if option == nil {
		option = new(HttpOptions)
	}
	if option.TryTime == 0 {
		option.TryTime = 1
	}
	if option.Timeout == 0 {
		option.Timeout = 10 * time.Second
	}
	return &HttpClient{
		HttpOptions: *option,
		client:      &http.Client{Timeout: option.Timeout},
		header:      make(map[string]string),
	}
}

func (c *HttpClient) AddCookie(cookie ...*http.Cookie) {
	c.cookies = append(c.cookies, cookie...)
}

func (c HttpClient) Do(req *http.Request) (*http.Response, error) {
	var res *http.Response
	err := errors.New("TryTime is zero, send no http request")
	if req == nil {
		return nil, errors.New("req is nil")
	}
	// 设置 header
	for key, val := range c.header {
		req.Header.Add(key, val)
	}
	// 添加 cookie
	for _, cookie := range c.cookies {
		if cookie == nil {
			continue
		}
		req.AddCookie(cookie)
	}
	// 进行指定次数的重试
	for i := 0; i < c.TryTime; i++ {
		res, err = c.client.Do(req)
		if err == nil {
			break
		}
	}
	return res, err
}
