package env_test

import (
	"testing"

	"github.com/caarlos0/env/v11"
	"github.com/stretchr/testify/assert"
)

func TestEnv_ParseAs(t *testing.T) {
	type Bar struct {
		Value string `env:"Bar"`
	}
	t.Setenv("Foo", "1")
	t.Setenv("Bar", "2")

	t.Run("Without reference", func(t *testing.T) {
		type Config struct {
			Foo string `env:"Foo"`
			Bar Bar
		}

		cfg, err := env.ParseAs[Config]()

		assert.NoError(t, err)
		assert.Equal(
			t,
			Config{
				Foo: "1",
				Bar: Bar{
					Value: "2",
				},
			},
			cfg,
		)
	})

	t.Run("With reference", func(t *testing.T) {
		type Config struct {
			Foo string `env:"Foo"`
			Bar *Bar
		}

		cfg, err := env.ParseAs[Config]()

		assert.NoError(t, err)
		assert.Equal(
			t,
			Config{
				Foo: "1",
				Bar: nil,
			},
			cfg,
		)
	})
}
