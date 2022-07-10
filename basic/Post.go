package basic

import (
	"fmt"
	"github.com/turnage/graw/reddit"
	"net/url"
	"strconv"
	"strings"
)

type vote int

// Reddit interprets -1, 0, 1 as downvote, no vote, and upvote, respectively.
const (
	downvote vote = iota - 1
	novote
	upvote
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
	permalink := ParsePermalink(url)
	post, err := r.Bot.Thread(permalink)
	return post, err
}

// DeletePost 删除帖子
func (r Robot) DeletePost(id string) error {
	api := "/api/del"
	v := url.Values{}
	v.Set("id", id)

	request, err := NewHttpRequest("POST", api, strings.NewReader(v.Encode()))
	client := NewHttpClient(nil)
	_, err = client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	return err

}

func (r Robot) vote(id string, vote vote) error {
	path := "api/vote"

	form := url.Values{}
	form.Set("id", id)
	form.Set("dir", strconv.Itoa(int(vote)))
	form.Set("rank", "10")

	_, err := NewHttpRequest("POST", path, strings.NewReader(form.Encode()))
	return err
}

// UpVote a post or a comment.
func (r Robot) UpVote(id string) error {
	return r.vote(id, upvote)
}

// DownVote a post or a comment.
func (r Robot) DownVote(id string) error {
	return r.vote(id, downvote)
}

// RemoveVote removes your vote on a post or a comment.
func (r Robot) RemoveVote(id string) error {
	return r.vote(id, novote)
}

// Lock a post or comment, preventing it from receiving new comments.
func (r Robot) Lock(id string) error {
	path := "api/lock"

	form := url.Values{}
	form.Set("id", id)

	_, err := NewHttpRequest("POST", path, strings.NewReader(form.Encode()))

	return err
}

// Unlock a post or comment, allowing it to receive new comments.
func (r Robot) Unlock(id string) error {
	path := "api/unlock"

	form := url.Values{}
	form.Set("id", id)

	_, err := NewHttpRequest("POST", path, strings.NewReader(form.Encode()))

	return err
}
