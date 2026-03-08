package main

import (
	"net/http"
	"os"
	"time"

	"github.com/moedev99/crawly/checker"
	"github.com/moedev99/crawly/crawly"
)

func main() {
	out := os.Stdout
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	os.Exit(crawly.Main(os.Args[1:], out, checker.NewChecker(out, client)))
}
