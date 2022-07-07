package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sbsz-reddit-bot/config"
	"strings"
	"time"
)

const (
	TOKEN_ACCESS_ENDPOINT = "https://www.reddit.com/api/v1/access_token"
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

//func (c HttpClient) Do(req *http.Request) (*http.Response, error) {
//	//req, err := http.NewRequest(method, url, payload)
//	//if err != nil {
//	//	return nil, err
//	//}
//	//
//	//// Set the auth for the request.
//	//req.SetBasicAuth(username, password)
//	//
//	//return http.DefaultClient.Do(req)
//}

func TokenAccess(path string) {
	botconf, err := config.ReadJsonFile(path)
	if err != nil {
		fmt.Println(err)
	}
	data := url.Values{}
	data.Set("username", botconf.Username)
	data.Set("password", botconf.Password)
	data.Set("grant_type", "password")
	body := strings.NewReader(data.Encode())
	req, err := http.NewRequest("POST", TOKEN_ACCESS_ENDPOINT, body)
	req.Header.Add("User-Agent", "sbsz/api")

	if err != nil {
		fmt.Println(err)
	}
	req.SetBasicAuth(botconf.ClientId, botconf.ClientSecret)

	var c http.Client
	response, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
	fmt.Println(response.Status)
	fmt.Println(response)
}
