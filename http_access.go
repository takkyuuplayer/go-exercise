package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
	"time"
)

var Transport http.RoundTripper = &http.Transport{
	TLSClientConfig: &tls.Config{Renegotiation: tls.RenegotiateOnceAsClient},
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	ctx := context.Background()
	for stdin.Scan() {
		url := stdin.Text()
		timeout, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		req, _ := http.NewRequestWithContext(timeout, http.MethodGet, url, nil)
		jar, _ := cookiejar.New(nil)
		client := http.Client{Jar: jar, Transport: Transport}

		res, _ := client.Do(req)
		defer res.Body.Close()

		lastUrl := res.Request.URL.String()

		fmt.Println(strings.Repeat("-", 50))
		fmt.Printf("Before: %s\tAfter: %s\n", url, lastUrl)
		fmt.Printf("Headers: %v", res.Header)
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Printf("Body: %v", string(body))
	}
}
