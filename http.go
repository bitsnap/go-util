package goutil

import (
	"io"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func HttpGetDocument(url string) *goquery.Document {
	res, err := http.Get(url)
	if err != nil {
		Logger().Fatalln(err)
		return nil
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	if res.StatusCode != 200 {
		Logger().Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		Logger().Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return nil
	}

	return doc
}
