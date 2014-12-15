package docomo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
)

func (d *DocomoClient) SendZatsudan(nickname, message string) ([]byte, error) {

	// Mode dialog or srtr
	// CharactorID なし:デフォルト 20:関西弁 30:あかちゃん
	type DialogueBody struct {
		Utt            string `json:"utt"`
		Context        string `json:"context"`
		Nickname       string `json:"nickname"`
//		NicknameYomi   string `json:"nickname_y"`
//		Sex            string `json:"sex"`
//		Bloodtype      string `json:"bloodtype"`
//		BirthdateY     int    `json:"birthdateY"`
//		BirthdateM     int    `json:"birthdateM"`
//		BirthdateD     int    `json:"birthdateD"`
//		Age            int    `json:"age"`
//		Constellations string `json:"constellations"`
//		Place          string `json:"place"`
//		Mode           string `json:"mode"`
//		CharactorID    int    `json:"t"`
	}

	b := DialogueBody{
		Utt:     message,
		Context: d.context,
		Nickname: nickname,
	}

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

	if res.StatusCode == http.StatusOK {
		var resMap map[string]string
		if err := json.Unmarshal(resData, &resMap); err != nil {
			return nil, err
		}
		d.context = resMap["context"]
	} else {
		fmt.Println(string(resData))
	}

	return resData, nil
}
