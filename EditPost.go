package main

import (
	"sbsz-reddit-bot/utils"
)

func Post() {
	bot, _ := utils.NewBotFromFile("config/config.json", 0)
	bot.PostSelf("/r/sbsz", "测试发帖", "测试内容")

	bot.GetPostLink("/r/sbsz", "测试内容", "")
}
