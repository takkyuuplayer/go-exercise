package test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go:generate go-assets-builder -s "/testdata" -p test -o test/asset.go testdata/

func TestAssets(t *testing.T) {
	f, _ := Assets.Open("/test.txt")

	buf := new(bytes.Buffer)
	buf.ReadFrom(f)

	assert.Equal(t, "Hello World\n", buf.String())
}
