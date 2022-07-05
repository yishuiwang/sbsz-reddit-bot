package main

import (
	"fmt"
	_ "github.com/turnage/graw"
	"sbsz-reddit-bot/basic"
	"sbsz-reddit-bot/utils"
)

func main() {
	bot, err := utils.NewBotFromFile("config/config.json", 0)
	if err != nil {
		fmt.Println("Failed to create bot handle: ", err)
		return
	}
	//sub, err := bot.GetPostLink("r/Megumin", "Megumin fantasy", "https://www.reddit.com/r/Megumin/comments/hbgrgk/megumin_fantasy")
	//fmt.Println(err)
	//fmt.Println(sub)
	//err = bot.Reply("sbsznmsl", "回复测试")
	//fmt.Println(err)

	//bot.PostSelf("/r/sbsz", "测试发帖", "测试内容")

	//if err := bot.PostLink("/r/sbsz", "PostLink测试", "https://www.reddit.com/r/sbsz/"); err != nil {
	//	fmt.Println(err)
	//}
	//

	post, err := bot.Thread("/r/sbsz/comments/vrr2jx/测试发帖/")
	if err != nil {
		fmt.Println(err)
	}
	//for _, reply := range post.Replies {
	//	fmt.Println(reply)
	//}

	basic.CommentTree(post.Replies)

}
