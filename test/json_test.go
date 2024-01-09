package test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	t.Parallel()
	t.Run("- is ignored", func(t *testing.T) {
		t.Parallel()
		var params struct {
			P1 bool
			P2 bool `json:"-"`
			P3 bool `json:""`
		}
		buf, err := json.Marshal(&params)
		assert.Equal(t, `{"P1":false,"P3":false}`, string(buf))
		assert.NoError(t, err)
	})
}

func TestUnmarshal(t *testing.T) {
	t.Parallel()
	t.Run("null to boolean", func(t *testing.T) {
		t.Parallel()
		code := `{"bool_field": null}`
		var params struct {
			BoolField bool `json:"bool_field"`
		}
		assert.NoError(t, json.Unmarshal([]byte(code), &params))
		assert.False(t, params.BoolField)
	})
}
