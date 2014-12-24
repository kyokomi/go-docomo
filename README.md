go-docomo
=========

go-docomo is a Go client library for accessing the [docomo API](https://dev.smt.docomo.ne.jp/).

## Usage

### 雑談 API

```go
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
	flag.StringVar(&apiKey, "APIKEY", "", "docomo developerで登録したAPIKEYを指定して下さい")
	flag.Parse()

	if apiKey == "" {
		log.Fatalln("APIKEYを指定して下さい")
	}

	d := docomo.New(apiKey)

	// 雑談API
	zatsu := docomo.DialogueRequest{
        Nickname: &nickName,
        Utt: &message,
    }
	if res, err := d.SendZatsudan(&zatsu, true); err != nil {
		log.Fatalln(err)
	}

	log.Println("%s", res.Utt)
}
```

### 知識Q&A API

```go
package main

import (
	"flag"
	"log"

	docomo "github.com/kyokomi/go-docomo"
)

func main() {
	var text, apiKey string
	flag.StringVar(&text, "qa-text", "三つ峠の標高は？", "質問内容")
	flag.StringVar(&apiKey, "APIKEY", "", "docomo developerで登録したAPIKEYを指定して下さい")
	flag.Parse()

	if apiKey == "" {
		log.Fatalln("APIKEYを指定して下さい")
	}

	d := docomo.New(apiKey)

    // 知識Q&A API
    qa := docomo.QARequest{
        QAText: qa-text,
    }
    if res, err := d.SendQA(&qa); err != nil {
		log.Fatalln(err)
	}

	log.Println("%s", res.Answers[0].AnswerText)
}
```

## License

[MIT](https://github.com/kyokomi/go-docomo/blob/master/LICENSE)
