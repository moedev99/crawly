package checker

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/fatih/color"
)

var DEAD, OKAY = color.RedString("DEAD"), color.GreenString("OKAY")

type Result struct {
	Link   string
	Status string
}

var notAllowedscheme = map[string]bool{"mailto": true}

type Checker struct {
	BaseUrl    *url.URL
	Output     io.Writer
	HttpClient *http.Client
	Visited    map[string]bool
	Results    []Result
}

func NewChecker(out io.Writer, client *http.Client) *Checker {
	return &Checker{
		Output:     out,
		HttpClient: client,
		Visited:    map[string]bool{},
		Results:    make([]Result, 0),
	}
}

// AddResutl will write to the output if there is an error otherwise it will print the successful hit later.
func (c *Checker) AddResult(link string, status string, isError bool) {

	if isError {
		fmt.Fprintf(c.Output, "%s --> %q \n", status, link)
	}

	c.Results = append(c.Results, Result{link, status})

}
func (c *Checker) Check(site string) {
	base, err := url.Parse(site)
	if err != nil {
		c.AddResult(site, DEAD, true)
		return
	}
	if !strings.HasSuffix(site, "/") {
		site += "/"
	}
	c.Visited[site] = true
	c.BaseUrl = base
	c.Crawl(base)
}
func (c *Checker) Crawl(page *url.URL) {
	req, err := http.NewRequest(http.MethodGet, page.String(), nil)
	if err != nil {
		c.AddResult(page.String(), DEAD, true)
		return
	}
	res, err := c.HttpClient.Do(req)
	if err != nil {
		c.AddResult(page.String(), DEAD, true)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		c.AddResult(page.String(), DEAD, true)
		return
	}
	doc, err := htmlquery.Parse(res.Body)
	if err != nil {
		c.AddResult(page.String(), DEAD, true)
		return
	}

	c.AddResult(page.String(), OKAY, false)
	// We do not want to scroll the entire internet so check the base
	if c.BaseUrl.Host != page.Host {
		return
	}
	// Find all A elements with href attribute and only return href value.
	list := htmlquery.Find(doc, "//a/@href")

	for _, anchor := range list {
		link := htmlquery.SelectAttr(anchor, "href")
		nextUrl, err := url.Parse(link)
		if err != nil {
			c.AddResult(link, DEAD, true)
			return
		}

		if nextUrl.Scheme != "" && notAllowedscheme[nextUrl.Scheme] {
			continue
		}
		resolvedUrl := page.ResolveReference(nextUrl)
		if !c.Visited[resolvedUrl.String()] {
			c.Visited[resolvedUrl.String()] = true
			c.Crawl(resolvedUrl)
		}

	}

}
