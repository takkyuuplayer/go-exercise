package playwright_test

import (
	"testing"

	"github.com/mxschmitt/playwright-go"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
)

func TestChromium(t *testing.T) {
	t.Parallel()
	runOptions := &playwright.RunOptions{
		Browsers: []string{"chromium"},
	}
	require.NoError(t, playwright.Install(runOptions))

	pw, err := playwright.Run()
	require.NoError(t, err)

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	require.NoError(t, err)

	page, err := browser.NewPage()
	require.NoError(t, err)

	err = page.SetContent(`
<table>
	<tr class="athing">
		<td class="title"><span><a>first title</a></span></td>
	</tr>
	<tr class="athing">
		<td class="title"><span><a>second title</a></span></td>
	</tr>
</table>
`)
	assert.NoError(t, err)

	entries, err := page.Locator(".athing").All()
	assert.NoError(t, err)

	for i, entry := range entries {
		titleElement := entry.Locator("td.title > span > a").First()
		title, err := titleElement.TextContent()
		assert.NoError(t, err)

		t.Logf("%d: %s\n", i+1, title)
	}
	assert.NoError(t, browser.Close())

	assert.NoError(t, pw.Stop())
}
