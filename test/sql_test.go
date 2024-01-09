// nolint: paralleltest
package test

import (
	"database/sql"
	"log"
	"os"
	"reflect"
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
    INNER JOIN group_users ug ON ug.user_id = u.id ` +
		"INNER JOIN `groups` g ON g.id = ug.group_id")
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

func TestScanWithReflection(t *testing.T) {
	db := mysqlDb(t)

	rows, err := db.Query(`SELECT u.*, g.* FROM users u
    INNER JOIN group_users ug ON ug.user_id = u.id ` +
		"INNER JOIN `groups` g ON g.id = ug.group_id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []*User
	var groups []*Group

	for rows.Next() {
		var user = &User{}
		var group = &Group{}
		values := make([]interface{}, 4)
		values[0] = reflect.ValueOf(user).Elem().Field(0).Addr().Interface()
		values[1] = reflect.ValueOf(user).Elem().Field(1).Addr().Interface()
		values[2] = reflect.ValueOf(group).Elem().Field(0).Addr().Interface()
		values[3] = reflect.ValueOf(group).Elem().Field(1).Addr().Interface()

		err := rows.Scan(values...)

		assert.Nil(t, err)

		users = append(users, user)
		groups = append(groups, group)
	}

	assert.Len(t, users, 1)
	assert.Len(t, groups, 1)
	assert.Equal(t, &User{1, "user1"}, users[0])
	assert.Equal(t, &Group{1, "group1"}, groups[0])
}

func TestScanLeftJoinWithReflection(t *testing.T) {
	db := mysqlDb(t)

	scan := func(rows *sql.Rows, models ...interface{}) {
		sum := 0
		numFields := map[int]int{}
		types := map[int]reflect.Type{}
		for modelIdx, model := range models {
			types[modelIdx] = reflect.TypeOf(model).Elem().Elem() // model
			numFields[modelIdx] = types[modelIdx].NumField()
			sum += numFields[modelIdx]
		}
		values := make([]interface{}, sum)

		idx := 0
		elements := make([]reflect.Value, len(models))
		for modelIdx := range models {
			elem := reflect.New(types[modelIdx]).Elem() // model
			elements[modelIdx] = elem
			for i := 0; i < numFields[modelIdx]; i++ {
				f := elem.Field(i)
				value := reflect.New(f.Addr().Type()) // **(model.field)
				values[idx] = value.Interface()
				idx++
			}
		}

		err := rows.Scan(values...)

		assert.Nil(t, err)

		idx = 0
		for modelIdx, model := range models {
			if reflect.ValueOf(values[idx]).Elem().IsNil() {
				idx += numFields[modelIdx]
				elem := reflect.ValueOf(model).Elem()
				elem.Set(reflect.Zero(elem.Type()))
				continue
			}
			elem := reflect.New(types[modelIdx]) // *model
			for i := 0; i < numFields[modelIdx]; i++ {
				f := elem.Elem().Field(i)
				f.Set(reflect.ValueOf(values[idx]).Elem().Elem())
				idx++
			}
			reflect.ValueOf(model).Elem().Set(elem)
		}
	}

	t.Run("using *model", func(t *testing.T) {
		rows, err := db.Query(`SELECT u.*, g.* FROM users u
        LEFT JOIN group_users ug ON ug.user_id = u.id ` +
			"LEFT JOIN `groups` g ON g.id = ug.group_id")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []*User
		var groups []*Group

		for rows.Next() {
			var user *User
			var group *Group

			scan(rows, &user, &group)

			users = append(users, user)
			groups = append(groups, group)
		}

		assert.Len(t, users, 2)
		assert.Len(t, groups, 2)
		assert.Equal(t, &User{1, "user1"}, users[0])
		assert.Equal(t, &Group{1, "group1"}, groups[0])
		assert.Equal(t, &User{2, "user2"}, users[1])
		assert.Equal(t, (*Group)(nil), groups[1])
	})

	t.Run("using &model{}", func(t *testing.T) {
		rows, err := db.Query(`SELECT u.*, g.* FROM users u
        LEFT JOIN group_users ug ON ug.user_id = u.id ` +
			"LEFT JOIN `groups` g ON g.id = ug.group_id")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []*User
		var groups []*Group

		for rows.Next() {
			user := &User{}
			group := &Group{}

			scan(rows, &user, &group)

			users = append(users, user)
			groups = append(groups, group)
		}

		assert.Len(t, users, 2)
		assert.Len(t, groups, 2)
		assert.Equal(t, &User{1, "user1"}, users[0])
		assert.Equal(t, &Group{1, "group1"}, groups[0])
		assert.Equal(t, &User{2, "user2"}, users[1])
		assert.Equal(t, (*Group)(nil), groups[1])
	})
}

func mysqlDb(t *testing.T) *sql.DB {
	t.Helper()

	if dsn, ok := os.LookupEnv("MYSQL_DSN"); ok {
		if db, err := sql.Open("mysql", dsn); err != nil {
			t.Fatal(err)
		} else {
			queries := []string{
				"SET FOREIGN_KEY_CHECKS = 0",
				"TRUNCATE `users`",
				"TRUNCATE `groups`",
				"TRUNCATE `group_users`",
				"SET FOREIGN_KEY_CHECKS = 1",
				"INSERT INTO `users` VALUES (1, 'user1'), (2, 'user2')",
				"INSERT INTO `groups` VALUES (1, 'group1')",
				"INSERT INTO `group_users` VALUES (1, 1, 1)",
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
