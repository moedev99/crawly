package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"

	"github.com/moedev/crawly/internal/checker"
	"github.com/moedev/crawly/internal/htmltraverser"
	"github.com/moedev/crawly/shared/logger"
	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please pass url to check, or run -help")
		return
	}

	arg := os.Args[1]
	if arg == "-help" {
		fmt.Println("Usage: crawly [OPTIONS]\nOptions:\n  crawly url \n")
		return
	}
	if _, err := url.ParseRequestURI(arg); err != nil {
		logger.Error(err)
	}

	res, err := http.Get(arg)
	if err != nil {
		logger.Error(err)
		return
	}
	defer res.Body.Close()

	doc, err := html.Parse(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var anchors []*html.Node
	htmltraverser.GetAnchorsDFS(doc, &anchors)

	var wg sync.WaitGroup

	urls := checker.RemoveDuplicateUrls(anchors, arg)
	wg.Add(len(urls))
	for _, url := range urls {
		go func(urls []string) {
			checker.GetStatusCode(url)
			defer wg.Done()
		}(urls)
	}

	wg.Wait()

}
