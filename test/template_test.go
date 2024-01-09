package test

import (
	"bytes"
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplate(t *testing.T) {
	t.Parallel()
	tpl := template.Must(template.New("foo").Parse(`<a href="{{.Url}}"></a>`))
	data := map[string]string{
		"Url": "https://example.com/?a=あいうえお&b=https://takkyuuplayer.com/",
	}
	var out = new(bytes.Buffer)
	tpl.Execute(out, data)

	assert.Equal(t, "<a href=\"https://example.com/?a=%e3%81%82%e3%81%84%e3%81%86%e3%81%88%e3%81%8a&amp;b=https://takkyuuplayer.com/\"></a>", out.String())
}
