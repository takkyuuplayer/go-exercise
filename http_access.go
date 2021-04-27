package main

import (
	"crypto/x509"
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	_, err := http.DefaultClient.Get("https://www.13hw.com/home/index.html")

	fmt.Printf("%#v\n", err)

	err = err.(*url.Error).Unwrap()

	fmt.Printf("%#v\n", err)

	err = err.(x509.UnknownAuthorityError)

	fmt.Printf("%#v\n", err)
}
