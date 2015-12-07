package docomo

import (
"net/url"
"golang.org/x/net/context"
)

const (
	// KnowledgeQAURL docomoAPIの知識Q&Aのmethod
	KnowledgeQAURL = "/knowledgeQA/v1/ask"
)

// ResponseCode 知識Q&Aのレスポンスコード
type ResponseCode string

const (
	// OkResponseCode 質問結果あり
	OkResponseCode ResponseCode = "S020000"
	// NoResponseCode 質問結果なし
	NoResponseCode ResponseCode = "S020011"
)

// KnowledgeQAService API docs: https://dev.smt.docomo.ne.jp/?p=docs.api.page&api_docs_id=6
type KnowledgeQAService struct {
	ctx context.Context
	client *Client
}

// KnowledgeQARequest 知識Q&Aのリクエスト
type KnowledgeQARequest struct {
	QAText string `json:"q"`
}

// KnowledgeQAResponse 知識Q&Aのレスポンス
type KnowledgeQAResponse struct {
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

func (d *KnowledgeQAService) WithContext(ctx context.Context) *KnowledgeQAService {
	d.ctx = ctx
	return d
}

// Success 質問成功
func (q KnowledgeQAResponse) Success() bool {
	return q.Code == OkResponseCode
}

// Get docomoAPIを呼び出して結果を返す
func (q *KnowledgeQAService) Get(req KnowledgeQARequest) (*KnowledgeQAResponse, error) {

	val := url.Values{}
	val.Set("q", req.QAText)

	var knowRes KnowledgeQAResponse
	res, err := q.client.get(q.ctx, KnowledgeQAURL, val, &knowRes)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return &knowRes, nil
}
