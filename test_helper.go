package docomo

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
)

// Stub test用のスタブ
func Stub(filename string) (*httptest.Server, *Client) {
	stub, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(stub))
	}))
	c := NewClient("")
	c.SetDomain(ts.URL)
	return ts, c
}
