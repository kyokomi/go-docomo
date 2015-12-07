package docomo

import (
	"bytes"
	"encoding/json"

	"golang.org/x/net/context"
)

const (
	// DialogueURL docomoAPIの雑談APIのmethod
	DialogueURL = "/dialogue/v1/dialogue"
)

// DialogueService API docs: https://dev.smt.docomo.ne.jp/?p=docs.api.page&api_docs_id=5
type DialogueService struct {
	ctx    context.Context
	client *Client
}

// DialogueRequest 雑談APIのリクエスト
// Mode dialog or srtr
// CharactorID なし:デフォルト 20:関西弁 30:あかちゃん
type DialogueRequest struct {
	Utt            *string `json:"utt"`
	Context        *string `json:"context"`
	Nickname       *string `json:"nickname"`
	NicknameYomi   *string `json:"nickname_y"`
	Sex            *string `json:"sex"`
	Bloodtype      *string `json:"bloodtype"`
	BirthdateY     *int    `json:"birthdateY"`
	BirthdateM     *int    `json:"birthdateM"`
	BirthdateD     *int    `json:"birthdateD"`
	Age            *int    `json:"age"`
	Constellations *string `json:"constellations"`
	Place          *string `json:"place"`
	Mode           *string `json:"mode"`
	CharactorID    *int    `json:"t"`
}

// DialogueResponse 雑談APIのレスポンス
type DialogueResponse struct {
	Context string `json:"context"`
	Da      string `json:"da"`
	Mode    string `json:"mode"`
	Utt     string `json:"utt"`
	Yomi    string `json:"yomi"`
	// error時
	RequestError struct {
		PolicyException struct {
			MessageID string `json:"messageId"`
			Text      string `json:"text"`
		} `json:"policyException"`
	} `json:"requestError"`
}

func (d *DialogueService) WithContext(ctx context.Context) *DialogueService {
	d.ctx = ctx
	return d
}

// Get 雑談APIを呼び出して結果を取得する.
// refreshContextがtrueの場合、DocomoClientのContextを更新する
func (d *DialogueService) Get(req DialogueRequest, refreshContext bool) (*DialogueResponse, error) {
	// context有効の場合、clientで保持しているcontextを設定する
	if refreshContext && req.Context != nil {
		d.SetContext(*req.Context)
	}

	var data []byte
	var err error
	if data, err = json.Marshal(req); err != nil {
		return nil, err
	}

	var dialogueRes DialogueResponse
	res, err := d.client.post(d.ctx, DialogueURL, "application/json", bytes.NewReader(data), &dialogueRes)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if refreshContext {
		d.SetContext(dialogueRes.Context)
	}

	return &dialogueRes, nil
}

// SetContext setter DocomoClient context.
func (d *DialogueService) SetContext(context string) {
	d.client.context = context
}
