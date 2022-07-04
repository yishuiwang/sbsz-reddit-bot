package main

import (
	"fmt"
	_ "github.com/turnage/graw"
	"sbsz-reddit-bot/utils"
)

func main() {
	bot, err := utils.NewBotFromFile("config/config.json", 0)
	if err != nil {
		fmt.Println("Failed to create bot handle: ", err)
		return
	}

	harvest, err := bot.Listing("/r/sbsz", "")
	if err != nil {
		fmt.Println("Failed to fetch /r/golang: ", err)
		return
	}

	for _, post := range harvest.Posts {
		fmt.Printf("[%s] posted [%s]\n", post.Author, post.Title)
	}

}
