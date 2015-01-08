package docomo

import "net/url"

const (
	// TrendGenreURL docomoAPIのトレンドAPIのジャンル情報の取得
	TrendGenreURL = "/webCuration/v3/genre"
)

// TrendService API docs: https://dev.smt.docomo.ne.jp/?p=docs.api.page&api_docs_id=26
type TrendService struct {
	client *DocomoClient
}

// TrendGenreRequest ジャンル情報取得のリクエスト
type TrendGenreRequest struct {
	// ja: 日本語（default）、en: 英語
	Lang *string `json:"lang"`
}

// TrendGenreResponse ジャンル情報取得のレスポンス
type TrendGenreResponse struct {
	Genre []Genre `json:"genre"`
}

// Genre ジャンル情報
type Genre struct {
	Description string  `json:"description"`
	GenreID     float64 `json:"genreId"`
	Title       string  `json:"title"`
}

// GetGenre ジャンル情報の取得する.
func (t *TrendService) GetGenre(req TrendGenreRequest) (*TrendGenreResponse, error) {

	v := url.Values{}
	if req.Lang != nil {
		v.Set("lang", *req.Lang)
	}

	res, err := t.client.get(TrendGenreURL, v)
	if err != nil {
		return nil, err
	}

	var trendRes TrendGenreResponse
	if err := responseUnmarshal(res.Body, &trendRes); err != nil {
		return nil, err
	}

	return &trendRes, nil
}
