package i18n

import (
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func Greet() string {
	localizer := i18n.NewLocalizer(i18n.NewBundle(language.Japanese), "ja")

	greet := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "おはよう%sさん",
			Other: "おはよう{{.Name}}さん",
		},
		TemplateData: map[string]interface{}{
			"Name": "takkyuuplayer",
		},
	})
	breakfast := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "今朝はりんごを{{.PluralCount}}個食べました",
			Other: "今朝はりんごを{{.PluralCount}}個食べました",
		},
		PluralCount: 5,
	})

	return fmt.Sprintf("%s %s", greet, breakfast)
}
