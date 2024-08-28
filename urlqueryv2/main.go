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
	return nil
}

// START // OMIT

func (p SearchParams) SetQueryParams(u *url.URL) {
	qry := u.Query()
	qry.Set("search", p.Search)
	qry.Set("offset", strconv.Itoa(p.Offset))
	qry.Set("limit", strconv.Itoa(p.Limit))
	u.RawQuery = qry.Encode()
}

// END // OMIT

type SearchParams struct {
	Search string
	Limit  int
	Offset int
}
