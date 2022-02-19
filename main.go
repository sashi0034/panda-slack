package main

import (
	"github.com/slack-go/slack"
)

func main() {
	// アクセストークンを使用してクライアントを生成する
	token := "xoxb-3069876617-3130854091778-fMxQjiXWMCIdLqclz5lON4aM"
	c := slack.New(token)

	// MsgOptionText() の第二引数に true を設定すると特殊文字をエスケープする
	_, _, err := c.PostMessage("C02BFCYV1DF", slack.MsgOptionText("Hello World\nわいわい", false))
	if err != nil {
		panic(err)
	}
}
