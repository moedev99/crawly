package crawly

import (
	"fmt"
	"sync"

	"github.com/fatih/color"
	"github.com/moedev99/crawly/checker"
)

var usage = `Usage: crawly URL
Example:
	crawly https://www.example.com`

func Main(args []string, c *checker.Checker) int {
	if len(args) == 0 || len(args) >= 2 {
		fmt.Println(usage)
		return 1
	}
	site := args[0]

	var wg sync.WaitGroup
	wg.Go(
		func() {
			c.Check(site)
		},
	)
	wg.Wait()
	fmt.Fprint(c.Output, color.YellowString("FINAL REPORT:\n"))
	for _, result := range c.Results {
		fmt.Fprintf(c.Output, "%s status -> %q\n", result.Status, result.Link)
	}

	return 0
}
