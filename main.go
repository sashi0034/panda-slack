package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/TKMAX777/panda"
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

	Panda := panda.NewClient()

	pandaErr := Panda.Login(os.Getenv("ECSID"), os.Getenv("ECSID_PW"))
	if pandaErr != nil {
		panic(pandaErr)
	}
	//ass, _ := Panda.GetAssignment()
	cont := Panda.GetFavoriteSites()
	fmt.Printf("%+v\n", cont)

	str := ":rest-panda: お気に入りサイト情報 :rest-panda:\n"
	for i := 0; i < len(cont.FavoriteSitesIDs); i++ {
		dat := Panda.GetContent(cont.FavoriteSitesIDs[i])
		fmt.Printf("%+v\n", dat)
		for j := 0; j < len(dat); j++ {
			str += "> " + dat[j].Author + " " + dat[j].Title + " " + dat[j].EntityTitle + dat[j].FromDate + " " + dat[j].EndDate + " " + dat[j].Quota + "\n"
		}
	}

	// for i := 0; i < len(ass); i++ {
	// 	fmt.Printf("%+v\n", ass[i])
	// 	mes := ass[i].Author + "\n" + ass[i].Title
	// 	_, _, err := c.PostMessage(config.PostChannel, slack.MsgOptionText(mes, false))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// MsgOptionText() の第二引数に true を設定すると特殊文字をエスケープする
	_, _, err := c.PostMessage(config.PostChannel, slack.MsgOptionText(str, false))
	if err != nil {
		panic(err)
	}
}
