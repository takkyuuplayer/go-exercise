package minna

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssets(t *testing.T) {
	t.Parallel()
	f, _ := Assets.Open("/test.txt")

	buf := new(bytes.Buffer)
	buf.ReadFrom(f)

	assert.Equal(t, "Hello World\n", buf.String())
}
