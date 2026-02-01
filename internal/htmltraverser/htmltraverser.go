package htmltraverser

import "golang.org/x/net/html"

// GetAnchorsDFS traverse though the html document in DFT.
func GetAnchorsDFS(node *html.Node, anchors *[]*html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		*anchors = append(*anchors, node)
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		GetAnchorsDFS(child, anchors)
	}

}
