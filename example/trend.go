package main

import (
	"flag"
	"log"

	"os"

	"fmt"

	"github.com/kyokomi/go-docomo/docomo"
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

	d := docomo.NewClient(apiKey)

	fmt.Println("\n---ジャンル取得---\n")

	gRes, err := d.Trend.GetGenre(docomo.TrendGenreRequest{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(gRes)

	fmt.Println("\n---記事取得---\n")

	var contentsReq docomo.TrendContentsRequest
	contentsReq.GenreID = &gRes.Genre[0].GenreID

	cRes, err := d.Trend.GetContents(contentsReq)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cRes)

	fmt.Println("\n---キーワード検索---\n")

	var searchReq docomo.TrendSearchRequest
	searchReq.Keyword = &keyword

	sRes, err := d.Trend.GetSearch(searchReq)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(sRes)

	fmt.Println("\n---関連記事検索---\n")

	var relatedReq docomo.TrendRelatedRequest

	// TODO: なかなかヒットしないからループして全部チェックしてる
	for _, cont := range sRes.ArticleContents {
		relatedReq.ContentID = &cont.ContentID
		rRes, err := d.Trend.GetRelated(relatedReq)
		if err != nil {
			log.Fatalln(err)
		}

		if rRes.TotalResults > 0 {
			fmt.Println(rRes)
			break
		}
	}
}
