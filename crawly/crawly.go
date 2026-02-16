package crawly

import (
	"net/http"
	"net/url"
	"sync"

	"github.com/moedev/crawly/checker"
	"github.com/moedev/crawly/internal/htmltraverser"
	"github.com/moedev/crawly/shared/logger"
	"golang.org/x/net/html"
)

func Crawl(arg string) {
	if _, err := url.ParseRequestURI(arg); err != nil {
		logger.Error(err)
	}

	res, err := http.Get(arg)
	if err != nil {
		logger.Error(err)
	}
	defer res.Body.Close()
	doc, err := html.Parse(res.Body)
	if err != nil {
		logger.Error(err)
	}

	var anchors []*html.Node
	htmltraverser.GetAnchorsDFS(doc, &anchors)

	var wg sync.WaitGroup

	urls := checker.RemoveDuplicateUrls(anchors, arg)
	for _, url := range urls {
		wg.Go(func() {
			checker.GetStatusCode(url)
		})
	}

	wg.Wait()
}
