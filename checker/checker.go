package checker

import (
	"net/http"

	"github.com/moedev99/crawly/internal/linkextractor"
	"github.com/moedev99/crawly/shared/logger"
	"golang.org/x/net/html"
)

func GetStatusCode(url string) {
	res, err := http.Get(url)
	if err != nil {
		logger.Infof("Error while calling %q❌", url)
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		logger.Infof("url=%s statusCode=%d ✅", url, res.StatusCode)
	} else {
		logger.Infof("url=%s statusCode=%d ❌", url, res.StatusCode)
	}
}

func RemoveDuplicateUrls(anchors []*html.Node, baseUrl string) []string {
	var urls []string

	for _, anchor := range anchors {
		url := linkextractor.ExtractUrl(anchor.Attr, baseUrl)
		if url != "" {
			urls = append(urls, url)
		}
	}
	return urls
}
