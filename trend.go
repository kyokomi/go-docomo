package docomo

import (
	"errors"
	"net/url"
	"strconv"
)

const (
	// TrendGenreURL docomoAPIのトレンドAPIのジャンル情報の取得
	TrendGenreURL = "/webCuration/v3/genre"
	// TrendContentsURL docomoAPIのトレンドAPIの記事取得 docs: https://dev.smt.docomo.ne.jp/?p=docs.api.page&api_docs_id=24#tag01
	TrendContentsURL = "/webCuration/v3/contents"
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
	Description string `json:"description"`
	GenreID     int    `json:"genreId"`
	Title       string `json:"title"`
}

// TrendContentsRequest 記事取得リクエスト
type TrendContentsRequest struct {
	// 必須
	GenreID *int `json:"genreId"`
	// ja: 日本語（default）、en: 英語
	Lang *string `json:"lang"`
	// 記事一覧の開始番号を指定(1以上999999以下の整数)。デフォルトは1
	StartNo *int `json:"s"`
	// カテゴリ内の記事一覧の取得件数を指定(0以上50以下の整数)。デフォルトは1
	Num *int `json:"n"`
}

// TrendContentsResponse 記事取得レスポンス
type TrendContentsResponse struct {
	ArticleContents []struct {
		ContentData struct {
			Body        string `json:"body"`
			CreatedDate string `json:"createdDate"`
			ImageSize   struct {
				Height float64 `json:"height"`
				Width  float64 `json:"width"`
			} `json:"imageSize"`
			ImageURL     string `json:"imageUrl"`
			LinkURL      string `json:"linkUrl"`
			SourceDomain string `json:"sourceDomain"`
			SourceName   string `json:"sourceName"`
			Title        string `json:"title"`
		} `json:"contentData"`
		ContentID   float64 `json:"contentId"`
		ContentType float64 `json:"contentType"`
		GenreID     float64 `json:"genreId"`
	} `json:"articleContents"`
	IssueDate    string  `json:"issueDate"`
	ItemsPerPage float64 `json:"itemsPerPage"`
	StartIndex   float64 `json:"startIndex"`
	TotalResults float64 `json:"totalResults"`
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

// GetContents 指定したジャンルの記事を取得する.
func (t *TrendService) GetContents(req TrendContentsRequest) (*TrendContentsResponse, error) {

	v := url.Values{}
	if req.GenreID == nil {
		// ジャンル指定は必須
		return nil, errors.New("genreId has not set")
	}
	v.Set("genreId", strconv.Itoa(*req.GenreID))

	if req.Lang != nil {
		v.Set("lang", *req.Lang)
	}
	if req.StartNo != nil {
		v.Set("s", strconv.Itoa(*req.StartNo))
	}
	if req.Num != nil {
		v.Set("n", strconv.Itoa(*req.Num))
	}

	res, err := t.client.get(TrendContentsURL, v)
	if err != nil {
		return nil, err
	}

	var trendRes TrendContentsResponse
	if err := responseUnmarshal(res.Body, &trendRes); err != nil {
		return nil, err
	}

	return &trendRes, nil
}
