package main

import (
	"fmt"
	"net/http"
)

func main() {
	_, err := http.DefaultClient.Get("https://onkyo.com/rakuraku/tenji/2014/abu02.htm")

	fmt.Printf("%#v", err)
}
