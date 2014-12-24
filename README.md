go-docomo
=========

[![GoDoc](https://godoc.org/github.com/kyokomi/go-docomo?status.svg)](https://godoc.org/github.com/kyokomi/go-docomo)

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

	log.Printf("%s\n", res.Utt)
}
```

実行結果

```
2014/12/24 11:52:11 ちわ
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

	log.Printf("%s\n", res.Answers[0].AnswerText)
}
```

実行結果

```
2014/12/24 11:52:11 1,785m
```


## License

[MIT](https://github.com/kyokomi/go-docomo/blob/master/LICENSE)
