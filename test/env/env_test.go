package env_test

import (
	"testing"

	"github.com/caarlos0/env/v11"
	"github.com/stretchr/testify/assert"
)

func TestEnv_ParseAs(t *testing.T) {
	t.Setenv("ENV_1", "1")
	t.Setenv("ENV_2", "2")
	t.Setenv("ENV_3", "")

	type Env2 struct {
		Value string `env:"ENV_2"`
	}

	t.Run("Without reference", func(t *testing.T) {
		t.Parallel()

		type Config struct {
			One   string `env:"ENV_1"`
			Two   Env2
			Three string `env:"ENV_3" envDefault:"default"`
		}

		cfg, err := env.ParseAs[Config]()

		assert.NoError(t, err)
		assert.Equal(
			t,
			Config{
				One: "1",
				Two: Env2{
					Value: "2",
				},
				Three: "default",
			},
			cfg,
		)
	})

	t.Run("With reference", func(t *testing.T) {
		t.Parallel()
		type Config struct {
			One   string `env:"ENV_1"`
			Two   *Env2
			Three string `env:"ENV_3" envDefault:"default"`
		}

		cfg, err := env.ParseAs[Config]()

		assert.NoError(t, err)
		assert.Equal(
			t,
			Config{
				One:   "1",
				Two:   nil,
				Three: "default",
			},
			cfg,
		)
	})
}
