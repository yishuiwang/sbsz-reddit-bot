package basic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	LoliconAPI = "https://api.lolicon.app/setu/v2"
	Proxy      = "i.pixiv.re"
	FilePath   = "./data/"
)

type ResponseInfo struct {
	Error string    `json:"error"`
	Data  []Picture `json:"data"`
}

type Picture struct {
	Pid        int64             `json:"pid"`
	P          int               `json:"p"`
	Uid        int64             `json:"uid"`
	Title      string            `json:"title"`
	Author     string            `json:"author"`
	R18        bool              `json:"r18"`
	Width      int               `json:"width"`
	Height     int               `json:"height"`
	Tags       []string          `json:"tags"`
	Ext        string            `json:"ext"`
	UploadDate int               `json:"uploadDate"`
	Urls       map[string]string `json:"urls"`
}

type APIRequest struct {
	Num   int      `json:"num,omitempty"`
	Tags  []string `json:"tag,omitempty"`
	Size  []string `json:"size,omitempty"`
	Proxy string   `json:"proxy,omitempty"`
	R18   int      `json:"r18"`
}

func GetImageInfo(tags []string) (ResponseInfo, error) {
	rule := APIRequest{
		Num:   1,
		Tags:  tags,
		Size:  []string{"original"},
		Proxy: Proxy,
		R18:   2,
	}
	body, _ := json.Marshal(rule)
	request, err := http.NewRequest("POST", LoliconAPI, bytes.NewReader(body))
	if err != nil {
		return ResponseInfo{}, err
	}
	c := NewHttpClient(&HttpOptions{TryTime: 3, Timeout: 15 * time.Second})
	c.SetHeader("Content-Type", "application/json")
	response, err := c.Do(request)
	if err != nil {
		return ResponseInfo{}, err
	}
	data, _ := ioutil.ReadAll(response.Body)
	responseInfo := ResponseInfo{}
	err = json.Unmarshal(data, &responseInfo)
	if err != nil {
		return ResponseInfo{}, err
	}
	return responseInfo, nil
}

func DownLoadImage(info ResponseInfo) {
	var filename []string
	for i := 0; i < len(info.Data); i++ {
		pid := info.Data[i].Pid
		filename = append(filename, strconv.FormatInt(pid, 10)+".jpg")
	}
	for i := 0; i < len(info.Data); i++ {
		url := info.Data[i].Urls["original"]
		img, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		content, _ := ioutil.ReadAll(img.Body)
		out, err := os.Create(FilePath + filename[i])
		if err != nil {
			fmt.Println(err)
		}
		out.Write(content)
		if err != nil {
			fmt.Println(err)
		}
	}

}
