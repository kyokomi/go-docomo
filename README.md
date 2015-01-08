go-docomo
=========

[![GoDoc](https://godoc.org/github.com/kyokomi/go-docomo?status.svg)](https://godoc.org/github.com/kyokomi/go-docomo)

go-docomo is a Go client library for accessing the [docomo API](https://dev.smt.docomo.ne.jp/).

## Usage

### dialogue API

https://dev.smt.docomo.ne.jp/?p=docs.api.page&api_docs_id=5

[example/dialogue.go](https://github.com/kyokomi/go-docomo/blob/master/example/dialogue.go)

```go
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
```

OutPut:

```
2014/12/24 11:52:11 ちわ
```

### KnowledgeQA API

https://dev.smt.docomo.ne.jp/?p=docs.api.page&api_docs_id=6

[example/knowledge.go](https://github.com/kyokomi/go-docomo/blob/master/example/knowledge.go)

```go
package main

import (
	"flag"
	"log"

	"os"

	docomo "github.com/kyokomi/go-docomo"
)

var apiKey, qa string

func init() {
	flag.StringVar(&apiKey, "APIKEY", os.Getenv("DOCOMO_APIKEY"), "docomo developerで登録したAPIKEYをして下さい")
	flag.StringVar(&qa, "qa", "三つ峠の標高は？", "質問内容を指定してください")
	flag.Parse()
}

func main() {
	if apiKey == "" {
		log.Fatalln("APIKEYを指定して下さい")
	}

	d := docomo.New(apiKey)

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
```

OutPut:

```
2014/12/24 11:52:11 1,785m
```

## License

[MIT](https://github.com/kyokomi/go-docomo/blob/master/LICENSE)
