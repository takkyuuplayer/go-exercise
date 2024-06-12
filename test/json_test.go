package test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null"
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

	t.Run("null.Time", func(t *testing.T) {
		t.Parallel()

		code := `{"field": "2024-01-01T00:00:00Z"}`
		type params struct {
			Field *null.Time `json:"field,omitempty"`
		}
		var p1 params
		assert.NoError(t, json.Unmarshal([]byte(code), &p1))
		assert.Equal(t,
			params{
				Field: &null.Time{
					Valid: true,
					Time:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				}},
			p1,
		)

		code = `{"field": null}`
		var p2 params
		assert.NoError(t, json.Unmarshal([]byte(code), &p2))
		//assert.Equal(t, params{Field: &null.Time{}}, p2) // FAIL!!

		code = `{}`
		var p3 params
		assert.NoError(t, json.Unmarshal([]byte(code), &p3))
		assert.Equal(t, params{Field: nil}, p3)
	})
}
