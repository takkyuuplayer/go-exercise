package main

import (
	"context"
	"log"
	"net/http"
	"net/http/cookiejar"
	"time"
)

func main() {
	rawurl := "https://docs.google.com/presentation/d/1ubhgnldLgquCWPm5lE19z0bhK9Iqcv26IPU-mvHENh0/edit#slide=id.gab8bba9c57_0_8"
	ctx := context.Background()
	timeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(timeout, http.MethodGet, rawurl, nil)
	if err != nil {
		log.Fatal(err)
	}
	// User-Agent is to avoid the default UA. Otherwise https://www.elastic.co/guide/en/elasticsearch/reference/6.8/search-request-sort.html returns 503.
	// There might be other website.
	req.Header.Add("User-Agent", "UA")
	// cookiejar is for handling cookie during redirection. Otherwise https://webun.jp/item/7724168 returns 404.
	// There may be other sites depending on cookie.
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{Jar: jar}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = res.Body.Close()
	}()

	log.Printf("%s\n", res.Request.URL.Hostname())
	log.Printf("%s\n", res.Request.URL.String())
}
