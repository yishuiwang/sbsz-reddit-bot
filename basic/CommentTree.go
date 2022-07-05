package basic

import (
	"fmt"
	"github.com/turnage/graw/reddit"
	"strings"
)

type Comment struct {
	Body     string
	Id       string
	ParentID string
}

func NewList(list []*Comment, comments []*reddit.Comment) []*Comment {
	for i := 0; i < len(comments); i++ {
		parentId := strings.Split(comments[i].ParentID, "_")
		list = append(list, &Comment{
			Body:     comments[i].Body,
			Id:       comments[i].ID,
			ParentID: parentId[1],
		})
		if comments[i].Replies != nil {
			list = NewList(list, comments[i].Replies)
		}
	}
	return list
}

func CommentTree(comments []*reddit.Comment) {
	var list []*Comment
	list = NewList(list, comments)
	printTree(list, list[0].ParentID, 0)
}

func printTree(list []*Comment, parent string, depth int) {
	for _, r := range list {
		if r.ParentID == parent {
			for i := 0; i < depth; i++ {
				if i == 0 {
					fmt.Print("|")
				}
				fmt.Print("--")
			}
			fmt.Print(r.Body, "\n")
			printTree(list, r.Id, depth+1)
		}
	}
}
