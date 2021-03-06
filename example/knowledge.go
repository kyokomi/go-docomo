package main

import (
	"flag"
	"log"

	"os"

	"github.com/kyokomi/go-docomo/docomo"
)

var apiKey, qa string

func init() {
	flag.StringVar(&apiKey, "APIKEY", os.Getenv("DOCOMO_APIKEY"), "docomo developerで登録したAPIKEY")
	flag.StringVar(&qa, "qa", "三つ峠の標高は？", "質問内容を指定してください")
	flag.Parse()
}

func main() {
	if apiKey == "" {
		log.Fatalln("APIKEYを指定して下さい")
	}

	d := docomo.NewClient(apiKey)

	qaReq := docomo.KnowledgeQARequest{
		QAText: qa,
	}

	res, err := d.KnowledgeQA.Get(qaReq)
	if err != nil {
		log.Fatalln(err)
	} else if !res.Success() {
		log.Fatalln("質問の答えがわかりません")
	}

	log.Printf("%s\n", res.Answers[0].AnswerText)
}
