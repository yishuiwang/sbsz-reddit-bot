package basic

import (
	"fmt"
	"github.com/turnage/graw/reddit"
)

type Comment struct {
	Body     string
	Id       string
	ParentID string
}

// NewList 返回一个简化版本的reddit comments
func NewList(list []*Comment, comments []*reddit.Comment) []*Comment {
	for i := 0; i < len(comments); i++ {
		list = append(list, &Comment{
			Body:     comments[i].Body,
			Id:       "t1_" + comments[i].ID,
			ParentID: comments[i].ParentID,
		})
		if comments[i].Replies != nil {
			list = NewList(list, comments[i].Replies)
		}
	}
	return list
}

// CommentTree todo 格式优化，显示问题，文本消息过多的处理
func CommentTree(p *reddit.Post) {
	var list []*Comment
	list = NewList(list, p.Replies)
	printTree(list, list[0].ParentID, 0)
}

// 根据父节点和子节点构造评论树
func printTree(list []*Comment, parent string, depth int) {
	for _, r := range list {
		if r.ParentID == parent {
			for i := 0; i < depth; i++ {
				if i == 0 {
					fmt.Print("|")
				}
				fmt.Print("--")
			}
			fmt.Printf("%s\n", r.Body)
			printTree(list, r.Id, depth+1)
		}
	}
}

// ReplyComment 根据title内容自动回复
func (r Robot) ReplyComment(title string, reply string, postUrl string) error {
	post, err := r.GetPostInfo(postUrl)
	if err != nil {
		return err
	}
	var list []*Comment
	list = NewList(list, post.Replies)
	text := reply + "\n\n" + tail
	for i := 0; i < len(list); i++ {
		if list[i].Body == title {
			if err := r.Bot.Reply(list[i].Id, text); err != nil {
				return err
			}
		}
	}
	return nil
}

// DeleteComment 删除回复
func (r Robot) DeleteComment() {

}
