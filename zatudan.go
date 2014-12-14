package docomo

import (
	"net/http"
	"os"
	"encoding/json"
	"log"
	"bytes"
	"io/ioutil"
	"fmt"
)

const DIALOGUE_URL = "dialogue/v1/dialogue"

type Dialogue struct {
	client *http.Client
	domain string
	apiKey string
	context string
}

func NewDialogue() *Dialogue {
	// TODO: Commandの引数にする
	apiKey := os.Getenv("DOCOMO_APIKEY")

	return &Dialogue{
		client: http.DefaultClient,
		domain: "https://api.apigw.smt.docomo.ne.jp",
		apiKey: apiKey,
		context: "",
	}
}

func (d *Dialogue) Send(message string) []byte {

	var client = http.DefaultClient

	apiKey := os.Getenv("DOCOMO_APIKEY")

	url := "https://api.apigw.smt.docomo.ne.jp/dialogue/v1/dialogue?APIKEY=" + apiKey

	// TODO: あとで増やす
	type DialogueBody struct {
		Utt     string `json:"utt"`
		Context string `json:"context"`
	}
	b := DialogueBody{
		Utt: message,
		Context: d.context,
	}

	var data []byte
	var err error
	if data, err = json.Marshal(b); err != nil {
		log.Fatalln(err)
	}

	res, err := client.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		log.Fatalln(err)
	}

	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var resMap map[string]string
	if err := json.Unmarshal(resData, &resMap); err != nil {
		log.Fatalln(err)
	}
	d.context = resMap["context"]
	fmt.Println("context = ", d.context)

	return resData
}
