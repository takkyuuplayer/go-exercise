package test

import (
	"context"
	"database/sql"
	"log"
	"os"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLStruct struct {
	Id int
}

func TestRowScan(t *testing.T) {
	db := mysqlDb(t)

	rows, err := db.QueryContext(context.Background(), `
    (SELECT 50 as id FROM information_schema.TABLES LIMIT 1)
    UNION ALL
    (SELECT NULL FROM information_schema.TABLES LIMIT 1)
    `)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var structs []*MySQLStruct

	for rows.Next() {
		item := &MySQLStruct{}
		field := reflect.ValueOf(item).Elem().Field(0).Interface()
		values := []interface{}{&field}
		if err := rows.Scan(values...); err != nil {
			log.Fatal(err)
		}
		if reflect.ValueOf(values[0]).IsNil() {
			item = nil
		} else {
			t.Logf("%#v", reflect.ValueOf(values[0]).Elem().Type())
			reflect.ValueOf(item).Elem().Field(0).Set(
				reflect.ValueOf(values[0]).Elem(),
			)
		}
		structs = append(structs, item)
	}
	// If the database is being written to ensure to check for Close
	// errors that may be returned from the driver. The query may
	// encounter an auto-commit error and be forced to rollback changes.
	rerr := rows.Close()
	if rerr != nil {
		log.Fatal(rerr)
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	t.Log(structs[0])
	t.Log(structs[1])
}

func mysqlDb(t *testing.T) *sql.DB {
	t.Helper()

	if dsn, ok := os.LookupEnv("MYSQL_DSN"); ok {
		if db, err := sql.Open("mysql", dsn); err != nil {
			t.Fatal(err)
		} else {
			return db
		}
	}

	t.Fatal("No MYSQL_DSN env")

	return nil
}
