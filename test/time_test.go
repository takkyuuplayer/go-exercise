package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	t.Parallel()

	assert.Equal(t, fmt.Sprintf("%s", time.Time{}), "0001-01-01 00:00:00 +0000 UTC")
}
