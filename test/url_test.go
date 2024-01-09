package test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Parallel()
	t.Run("Parse", func(t *testing.T) {
		t.Parallel()
		u, _ := url.Parse("http://あ.example/い?search=う#え")

		assert.Equal(
			t,
			&url.URL{
				Scheme:      "http",
				Host:        "あ.example",
				Path:        "/い",
				RawPath:     "/い",
				RawQuery:    "search=う",
				Fragment:    "え",
				RawFragment: "え",
			},
			u,
		)
		assert.Equal(t, "http://%E3%81%82.example/%E3%81%84?search=う#%E3%81%88", u.String())
		assert.Equal(t, "/%E3%81%84", u.EscapedPath())
		assert.Equal(t, "search=%E3%81%86", u.Query().Encode())
		assert.Equal(t, "%E3%81%88", u.EscapedFragment())

		u.RawQuery = u.Query().Encode()
		assert.Equal(t, "http://%E3%81%82.example/%E3%81%84?search=%E3%81%86#%E3%81%88", u.String())
	})
}
