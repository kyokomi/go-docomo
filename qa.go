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

// QARequest 知識Q&Aのリクエスト
type QARequest struct {
	QAText string `json:"q"`
}

// QAResponse 知識Q&Aのレスポンス
type QAResponse struct {
	Code    string `json:"code"`
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

// SendQA docomoAPIを呼び出して結果を返す
func (d *DocomoClient) SendQA(req *QARequest) (*QAResponse, error) {

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
