package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"./config/words"
	"./keys"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	fmt.Println("starting update-sai...")
	var config = oauth1.NewConfig(keys.ConsumerKey, keys.ConsumerSecret)
	var token = oauth1.NewToken(keys.AccessToken, keys.AccessTokenSecret)

	var httpClient = config.Client(oauth1.NoContext, token)

	var twitterClient = twitter.NewClient(httpClient)

	const wordMatch = regexp.Compile(words.Update)
	demux := twitter.NewSwitchDemux()

	// ここでワード処理
	demux.Tweet = func(tweet *twitter.Tweet) {
		fmt.Printf("%s: %s\n", tweet.User.Name, tweet.Text)

	}

	params := &twitter.StreamUserParams{
		With:          "followings",
		StallWarnings: twitter.Bool(true),
	}

	stream, err := twitterClient.Streams.User(params)

	if err == nil {
		go demux.HandleChan(stream.Messages)
	} else {
		fmt.Println("Failed get stream... exit.")
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")
	stream.Stop()

	fmt.Println("finished.")
}

type Tweet struct {
	text string
	user struct {
		screen_name string
	}
}
