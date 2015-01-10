package docomo

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// DomainURL DocomoAPIのhost
	DomainURL = "https://api.apigw.smt.docomo.ne.jp"
)

// DocomoClient DocomoAPIへのpostやgetを行うクライアント
type DocomoClient struct {
	client  *http.Client
	domain  string
	apiKey  string
	context string

	Trend       *TrendService
	KnowledgeQA *KnowledgeQAService
	Dialogue    *DialogueService
}

// New DocomoClientを生成する
func New(apiKey string) *DocomoClient {
	c := &DocomoClient{}
	c.client = http.DefaultClient
	c.domain = DomainURL
	c.apiKey = apiKey
	c.context = ""

	c.Trend = &TrendService{client: c}
	c.KnowledgeQA = &KnowledgeQAService{client: c}
	c.Dialogue = &DialogueService{client: c}

	return c
}

func (d *DocomoClient) createURL(docomoURL string) string {
	return d.domain + docomoURL + "?APIKEY=" + d.apiKey
}

func (d *DocomoClient) post(docomoURL string, bodyType string, body io.Reader) (resp *http.Response, err error) {
	return d.client.Post(d.createURL(docomoURL), bodyType, body)
}

func (d *DocomoClient) get(docomoURL string, query url.Values) (resp *http.Response, err error) {

	path := d.createURL(docomoURL)
	for key, value := range query {
		path += "&" + key + "=" + url.QueryEscape(value[0])
	}

	u := url.URL{Path: path}
	return d.client.Get(u.String())
}

func responseUnmarshal(body io.ReadCloser, v interface{}) error {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}
