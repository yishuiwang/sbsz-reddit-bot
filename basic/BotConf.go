package basic

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/turnage/graw/reddit"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Robot struct {
	Bot reddit.Bot
}

var (
	tail = "——[I am a robot](https://github.com/yishuiwang/sbsz-reddit-bot)"
)

type BotConfig struct {
	UserAgent    string `json:"user_agent"`
	ClientId     string `json:"client_id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientSecret string `json:"client_secret"`
}

// NewRobot 返回Robot
func NewRobot(filename string) (Robot, error) {
	b, err := NewBotFromFile(filename, 0)
	return Robot{Bot: b}, err
}

func NewBotFromFile(filename string, rate time.Duration) (reddit.Bot, error) {
	conf, _ := ReadJsonFile(filename)

	app := reddit.App{
		ID:       conf.ClientId,
		Secret:   conf.ClientSecret,
		Username: conf.Username,
		Password: conf.Password,
	}

	agent := conf.UserAgent

	return reddit.NewBot(
		reddit.BotConfig{
			Agent: agent,
			App:   app,
			Rate:  rate,
		},
	)
}

// ReadJsonFile 从json文件读取配置
func ReadJsonFile(path string) (BotConfig, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	body, _ := ioutil.ReadAll(jsonFile)
	config := BotConfig{}
	json.Unmarshal(body, &config)
	return config, err
}

func TokenAccess(path string) (string, error) {
	botConf, err := ReadJsonFile(path)
	if err != nil {
		return "", err
	}
	data := url.Values{}
	data.Set("username", botConf.Username)
	data.Set("password", botConf.Password)
	data.Set("grant_type", "password")
	body := strings.NewReader(data.Encode())
	req, _ := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", body)
	req.Header.Add("User-Agent", "<reddit>:<version 1.0.0> (by /u/sbsznmsl)")

	req.SetBasicAuth(botConf.ClientId, botConf.ClientSecret)

	var c http.Client
	response, err := c.Do(req)
	if err != nil {
		return "", err
	}
	bytes, _ := ioutil.ReadAll(response.Body)
	token := gjson.Get(string(bytes), "access_token")
	return token.String(), nil
}
