package basic

import (
	"encoding/json"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"io/ioutil"
)

var (
	tail = "——[I am a robot](https://github.com/yishuiwang/sbsz-reddit-bot)"
)

type botconf struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

// NewBot 返回Robot
func NewBot() (*reddit.Client, error) {
	var conf botconf
	data, _ := ioutil.ReadFile("config/reddit.json")
	json.Unmarshal(data, &conf)
	credentials := reddit.Credentials{ID: conf.ClientId, Secret: conf.ClientSecret, Username: conf.Username, Password: conf.Password}
	client, err := reddit.NewClient(credentials)
	return client, err
}
