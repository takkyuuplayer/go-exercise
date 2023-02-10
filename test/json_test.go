package test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	t.Run("null to boolean", func(t *testing.T) {
		code := `{"bool_field": null}`
		var params struct {
			BoolField bool `json:"bool_field"`
		}
		assert.NoError(t, json.Unmarshal([]byte(code), &params))
		assert.False(t, params.BoolField)
	})
}
