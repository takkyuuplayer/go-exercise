package test

import (
	"io/ioutil"
	"testing"
)

func TestOpen(t *testing.T) {
	data, err := ioutil.ReadFile("../testdata/test.txt")

	if err != nil {
		t.Fatal(err)
	}

	if string(data) != "Hello World\n" {
		t.Errorf(`string(data) = %#v, want %#v`, string(data), "Hello World\n")
	}
}
