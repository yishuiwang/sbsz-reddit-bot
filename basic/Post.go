package basic

import (
	"fmt"
	"github.com/turnage/graw/reddit"
	"sbsz-reddit-bot/utils"
)

//// PostLink 发送一篇带链接的帖子
//func PostLink(community string, title string, link string) {
//	//if err := bot.PostLink("/r/sbsz", "PostLink测试", "https://www.reddit.com/r/sbsz/"); err != nil {
//	//	fmt.Println(err)
//	//}
//}

// GetPostInfo 获取帖子的所有信息
func (r Robot) GetPostInfo(url string) *reddit.Post {
	permalink := utils.ParsePermalink(url)
	post, err := r.Bot.Thread(permalink)
	if err != nil {
		fmt.Println(err)
	}
	return post
}
