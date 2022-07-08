package basic

import (
	"fmt"
	"testing"
)

var (
	sbsz, _ = NewRobot("config/config.json")
)

func TestReplyComment(t *testing.T) {
	err := sbsz.ReplyComment("套娃2", "165711449", "/r/sbsz/comments/vrr2jx/测试发帖/")
	if err != nil {
		fmt.Println("Failed to create bot handle: ", err)
		return
	}
}
