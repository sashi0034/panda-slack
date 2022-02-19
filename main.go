package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/slack-go/slack"
)

type configData struct {
	BotToken    string `json:'botToken`
	PostChannel string `json:postChannel`
}

func loadConfig() (*configData, error) {
	f, err := os.Open("config.json")
	if err != nil {
		log.Fatal("loadConfig os.Open err:", err)
		return nil, err
	}
	defer f.Close()

	var cfg configData
	err = json.NewDecoder(f).Decode(&cfg)
	return &cfg, err
}

var config *configData

func main() {
	config, _ = loadConfig()
	// アクセストークンを使用してクライアントを生成する
	token := config.BotToken
	c := slack.New(token)

	// MsgOptionText() の第二引数に true を設定すると特殊文字をエスケープする
	_, _, err := c.PostMessage(config.PostChannel, slack.MsgOptionText("Hello World\nわいわい", false))
	if err != nil {
		panic(err)
	}
}
