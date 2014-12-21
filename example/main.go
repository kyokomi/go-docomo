package main

import (
	"flag"
	"log"

	docomo "github.com/kyokomi/go-docomo"
)

func main() {
	var nickName, message, apiKey string
	flag.StringVar(&nickName, "nickName", "foo", "ニックネーム")
	flag.StringVar(&message, "message", "こんにちわ", "雑談のメッセージ")
	flag.StringVar(&apiKey, "APIKEY", "", "docomo developerで登録したAPIKEYをして下さい")
	flag.Parse()

	if apiKey == "" {
		log.Fatalln("APIKEYを指定して下さい")
	}

	d := docomo.New(apiKey)

	if res, err := d.SendZatsudan(nickName, message); err != nil {
		log.Fatalln(err)
	} else {
		log.Println(string(res))
	}

	if res, err := d.SendQA(&docomo.QARequest{
		QAText: "プラネテューヌの女神は誰",
	}); err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("%+v\n", res)
	}
}
