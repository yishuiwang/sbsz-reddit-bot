package main

import (
	"fmt"
	_ "github.com/turnage/graw"
	"sbsz-reddit-bot/basic"
)

func main() {
	sbsz, err := basic.NewRobot("config/config.json")
	if err != nil {
		fmt.Println("Failed to create bot handle: ", err)
		return
	}

	sbsz.DeletePost("t3_vthyma")

}
