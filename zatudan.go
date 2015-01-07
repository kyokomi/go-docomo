package docomo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	// DialogueURL docomoAPIの雑談APIのmethod
	DialogueURL = "/dialogue/v1/dialogue"
)

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

// ZatsudanResponse 雑談APIのレスポンス
type ZatsudanResponse struct {
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

// SendZatsudan 雑談APIに送信
// refreshContextがtrueの場合、DocomoClientのContextを更新する
func (d *DocomoClient) SendZatsudan(req DialogueRequest, refreshContext bool) (*ZatsudanResponse, error) {

	// context有効の場合、clientで保持しているcontextを設定する
	if refreshContext && req.Context != nil {
		d.context = *req.Context
	}

	var data []byte
	var err error
	if data, err = json.Marshal(req); err != nil {
		return nil, err
	}

	res, err := d.post(DialogueURL, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var zatsudanRes ZatsudanResponse
	if err := json.Unmarshal(resData, &zatsudanRes); err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusOK {
		if refreshContext {
			d.context = zatsudanRes.Context
		}
	} else {
		return nil, errors.New("zatsudan error response: " + zatsudanRes.RequestError.PolicyException.Text)
	}

	return &zatsudanRes, nil
}
