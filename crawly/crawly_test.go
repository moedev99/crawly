package crawly_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/fstest"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/moedev99/crawly/checker"
	"github.com/moedev99/crawly/crawly"
)

func TestCrawlyMain(t *testing.T) {
	out := io.Discard
	clientTest := httptest.NewTLSServer(
		http.FileServerFS(testFS),
	)
	defer clientTest.Close()
	c := checker.NewChecker(out, clientTest.Client())
	t.Logf("Server URL: %s", clientTest.URL)

	for i, r := range c.Results {
		t.Logf("[%d] %+v", i, r)
	}
	crawly.Main([]string{clientTest.URL}, out, c)
	base := clientTest.URL
	want := []checker.Result{
		{Link: base, Status: checker.OKAY},
		{Link: base + "/blog/hello-world.html", Status: checker.OKAY},
		{Link: base + "/blog/follow-up.html", Status: checker.OKAY},
		{Link: base + "/about.html", Status: checker.OKAY},
		{Link: base + "/nonexistent.html", Status: checker.DEAD},
		{Link: base + "/projects.html", Status: checker.DEAD},
		{Link: "httpq://golang.org", Status: checker.DEAD},
	}
	// sort by link
	opts := cmpopts.SortSlices(func(a, b checker.Result) bool {
		return a.Link < b.Link
	})

	got := c.Results
	if !cmp.Equal(want, got, opts) {
		t.Error(cmp.Diff(want, got, opts))
	}

}

var testFS = fstest.MapFS{
	"blog/hello-world.html": {
		Data: []byte(`<html><head><title>Hello World</title></head>
<body>
  <h1>My First Post</h1>
  <p>Welcome! Read more at <a href="httpq://golang.org">Go's website</a></p>

  <a href="/">Home</a>
  <a href="follow-up.html">Follow-up post</a>
  <a href="/about.html">About Me</a>
</body>
</html>`),
	},
	"blog/follow-up.html": {},
	"about.html": {
		Data: []byte(`<html><head><title>About Me</title></head>
<body>
  <h1>About</h1>
  <p>I write about tech.</p>

  <a href="/">Home</a>
  <a href="/blog/hello-world.html">My First Post</a>
  <a href="mailto:me@example.com">Email Me</a>
  <a href="/nonexistent.html">Broken Link</a>
</body>
</html>`),
	},
	"index.html": {
		Data: []byte(`<html><head><title>Home</title></head>
<body>
  <h1>Welcome to My Site</h1>

  <ul>
    <li><a href="blog/hello-world.html">Hello World</a></li>
    <li><a href="about.html">About Me</a></li>
    <li><a href="projects.html">Projects</a></li>
  </ul>

  <a href="mailto:me@example.com">Contact</a>
</body>
</html>`),
	},
	"broken_links.html": {
		Data: []byte(`<html><head><title>Broken Links</title></head>
<body>
  <ul>
    <li><a href="ftp://old-school.html">Bad scheme</a></li>
    <li><a href="httpq://www.invalid_scheme.com">Invalid scheme</a></li>
    <li><a href="">Empty href</a></li>
    <li><a href="/also/missing.html">Missing page</a></li>
  </ul>
</body>
</html>`),
	},
}
