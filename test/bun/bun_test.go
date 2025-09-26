package bun_test

import (
	"database/sql"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func TestBung(t *testing.T) {
	t.Parallel()

	type User struct {
		bun.BaseModel `bun:"table:users,alias:u"`

		ID   int64 `bun:",pk,autoincrement"`
		Name string
	}

	tmpDir := t.TempDir()
	file := filepath.Join(tmpDir, "test.db")

	sqldb, err := sql.Open(sqliteshim.ShimName, file)
	require.NoError(t, err)

	db := bun.NewDB(sqldb, sqlitedialect.New())
	_, err = db.NewCreateTable().Model((*User)(nil)).Exec(t.Context())
	require.NoError(t, err)

	user := &User{Name: "admin"}
	_, err = db.NewInsert().Model(user).Exec(t.Context())
	require.NoError(t, err)

	users := make([]*User, 0)
	err = db.NewRaw(
		"SELECT id, name FROM ? LIMIT ?",
		bun.Ident("users"), 100,
	).Scan(t.Context(), &users)
	require.NoError(t, err)

	require.Len(t, users, 1)
	require.Equal(
		t,
		[]*User{
			{
				ID:   1,
				Name: "admin",
			},
		},
		users,
	)

}
