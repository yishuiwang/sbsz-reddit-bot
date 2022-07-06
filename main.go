package main

import (
	"fmt"
	_ "github.com/turnage/graw"
	"sbsz-reddit-bot/basic"
	"sbsz-reddit-bot/utils"
)

func main() {
	sbsz, err := utils.NewRobot("config/config.json")
	if err != nil {
		fmt.Println("Failed to create bot handle: ", err)
		return
	}

	p := sbsz.GetPostInfo("/r/sbsz/comments/vrr2jx/测试发帖/")
	basic.CommentTree(p)
}
