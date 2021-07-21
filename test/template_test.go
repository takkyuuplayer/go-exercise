package test

import (
	"bytes"
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplate(t *testing.T) {
	tpl := template.Must(template.New("foo").Parse(`<a href="{{.Url}}"></a>`))
	data := map[string]string{
		"Url": "https://example.com/?a=あいうえお&b=https://takkyuuplayer.com/",
	}
	var out = new(bytes.Buffer)
	tpl.Execute(out, data)

	assert.Equal(t, out, "foo")

	t.Log(out)
	t.Log(template.HTMLEscape("a=あいうえお&b=https://takkyuuplayer.com/"))
}
