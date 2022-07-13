package plugin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sbsz-reddit-bot/basic"
	"time"
)

const (
	LoliconAPI = "https://api.lolicon.app/setu/v2"
	proxy      = "i.pixiv.re"
)

type APIRequest struct {
	Num   int      `json:"num,omitempty"`
	Tags  []string `json:"tag,omitempty"`
	Size  []string `json:"size,omitempty"`
	Proxy string   `json:"proxy,omitempty"`
	R18   int      `json:"r18"`
}

func DownLoadImage() {
	tags := []string{"原神"}
	rule := APIRequest{
		Num:   1,
		Tags:  tags,
		Size:  []string{"original"},
		Proxy: proxy,
		R18:   1,
	}
	body, _ := json.Marshal(rule)
	request, err := http.NewRequest("POST", LoliconAPI, bytes.NewReader(body))
	if err != nil {
		fmt.Println(err)
	}
	c := basic.NewHttpClient(&basic.HttpOptions{TryTime: 3, Timeout: 15 * time.Second})
	c.SetHeader("Content-Type", "application/json")
	response, err := c.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))

}
