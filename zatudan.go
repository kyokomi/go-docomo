package docomo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	DIALOGUE_URL = "/dialogue/v1/dialogue"
)

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

// ZatsudanResponse 雑談APIの結果
type ZatsudanResponse struct {
	Context string `json:"context"`
	Da      string `json:"da"`
	Mode    string `json:"mode"`
	Utt     string `json:"utt"`
	Yomi    string `json:"yomi"`
	// error時
	RequestError struct {
		PolicyException struct {
			MessageId string `json:"messageId"`
			Text      string `json:"text"`
		} `json:"policyException"`
	} `json:"requestError"`
}

// SendZatsudan 雑談APIに送信。Contextを更新する
func (d *DocomoClient) SendZatsudan(nickname, message string, refreshContext bool) (*ZatsudanResponse, error) {
	return d.sendZatsudan(DialogueRequest{
		Utt:      &message,
		Context:  &d.context,
		Nickname: &nickname,
	}, refreshContext)
}

func (d *DocomoClient) sendZatsudan(b DialogueRequest, refreshContext bool) (*ZatsudanResponse, error) {

	var data []byte
	var err error
	if data, err = json.Marshal(b); err != nil {
		return nil, err
	}

	res, err := d.post(DIALOGUE_URL, "application/json", bytes.NewReader(data))
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
