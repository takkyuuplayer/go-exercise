package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	t.Parallel()

	assert.Equal(t, time.Time{}.String(), "0001-01-01 00:00:00 +0000 UTC")
}
