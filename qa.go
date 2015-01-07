package docomo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// QaURL docomoAPIの知識Q&Aのmethod
	QaURL = "/knowledgeQA/v1/ask"
)

// ResponseCode 知識Q&Aのレスポンスコード
type ResponseCode string

const (
	// OkResponseCode 質問結果あり
	OkResponseCode ResponseCode = "S020000"
	// NoResponseCode 質問結果なし
	NoResponseCode ResponseCode = "S020011"
)

// QARequest 知識Q&Aのリクエスト
type QARequest struct {
	QAText string `json:"q"`
}

// QAResponse 知識Q&Aのレスポンス
type QAResponse struct {
	Code    ResponseCode `json:"code"`
	Answers []struct {
		AnswerText string `json:"answerText"`
		LinkText   string `json:"linkText"`
		LinkURL    string `json:"linkUrl"`
		Rank       string `json:"rank"`
	} `json:"answers"`
	Message struct {
		TextForDisplay string `json:"textForDisplay"`
		TextForSpeech  string `json:"textForSpeech"`
	} `json:"message"`
}

// Success 質問成功
func (q QAResponse) Success() bool {
	return q.Code == OkResponseCode
}

// SendQA docomoAPIを呼び出して結果を返す
func (d *DocomoClient) SendQA(req QARequest) (*QAResponse, error) {

	val := url.Values{}
	val.Set("q", req.QAText)
	res, err := d.get(QaURL, val)
	if err != nil {
		return nil, err
	}

	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("response error : " + string(resData))
	}

	var q QAResponse
	if err := json.Unmarshal(resData, &q); err != nil {
		return nil, err
	}
	return &q, nil
}
