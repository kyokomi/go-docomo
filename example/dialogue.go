package main

import (
	"flag"
	"log"

	"os"

	docomo "github.com/kyokomi/go-docomo"
)

var nickName, message, apiKey string

func init() {
	flag.StringVar(&nickName, "nickName", "foo", "ニックネーム")
	flag.StringVar(&message, "message", "こんにちわ", "雑談のメッセージ")
	flag.StringVar(&apiKey, "APIKEY", os.Getenv("DOCOMO_APIKEY"), "docomo developerで登録したAPIKEYをして下さい")
	flag.Parse()
}

func main() {
	if apiKey == "" {
		log.Fatalln("APIKEYを指定して下さい")
	}

	d := docomo.New(apiKey)

	zatsu := docomo.DialogueRequest{
		Nickname: &nickName,
		Utt:      &message,
	}
	res, err := d.Dialogue.Get(zatsu, true)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%s\n", res.Utt)
}
