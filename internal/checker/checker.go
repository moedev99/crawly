package checker

import (
	"net/http"

	"github.com/moedev/crawly/shared/logger"
	"golang.org/x/net/html"
)

var uniqueUrls = make(map[string]bool)

func GetStatusCode(url string) {
	res, err := http.Get(url)
	if err != nil {
		logger.Infof("Error while calling %q❌", url)
		return
	}
	defer res.Body.Close()

	logger.Infof("url=%s statusCode=%d ✅", url, res.StatusCode)
}

func RemoveDuplicateUrls(anchors []*html.Node, baseUrl string) []string {
	var urls []string

	for _, anchor := range anchors {
		url := extractUrl(anchor.Attr, baseUrl)
		if url != "" {
			urls = append(urls, url)
		}
	}
	return urls
}
