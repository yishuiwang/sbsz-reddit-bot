package basic

import (
	"github.com/turnage/graw/reddit"
	"sbsz-reddit-bot/utils"
)

// PostLink 发送一篇带链接的帖子
func (r Robot) PostLink(subreddit string, title string, link string) error {
	err := r.Bot.PostLink(subreddit, title, link)
	return err
}

// Post 普通发帖
func (r Robot) Post(subreddit string, title string, text string) error {
	err := r.Bot.PostSelf(subreddit, title, text)
	return err
}

// GetPostInfo 获取帖子的所有信息
func (r Robot) GetPostInfo(url string) (*reddit.Post, error) {
	permalink := utils.ParsePermalink(url)
	post, err := r.Bot.Thread(permalink)
	return post, err
}

// DeletePost 删除帖子
func (r Robot) DeletePost() {

}
