package docomo

import (
	"net/http"
	"io"
)

const (
	DOMAIN_URL = "https://api.apigw.smt.docomo.ne.jp"
	// docomoURL
	DIALOGUE_URL = "/dialogue/v1/dialogue"
	IMAGE_URL = "/characterRecognition/v1/document"
)

type DocomoClient struct {
	client  *http.Client
	domain  string
	apiKey  string
	context string
}

func New(apiKey string) *DocomoClient {

	return &DocomoClient{
		client:  http.DefaultClient,
		domain:  DOMAIN_URL,
		apiKey:  apiKey,
		context: "",
	}
}

func (d *DocomoClient) createURL(docomoURL string) string {
	return d.domain + docomoURL + "?APIKEY=" + d.apiKey
}

func (d *DocomoClient) post(docomoURL string, bodyType string, body io.Reader) (resp *http.Response, err error) {
	return d.client.Post(d.createURL(docomoURL), bodyType, body)
}
