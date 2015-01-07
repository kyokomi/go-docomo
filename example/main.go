package main

import (
	"flag"
	"log"

	docomo "github.com/kyokomi/go-docomo"
)

var nickName, message, apiKey, qa string

func init() {
	flag.StringVar(&nickName, "nickName", "foo", "ニックネーム")
	flag.StringVar(&message, "message", "こんにちわ", "雑談のメッセージ")
	flag.StringVar(&apiKey, "APIKEY", "", "docomo developerで登録したAPIKEYをして下さい")
	flag.StringVar(&qa, "qa", "三つ峠の標高は？", "質問内容を指定してください")
	flag.Parse()
}

func main() {
	if apiKey == "" {
		log.Fatalln("APIKEYを指定して下さい")
	}

	d := docomo.New(apiKey)

	// 雑談API
	zatsu := docomo.DialogueRequest{
		Nickname: &nickName,
		Utt:      &message,
	}
	if res, err := d.SendZatsudan(zatsu, true); err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("%s\n", res.Utt)
	}

	// 知識Q&A
	qaReq := docomo.QARequest{
		QAText: qa,
	}
	if res, err := d.SendQA(qaReq); err != nil {
		log.Fatalln(err)
	} else {
		if !res.Success() {
			log.Println("質問の答えがわかりません")
		} else {
			log.Printf("%s\n", res.Answers[0].AnswerText)
		}
	}
}
