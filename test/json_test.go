package test

import (
	"database/sql"
	"encoding/json"
	"testing"
	"time"

	null2 "github.com/guregu/null/v5"
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
			Field OptionalParam[null.Time] `json:"field"`
		}
		var p1 params
		assert.NoError(t, json.Unmarshal([]byte(code), &p1))
		assert.Equal(t, params{
			Field: OptionalParam[null.Time]{
				Defined: true,
				Value: null.Time{
					Time:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
					Valid: true,
				},
			},
		}, p1)

		code = `{"field": null}`
		var p2 params
		assert.NoError(t, json.Unmarshal([]byte(code), &p2))
		assert.Equal(t, params{
			Field: OptionalParam[null.Time]{
				Defined: true,
				Value:   null.Time{},
			},
		}, p2)

		code = `{}`
		var p3 params
		assert.NoError(t, json.Unmarshal([]byte(code), &p3))
		assert.Equal(t, params{
			Field: OptionalParam[null.Time]{
				Defined: false,
				Value:   null.Time{},
			},
		}, p3)
	})

	t.Run("null2.Time", func(t *testing.T) {
		t.Parallel()

		code := `{"field": "2024-01-01T00:00:00Z"}`
		type params struct {
			Field *null2.Time `json:"field,omitempty"`
		}
		var p1 params
		assert.NoError(t, json.Unmarshal([]byte(code), &p1))
		assert.Equal(t,
			params{
				Field: &null2.Time{
					NullTime: sql.NullTime{
						Valid: true,
						Time:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
					},
				}},
			p1,
		)

		code = `{"field": null}`
		var p2 params
		assert.NoError(t, json.Unmarshal([]byte(code), &p2))
		//assert.Equal(t, params{Field: &null2.Time{}}, p2) // FAIL!

		code = `{}`
		var p3 params
		assert.NoError(t, json.Unmarshal([]byte(code), &p3))
		assert.Equal(t, params{Field: nil}, p3)
	})
}

type OptionalParam[T any] struct {
	Defined bool
	Value   T
}

func (u *OptionalParam[T]) UnmarshalJSON(data []byte) error {
	u.Defined = true
	return json.Unmarshal(data, &u.Value)
}
