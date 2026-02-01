package checker

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func extractUrl(attrs []html.Attribute, baseUrl string) string {
	for _, val := range attrs {
		if val.Key == "href" {
			normalizedUrl := normalizeUrl(val.Val, baseUrl)
			if normalizedUrl == "" {
				continue
			}

			if !uniqueUrls[normalizedUrl] {
				uniqueUrls[normalizedUrl] = true

				return normalizedUrl
			}
		}
	}
	return ""
}

func normalizeUrl(rawUrl string, baseUrl string) string {
	base, err := url.Parse(baseUrl)
	if err != nil {
		return ""
	}

	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		return ""
	}

	// This handles BOTH relative and absolute URLs correctly!
	absolute := base.ResolveReference(parsedURL)

	absolute.Scheme = strings.ToLower(absolute.Scheme)
	absolute.Host = strings.ToLower(absolute.Host)
	absolute.Fragment = ""

	if absolute.Path != "/" {
		absolute.Path = strings.TrimSuffix(absolute.Path, "/")
	}

	return absolute.String()
}
