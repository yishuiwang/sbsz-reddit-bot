package basic

import (
	"fmt"
	"sbsz-reddit-bot/utils"
	"testing"
)

var (
	sbsz, _ = utils.NewRobot("config/config.json")
)

func TestReplyComment(t *testing.T) {
	err := sbsz.ReplyComment("套娃2", "165711449", "/r/sbsz/comments/vrr2jx/测试发帖/")
	if err != nil {
		fmt.Println("Failed to create bot handle: ", err)
		return
	}
}
