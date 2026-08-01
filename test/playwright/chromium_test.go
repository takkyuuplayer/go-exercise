package playwright_test

import (
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChromium(t *testing.T) {
	t.Parallel()
	runOptions := &playwright.RunOptions{
		Browsers: []string{"chromium"},
	}
	if err := playwright.Install(runOptions); err != nil {
		t.Skipf("skipping playwright test: %v", err)
	}

	pw, err := playwright.Run()
	require.NoError(t, err)

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	require.NoError(t, err)

	page, err := browser.NewPage()
	assert.NoError(t, err)

	_, err = page.Goto("https://news.ycombinator.com")
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
