package docomo

import (
	"testing"
	"reflect"
)

func TestTrendGetGenre(t *testing.T) {
	type TestCase struct {
		in  string
		out TrendGenreResponse
	}

	testCase := TestCase{
		in:  "../tests/stubs/trend_genre.json",
	}
	serve, client := Stub(testCase.in, &testCase.out)
	defer serve.Close()

	req := TrendGenreRequest{}
	res, err := client.Trend.GetGenre(req)
	if err != nil {
		t.Errorf("error Request %s\n", err)
	}

	if !reflect.DeepEqual(*res, testCase.out) {
		t.Errorf("error Response %s != %s\n", res, testCase.out)
	}
}

func TestTrendGetContents(t *testing.T) {
	type TestCase struct {
		in  string
		out TrendContentsResponse
	}

	testCase := TestCase{
		in:  "../tests/stubs/trend_contents.json",
	}
	serve, client := Stub(testCase.in, &testCase.out)
	defer serve.Close()

	req := TrendContentsRequest{}
	var genreID int = 1
	req.GenreID = &genreID
	res, err := client.Trend.GetContents(req)
	if err != nil {
		t.Errorf("error Request %s\n", err)
	}

	if !reflect.DeepEqual(*res, testCase.out) {
		t.Errorf("error Response %s != %s\n", res, testCase.out)
	}
}

func TestTrendGetSearch(t *testing.T) {
	type TestCase struct {
		in  string
		out TrendSearchResponse
	}

	testCase := TestCase{
		in:  "../tests/stubs/trend_search.json",
	}
	serve, client := Stub(testCase.in, &testCase.out)
	defer serve.Close()

	req := TrendSearchRequest{}
	var keyword string = "test"
	req.Keyword = &keyword
	res, err := client.Trend.GetSearch(req)
	if err != nil {
		t.Errorf("error Request %s\n", err)
	}

	if !reflect.DeepEqual(*res, testCase.out) {
		t.Errorf("error Response %s != %s\n", res, testCase.out)
	}
}

func TestTrendGetRelated(t *testing.T) {
	type TestCase struct {
		in  string
		out TrendRelatedResponse
	}

	testCase := TestCase{
		in:  "../tests/stubs/trend_related.json",
	}
	serve, client := Stub(testCase.in, &testCase.out)
	defer serve.Close()

	req := TrendRelatedRequest{}
	var contentID int = 1
	req.ContentID = &contentID
	res, err := client.Trend.GetRelated(req)
	if err != nil {
		t.Errorf("error Request %s\n", err)
	}

	if !reflect.DeepEqual(*res, testCase.out) {
		t.Errorf("error Response %s != %s\n", res, testCase.out)
	}
}
