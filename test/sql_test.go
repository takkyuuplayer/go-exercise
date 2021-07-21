package test

import (
	"context"
	"database/sql"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestRowScan(t *testing.T) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:53306)/")
	if err != nil {
		panic(err)
	}
	rows, err := db.QueryContext(context.Background(), `
    (SELECT 50 FROM information_schema.TABLES LIMIT 1)
    UNION ALL
    (SELECT NULL FROM information_schema.TABLES LIMIT 1)
    `)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	numbers := make([]*int64, 0)

	for rows.Next() {
		var number *int64
		if err := rows.Scan(&number); err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
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
	t.Log(numbers)
	t.Log(*numbers[0])
	t.Log(numbers[1])
}
