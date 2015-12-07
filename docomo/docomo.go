package docomo

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/kyokomi/go-docomo/internal"
	"golang.org/x/net/context"
)

const (
	// DomainURL DocomoAPIのhost
	DomainURL = "https://api.apigw.smt.docomo.ne.jp"
)

// Client DocomoAPIへのpostやgetを行うクライアント
type Client struct {
	client  *http.Client
	domain  string
	apiKey  string
	context string

	Trend       *TrendService
	KnowledgeQA *KnowledgeQAService
	Dialogue    *DialogueService
}

// NewClient docomo APIのClientを生成する
func NewClient(apiKey string) *Client {
	c := &Client{}
	cli, err := internal.ContextClient(context.Background())
	if err != nil {
		cli = http.DefaultClient
	}
	c.client = cli
	c.domain = DomainURL
	c.apiKey = apiKey
	c.context = ""

	c.Trend = &TrendService{client: c}
	c.KnowledgeQA = &KnowledgeQAService{client: c}
	c.Dialogue = &DialogueService{client: c}

	return c
}

func (c *Client) createURL(docomoURL string) string {
	return c.domain + docomoURL + "?APIKEY=" + c.apiKey
}

func (c *Client) post(docomoURL string, bodyType string, body io.Reader, v interface{}) (resp *http.Response, err error) {
	res, err := c.client.Post(c.createURL(docomoURL), bodyType, body)
	if err != nil {
		return nil, err
	}

	if err := responseUnmarshal(res.Body, v); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) get(docomoURL string, query url.Values, v interface{}) (resp *http.Response, err error) {

	path := c.createURL(docomoURL)
	for key, value := range query {
		path += "&" + key + "=" + url.QueryEscape(value[0])
	}

	u := url.URL{Path: path}

	res, err := c.client.Get(u.String())
	if err != nil {
		return nil, err
	}

	if err := responseUnmarshal(res.Body, v); err != nil {
		return nil, err
	}

	return res, nil
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
