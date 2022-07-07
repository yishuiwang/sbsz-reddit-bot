package config

import (
	"encoding/json"
	"fmt"
	"github.com/turnage/graw/reddit"
	"io/ioutil"
	"os"
	"sbsz-reddit-bot/basic"
	"time"
)

type BotConfig struct {
	UserAgent    string `json:"user_agent"`
	ClientId     string `json:"client_id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientSecret string `json:"client_secret"`
}

// NewRobot 返回Robot
func NewRobot(filename string) (basic.Robot, error) {
	b, err := NewBotFromFile(filename, 0)
	return basic.Robot{Bot: b}, err
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
