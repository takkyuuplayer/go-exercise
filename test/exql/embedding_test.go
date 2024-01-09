package exql_test

import (
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/loilo-inc/exql/v2"
	"github.com/stretchr/testify/assert"
)

type exqlDB exql.DB

type (
	DB struct {
		exqlDB
	}
)

func TestEmbedding(t *testing.T) {
	t.Parallel()
	db, err := exql.Open(&exql.OpenOptions{Url: os.Getenv("MYSQL_DSN")})
	assert.NoError(t, err)

	var d exql.DB = &DB{db}
	assert.NotNil(t, d)
}
