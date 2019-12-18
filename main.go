// outline prints the structure of an HTML doc.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	rootNode, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}

	outline(nil, rootNode)
}

func outline(stack []string, htmlNode *html.Node) {
	if htmlNode.Type == html.ElementNode {
		stack = append(stack, htmlNode.Data)
		fmt.Println(stack)
	}

	fmt.Println("Entering loop...")
	for child := htmlNode.FirstChild; child != nil; child = child.NextSibling {
		outline(stack, child)
	}
}
