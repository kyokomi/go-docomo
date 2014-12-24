package docomo

import (
	"io"
	"net/http"
	"net/url"
)

const (
	DOMAIN_URL = "https://api.apigw.smt.docomo.ne.jp"
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

func (d *DocomoClient) get(docomoURL string, query url.Values) (resp *http.Response, err error) {

	u := d.createURL(docomoURL)
	for key, value := range query {
		u += "&" + key + "=" + url.QueryEscape(value[0])
	}
	return d.client.Get(u)
}
