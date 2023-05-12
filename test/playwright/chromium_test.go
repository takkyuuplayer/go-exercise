package playwright_test

import (
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/assert"
)

func TestChromium(t *testing.T) {
	runOptions := &playwright.RunOptions{
		Browsers: []string{"chromium"},
	}
	assert.NoError(t, playwright.Install(runOptions))

	pw, err := playwright.Run()
	assert.NoError(t, err)

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	assert.NoError(t, err)

	page, err := browser.NewPage()
	assert.NoError(t, err)

	_, err = page.Goto("https://news.ycombinator.com")
	assert.NoError(t, err)

	entries, err := page.QuerySelectorAll(".athing")
	assert.NoError(t, err)

	for i, entry := range entries {
		titleElement, err := entry.QuerySelector("td.title > span > a")
		assert.NoError(t, err)

		title, err := titleElement.TextContent()
		assert.NoError(t, err)

		t.Logf("%d: %s\n", i+1, title)
	}
	assert.NoError(t, browser.Close())

	assert.NoError(t, pw.Stop())
}
