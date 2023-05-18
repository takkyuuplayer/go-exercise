package i18n_test

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

var GolangMessage = map[language.Tag]*message.Printer{}

func init() {
	cat := catalog.NewBuilder()

	cat.SetString(language.Japanese, "おはよう%sさん", "おはよう%sさん")
	cat.SetString(language.Japanese, "今朝はりんごを%d個食べました", "今朝はりんごを%d個食べました")

	cat.SetString(language.English, "おはよう%sさん", "Good morning %s")
	cat.SetString(language.English, "今朝はりんごを%d個食べました", "I ate %d ${apples(1)} this morning")
	cat.SetMacro(language.English, "apples", plural.Selectf(1, "", plural.One, "apple"))

	GolangMessage[language.Japanese] = message.NewPrinter(language.Japanese, message.Catalog(cat))
	GolangMessage[language.English] = message.NewPrinter(language.English, message.Catalog(cat))
}

func TestSprintf(t *testing.T) {
	assert.Equal(t, "おはようtakkyuuplayerさん", GolangMessage[language.Japanese].Sprintf("おはよう%sさん", "takkyuuplayer"))
	assert.Equal(t, "今朝はりんごを1個食べました", GolangMessage[language.Japanese].Sprintf("今朝はりんごを%d個食べました", 1))
	assert.Equal(t, "今朝はりんごを2個食べました", GolangMessage[language.Japanese].Sprintf("今朝はりんごを%d個食べました", 2))
	assert.Equal(t, "Good morning takkyuuplayer", GolangMessage[language.English].Sprintf("おはよう%sさん", "takkyuuplayer"))
	GolangMessage[language.English].Sprintf("今朝はりんごを%d個食べました", 1)
	assert.Equal(t, "I ate 1 apple this morning", GolangMessage[language.English].Sprintf("今朝はりんごを%d個食べました", 1))
	assert.Equal(t, "I ate 2 apples this morning", GolangMessage[language.English].Sprintf("今朝はりんごを%d個食べました", 2))
}

func BenchmarkGolangMessage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GolangMessage[language.Japanese].Sprintf("おはよう%sさん", "takkyuuplayer")
		GolangMessage[language.Japanese].Sprintf("今朝はりんごを%d個食べました", 1)
		GolangMessage[language.Japanese].Sprintf("今朝はりんごを%d個食べました", 2)

		GolangMessage[language.English].Sprintf("おはよう%sさん", "takkyuuplayer")
		GolangMessage[language.English].Sprintf("今朝はりんごを%d個食べました", 1)
		GolangMessage[language.English].Sprintf("今朝はりんごを%d個食べました", 2)
	}
}

var Localizer = map[language.Tag]*i18n.Localizer{}

func init() {
	ja := i18n.NewBundle(language.Japanese)
	ja.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	ja.LoadMessageFile("active.ja.toml")
	Localizer[language.Japanese] = i18n.NewLocalizer(ja, "ja")

	en := i18n.NewBundle(language.English)
	en.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	en.LoadMessageFile("translate.en.toml")
	Localizer[language.English] = i18n.NewLocalizer(en, "en")
}

func TestMustLocalize(t *testing.T) {
	assert.Equal(t, "おはようtakkyuuplayerさん",
		Localizer[language.Japanese].MustLocalize(
			&i18n.LocalizeConfig{MessageID: "おはよう%sさん", TemplateData: map[string]interface{}{"Name": "takkyuuplayer"}},
		))
	assert.Equal(t, "今朝はりんごを1個食べました",
		Localizer[language.Japanese].MustLocalize(
			&i18n.LocalizeConfig{MessageID: "今朝はりんごを%d個食べました", PluralCount: 1},
		))
	assert.Equal(t, "今朝はりんごを2個食べました",
		Localizer[language.Japanese].MustLocalize(
			&i18n.LocalizeConfig{MessageID: "今朝はりんごを%d個食べました", PluralCount: 2},
		))

	assert.Equal(t, "Good morning takkyuuplayer",
		Localizer[language.English].MustLocalize(
			&i18n.LocalizeConfig{MessageID: "おはよう%sさん", TemplateData: map[string]interface{}{"Name": "takkyuuplayer"}},
		))
	assert.Equal(t, "I ate 1 apple",
		Localizer[language.English].MustLocalize(
			&i18n.LocalizeConfig{MessageID: "今朝はりんごを%d個食べました", PluralCount: 1},
		))
	assert.Equal(t, "I ate 2 apples",
		Localizer[language.English].MustLocalize(
			&i18n.LocalizeConfig{MessageID: "今朝はりんごを%d個食べました", PluralCount: 2},
		))
}

func BenchmarkMustLocalize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Localizer[language.Japanese].MustLocalize(
			&i18n.LocalizeConfig{MessageID: "おはよう%sさん", TemplateData: map[string]interface{}{"Name": "takkyuuplayer"}},
		)
		Localizer[language.Japanese].MustLocalize(
			&i18n.LocalizeConfig{MessageID: "今朝はりんごを%d個食べました", PluralCount: 1},
		)
		Localizer[language.Japanese].MustLocalize(
			&i18n.LocalizeConfig{MessageID: "今朝はりんごを%d個食べました", PluralCount: 2},
		)
		Localizer[language.English].MustLocalize(
			&i18n.LocalizeConfig{MessageID: "おはよう%sさん", TemplateData: map[string]interface{}{"Name": "takkyuuplayer"}},
		)
		Localizer[language.English].MustLocalize(
			&i18n.LocalizeConfig{MessageID: "今朝はりんごを%d個食べました", PluralCount: 1},
		)
		Localizer[language.English].MustLocalize(
			&i18n.LocalizeConfig{MessageID: "今朝はりんごを%d個食べました", PluralCount: 2},
		)
	}
}
