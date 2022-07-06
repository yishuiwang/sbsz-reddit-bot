package utils

import (
	"strings"
)

// ParsePermalink 将帖子url转为permalink形式
func ParsePermalink(url string) string {
	slice := strings.Split(url, "reddit.com")
	return slice[1]
}
