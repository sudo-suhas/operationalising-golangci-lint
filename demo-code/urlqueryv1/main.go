package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
)

func main() {
	if err := demo(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// START // OMIT

func demo() error {
	u, err := url.Parse("http://my-search-engine.com")
	if err != nil {
		return fmt.Errorf("demo: parse url: %w", err)
	}

	p := SearchParams{
		Search: "linters",
		Limit:  5,
		Offset: 10,
	}
	p.SetQueryParams(u)

	fmt.Println("URL:", u)
	// [Expected] URL: http://my-search-engine.com?limit=5&offset=10&search=linters // HL
	return nil
}

func (p SearchParams) SetQueryParams(u *url.URL) {
	u.Query().Set("search", p.Search)
	u.Query().Set("limit", strconv.Itoa(p.Limit))
	u.Query().Set("offset", strconv.Itoa(p.Offset))
}

// END // OMIT

type SearchParams struct {
	Search string
	Limit  int
	Offset int
}
