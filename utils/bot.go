package utils

import (
	"encoding/json"
	"fmt"
	"github.com/turnage/graw/reddit"
	"io/ioutil"
	"os"
	"time"
)

type BotConfig struct {
	UserAgent    string `json:"user_agent"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
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
