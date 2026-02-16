package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/moedev99/crawly/crawly"
	"github.com/moedev99/crawly/shared/logger"
)

func main() {
	if len(os.Args) == 1 {
		logger.Error(errors.New("Please pass url to check, or run -help"))
	}

	arg := os.Args[1]
	if arg == "-help" {
		fmt.Printf("Usage: crawly [OPTIONS]\nOptions:\n  crawly url \n")
		return
	}

	crawly.Crawl(arg)

}
