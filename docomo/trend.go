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
	// TrendSearchURL docomoAPIのトレンドAPIのキーワード検索 docs: https://dev.smt.docomo.ne.jp/?p=docs.api.page&api_docs_id=25#tag01
	TrendSearchURL = "/webCuration/v3/search"
	// TrendRelatedURL docomoAPIのトレンドAPIの関連記事取得 docs: https://dev.smt.docomo.ne.jp/?p=docs.api.page&api_docs_id=109#tag01
	TrendRelatedURL = "/webCuration/v3/relatedContents"
)

// TrendService API docs: https://dev.smt.docomo.ne.jp/?p=docs.api.page&api_docs_id=26
type TrendService struct {
	client *Client
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

// TrendSearchRequest キーワード検索リクエスト
type TrendSearchRequest struct {
	// ジャンルID
	GenreID *int `json:"genreId"`
	// キーワード（UTF-8）
	Keyword *string `json:"keyword"`
	// ja: 日本語（default）、en: 英語
	Lang *string `json:"lang"`
	// 記事一覧の開始番号を指定(1以上999999以下の整数)。デフォルトは1
	StartNo *int `json:"s"`
	// カテゴリ内の記事一覧の取得件数を指定(0以上50以下の整数)。デフォルトは1
	Num *int `json:"n"`
}

// TrendSearchResponse キーワード検索レスポンス
type TrendSearchResponse struct {
	ArticleContents []struct {
		ContentData struct {
			Body        string `json:"body"`
			CreatedDate string `json:"createdDate"`
			ImageSize   struct {
				Height int `json:"height"`
				Width  int `json:"width"`
			} `json:"imageSize"`
			ImageURL     string `json:"imageUrl"`
			LinkURL      string `json:"linkUrl"`
			SourceDomain string `json:"sourceDomain"`
			SourceName   string `json:"sourceName"`
			Title        string `json:"title"`
		} `json:"contentData"`
		ContentID   int `json:"contentId"`
		ContentType int `json:"contentType"`
		GenreID     int `json:"genreId"`
	} `json:"articleContents"`
	IssueDate    string `json:"issueDate"`
	ItemsPerPage int    `json:"itemsPerPage"`
	StartIndex   int    `json:"startIndex"`
	TotalResults int    `json:"totalResults"`
}

// TrendRelatedRequest 関連記事取得リクエスト
type TrendRelatedRequest struct {
	// 関連記事を取得する記事ID 必須
	ContentID *int `json:"contentId"`
}

// TrendRelatedResponse 関連記事取得レスポンス
type TrendRelatedResponse struct {
	ArticleContents []struct {
		ContentData struct {
			Body        string `json:"body"`
			CreatedDate string `json:"createdDate"`
			ImageSize   struct {
				Height int `json:"height"`
				Width  int `json:"width"`
			} `json:"imageSize"`
			ImageURL     string `json:"imageUrl"`
			LinkURL      string `json:"linkUrl"`
			SourceDomain string `json:"sourceDomain"`
			SourceName   string `json:"sourceName"`
			Title        string `json:"title"`
		} `json:"contentData"`
		ContentID       int    `json:"contentId"`
		ContentType     int    `json:"contentType"`
		GenreID         int    `json:"genreId"`
		RelatedContents string `json:"relatedContents"`
	} `json:"articleContents"`
	TotalResults int `json:"totalResults"`
}

// GetGenre ジャンル情報の取得する.
func (t *TrendService) GetGenre(req TrendGenreRequest) (*TrendGenreResponse, error) {

	v := url.Values{}
	if req.Lang != nil {
		v.Set("lang", *req.Lang)
	}

	var trendRes TrendGenreResponse
	_, err := t.client.get(TrendGenreURL, v, &trendRes)
	if err != nil {
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

	var trendRes TrendContentsResponse
	_, err := t.client.get(TrendContentsURL, v, &trendRes)
	if err != nil {
		return nil, err
	}

	return &trendRes, nil
}

// GetSearch キーワード検索
func (t *TrendService) GetSearch(req TrendSearchRequest) (*TrendSearchResponse, error) {

	v := url.Values{}
	if err := validation(req.Keyword); err != nil {
		return nil, err
	}
	v.Set("keyword", *req.Keyword)

	if req.GenreID != nil {
		v.Set("genreId", strconv.Itoa(*req.GenreID))
	}
	if req.Lang != nil {
		v.Set("lang", *req.Lang)
	}
	if req.StartNo != nil {
		v.Set("s", strconv.Itoa(*req.StartNo))
	}
	if req.Num != nil {
		v.Set("n", strconv.Itoa(*req.Num))
	}

	var trendRes TrendSearchResponse
	_, err := t.client.get(TrendSearchURL, v, &trendRes)
	if err != nil {
		return nil, err
	}

	return &trendRes, nil
}

// GetRelated 関連記事取得
func (t *TrendService) GetRelated(req TrendRelatedRequest) (*TrendRelatedResponse, error) {

	v := url.Values{}
	if err := validation(req.ContentID); err != nil {
		return nil, err
	}
	v.Set("contentId", strconv.Itoa(*req.ContentID))

	var trendRes TrendRelatedResponse
	_, err := t.client.get(TrendRelatedURL, v, &trendRes)
	if err != nil {
		return nil, err
	}

	return &trendRes, nil
}

func validation(param interface{}) error {
	if param == nil {
		return errors.New("genreId has not set")
	}

	return nil
}
