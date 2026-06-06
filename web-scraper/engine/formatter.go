package engine

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
	"text/tabwriter"

	"golang.org/x/net/html"
)

func MatchUrls(r io.Reader, src string) {
	z := html.NewTokenizer(r)

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			if z.Err() == io.EOF {
				return
			}
			fmt.Printf("Error: %s\n", z.Err())
			return
		case html.StartTagToken, html.SelfClosingTagToken:
			tn, hasAttr := z.TagName()
			if string(tn) != "a" || !hasAttr {
				continue
			}
			for {
				key, val, more := z.TagAttr()
				if string(key) == "href" {
					url := string(val)

					if !strings.HasPrefix(url, "/") && !strings.HasPrefix(url, "http") {
						continue
					}

					if len(url) < 2 {
						continue
					}

					if !strings.HasPrefix(url, "http") && strings.HasPrefix(url, "/") {
						url = strings.Join([]string{src, url}, "")
					}

					link := Link{Url: url}

					if strings.Contains(url, "#") || slices.Contains(Links, link) {
						continue
					}
					Links = append(Links, link)
				}
				if !more {
					break
				}
			}
		}
	}
}

func Format(urls []Link) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)

	header := []string{"Url", "IsDead"}

	fmt.Fprintln(w, "")
	fmt.Fprintln(w, strings.Join(header, "\t"))

	for _, u := range urls {
		line := []string{u.Url, strconv.FormatBool(u.IsDead)}

		fmt.Fprintln(w, strings.Join(line, "\t"))
	}
	w.Flush()
	fmt.Println()
}
