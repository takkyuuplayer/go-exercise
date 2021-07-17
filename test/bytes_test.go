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
	in := strings.NewReader("test")
	var out bytes.Buffer

	upper(in, &out)

	assert.Equal(t, "TEST\n", out.String())
}

func upper(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		if err := scanner.Err(); err == nil {
			fmt.Fprintln(out, strings.ToUpper(scanner.Text()))
		}
	}
}
