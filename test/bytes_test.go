package test

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmurateStdinOut(t *testing.T) {
	t.Parallel()
	in := strings.NewReader("test")
	var out bytes.Buffer

	upper(t, in, &out)

	assert.Equal(t, "TEST\n", out.String())
}

func upper(t *testing.T, in io.Reader, out io.Writer) {
	t.Helper()

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		assert.NoError(t, scanner.Err())
		_, err := fmt.Fprintln(out, strings.ToUpper(scanner.Text()))
		assert.NoError(t, err)
	}
}
