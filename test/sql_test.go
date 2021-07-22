package test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Id   int
	Name string
}

type Group struct {
	Id   int
	Name string
}

func TestScan(t *testing.T) {
	db := mysqlDb(t)

	rows, err := db.Query(`SELECT u.*, g.* FROM users u
    INNER JOIN group_users ug ON ug.user_id = u.id
    INNER JOIN groups g ON g.id = ug.group_id`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []*User
	var groups []*Group

	for rows.Next() {
		var user User
		var group Group
		err := rows.Scan(&user.Id, &user.Name, &group.Id, &group.Name)

		assert.Nil(t, err)

		users = append(users, &user)
		groups = append(groups, &group)
	}

	assert.Len(t, users, 1)
	assert.Len(t, groups, 1)
	assert.Equal(t, &User{1, "user1"}, users[0])
	assert.Equal(t, &Group{1, "group1"}, groups[0])
}

func mysqlDb(t *testing.T) *sql.DB {
	t.Helper()

	if dsn, ok := os.LookupEnv("MYSQL_DSN"); ok {
		if db, err := sql.Open("mysql", dsn); err != nil {
			t.Fatal(err)
		} else {
			queries := []string{
				"SET FOREIGN_KEY_CHECKS = 0",
				"TRUNCATE users",
				"TRUNCATE groups",
				"TRUNCATE group_users",
				"SET FOREIGN_KEY_CHECKS = 1",
				"INSERT INTO users VALUES (1, 'user1'), (2, 'user2')",
				"INSERT INTO groups VALUES (1, 'group1')",
				"INSERT INTO group_users VALUES (1, 1, 1)",
			}
			for _, query := range queries {
				if _, err := db.Exec(query); err != nil {
					t.Fatal(err)
				}
			}
			return db
		}
	}

	t.Fatal("No MYSQL_DSN env")

	return nil
}
