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
	flag.StringVar(&apiKey, "APIKEY", os.Getenv("DOCOMO_APIKEY"), "docomo developerで登録したAPIKEY")
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
	flag.StringVar(&apiKey, "APIKEY", os.Getenv("DOCOMO_APIKEY"), "docomo developerで登録したAPIKEY")
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

### Trend API

https://dev.smt.docomo.ne.jp/?p=docs.api.page&api_docs_id=26

[example/trend.go](https://github.com/kyokomi/go-docomo/blob/master/example/trend.go)

```go
package main

import (
	"flag"
	"log"

	"os"

	"fmt"

	docomo "github.com/kyokomi/go-docomo"
)

var apiKey string

func init() {
	flag.StringVar(&apiKey, "APIKEY", os.Getenv("DOCOMO_APIKEY"), "docomo developerで登録したAPIKEY")
	flag.Parse()
}

func main() {
	if apiKey == "" {
		log.Fatalln("APIKEYを指定して下さい")
	}

	d := docomo.New(apiKey)

	// ジャンル取得

	gRes, err := d.Trend.GetGenre(docomo.TrendGenreRequest{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(gRes)

	// 記事取得（ジャンルID指定）

	var contentsReq docomo.TrendContentsRequest
	contentsReq.GenreID = &gRes.Genre[0].GenreID

	cRes, err := d.Trend.GetContents(contentsReq)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cRes)

	// TODO: キーワード検索

}
```

## License

[MIT](https://github.com/kyokomi/go-docomo/blob/master/LICENSE)
