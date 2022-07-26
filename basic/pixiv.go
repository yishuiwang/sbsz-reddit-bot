package basic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

const (
	API      = "https://api.lolicon.app/setu/v2"
	Proxy    = "i.pixiv.re"
	FilePath = "./data/"
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

func GetImage(tags []string) error {
	rule := APIRequest{
		Num:   6,
		Tags:  tags,
		Size:  []string{"original"},
		Proxy: Proxy,
		R18:   0,
	}
	body, _ := json.Marshal(rule)
	request, err := http.NewRequest("POST", API, bytes.NewReader(body))
	if err != nil {
		return err
	}
	c := http.Client{}
	//c.SetHeader("Content-Type", "application/json")
	response, err := c.Do(request)
	if err != nil {
		return err
	}
	data, _ := ioutil.ReadAll(response.Body)
	responseInfo := ResponseInfo{}
	err = json.Unmarshal(data, &responseInfo)
	if err != nil {
		return err
	}

	var filename []string
	for i := 0; i < len(responseInfo.Data); i++ {
		pid := responseInfo.Data[i].Pid
		filename = append(filename, strconv.FormatInt(pid, 10)+".jpg")
	}

	for i := 0; i < len(responseInfo.Data); i++ {
		downLoadImage(FilePath+filename[i], responseInfo.Data[i].Urls["original"])
	}
	return nil
}

func downLoadImage(filename, path string) {

	img, err := http.Get(path)
	if err != nil {
		fmt.Println(err)
	}
	content, _ := ioutil.ReadAll(img.Body)
	defer img.Body.Close()
	out, err := os.Create(filename)
	defer out.Close()

	if err != nil {
		fmt.Println(err)
	}
	out.Write(content)
	if err != nil {
		fmt.Println(err)
	}

}
