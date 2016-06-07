package ui

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"github.com/kr/pretty"
)

func parseMeta(htmlSrc []byte) (*meta, error) {
	root, err := html.Parse(bytes.NewReader(htmlSrc))
	if err != nil {
		return nil, err
	}

	var m meta
	walk(root, func(n *html.Node) bool {
		if n.Type == html.DocumentNode {
			return true
		}
		switch n.DataAtom {
		case atom.Title:
			m.Title = n.FirstChild.Data

		case atom.Html, atom.Head:
			return true // traverse
		}
		return false
	})

	return &m, nil
}

func walk(n *html.Node, fn func(*html.Node) bool) {
	if fn(n) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c, fn)
		}
	}
}

func TestParseMeta(t *testing.T) {
	tests := []struct {
		html string
		want meta
	}{
		{
			html: "<html><head></head></html>",
			want: meta{Title: ""},
		},
		{
			html: "<html><head><title>mytitle</title></head><body></body></html>",
			want: meta{Title: "mytitle"},
		},
	}
	for _, test := range tests {
		meta, err := parseMeta([]byte(test.html))
		if err != nil {
			t.Errorf("%q: %s", test.html, err)
			continue
		}
		if !reflect.DeepEqual(*meta, test.want) {
			t.Errorf("meta mismatch for HTML: %q\n\n%s", test.html, strings.Join(pretty.Diff(*meta, test.want), "\n"))
		}
	}
}
