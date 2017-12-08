package words

import (
	"regexp"

	. "github.com/dghubble/go-twitter/twitter"
)

const (
	Update = `^(?!(RT\s|@)).+菜$`
)

var (
	NGWords = []string{`.*お\s*っ\s*ぱ\s*い\s*菜.*`, `.*脱法菜.*`, `.*うんこ.*`, `.*ちんこ.*`, `.*まんこ.*`, `.*麻薬.*`}
)

type Context struct {
	tweet  *Tweet
	client *Client
}

func (context *Context) Matched(pattern string) bool {
	r, err := regexp.MatchString(pattern, context.tweet.Text)
	if err == nil {
		return r
	}
	return false
}
