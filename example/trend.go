package main

import (
	"flag"
	"log"

	"os"

	"fmt"

	docomo "github.com/kyokomi/go-docomo"
)

var apiKey, keyword string

func init() {
	flag.StringVar(&apiKey, "APIKEY", os.Getenv("DOCOMO_APIKEY"), "docomo developerで登録したAPIKEY")
	flag.StringVar(&keyword, "keyword", "", "記事検索キーワード")
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

	var searchReq docomo.TrendSearchRequest
	searchReq.Keyword = &keyword

	sRes, err := d.Trend.GetSearch(searchReq)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(sRes)
}