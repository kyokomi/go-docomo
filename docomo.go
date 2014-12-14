package docomo

import (
	"net/http"
	"os"
	"io"
	"mime/multipart"
	"bytes"
	"fmt"
)

const DOMAIN_URL = "https://api.apigw.smt.docomo.ne.jp"

// docomoURL
const DIALOGUE_URL = "/dialogue/v1/dialogue"
const IMAGE_URL = "/characterRecognition/v1/document"

type DocomoClient struct {
	client  *http.Client
	domain  string
	apiKey  string
	context string
}

func New() *DocomoClient {
	// TODO: Commandの引数にする
	apiKey := os.Getenv("DOCOMO_APIKEY")

	return &DocomoClient{
		client:  http.DefaultClient,
		domain:  DOMAIN_URL,
		apiKey:  apiKey,
		context: "",
	}
}

func (d *DocomoClient) PostJSON(docomoURL string, body io.Reader) (resp *http.Response, err error) {
	return d.client.Post(d.domain + docomoURL + "?APIKEY=" + d.apiKey, "application/json", body)
}

func (d *DocomoClient) PostImage(file string) (resp *http.Response, err error) {

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	// Add your image file
	f, err := os.Open(file)
	if err != nil {
		return
	}
	fw, err := w.CreateFormFile("image", file)
	if err != nil {
		return
	}
	if _, err = io.Copy(fw, f); err != nil {
		return
	}

	// Add the other fields
	if fw, err = w.CreateFormField("lang"); err != nil {
		return
	}
	if _, err = fw.Write([]byte("eng")); err != nil {
		return
	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", d.domain + IMAGE_URL + "?APIKEY=" + d.apiKey, &b)
	if err != nil {
		return
	}

	fmt.Println("ContentLength = ", req.ContentLength)

	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}

	return res, nil
}
