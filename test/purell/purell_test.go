package purell_test

import (
	"testing"

	"github.com/PuerkitoBio/purell"
	"github.com/stretchr/testify/assert"
	"github.com/sekimura/go-nomalize-url"
)

func TestURLString(t *testing.T) {
	normalized := purell.MustNormalizeURLString("https://example.com/?a=あいうえお&b=https://takkyuuplayer.com/", purell.FlagsSafe|purell.FlagSortQuery|purell.FlagDecodeUnnecessaryEscapes)

	assert.Equal(t, "https://example.com/?a=%E3%81%82%E3%81%84%E3%81%86%E3%81%88%E3%81%8A&b=https://takkyuuplayer.com/", normalized)

	s, _ := normalizeurl.Normalize("sekimura.org")
	assert.Equal(t, "https://example.com/?a=%E3%81%82%E3%81%84%E3%81%86%E3%81%88%E3%81%8A&b=https://takkyuuplayer.com/", s)
}
