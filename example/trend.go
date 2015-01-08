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
	flag.StringVar(&apiKey, "APIKEY", os.Getenv("DOCOMO_APIKEY"), "docomo developerで登録したAPIKEYをして下さい")
	flag.Parse()
}

func main() {
	if apiKey == "" {
		log.Fatalln("APIKEYを指定して下さい")
	}

	d := docomo.New(apiKey)

	res, err := d.Trend.GetGenre(docomo.TrendGenreRequest{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)
}
