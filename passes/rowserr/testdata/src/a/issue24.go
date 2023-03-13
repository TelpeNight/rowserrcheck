package a

import (
	"database/sql"
)

func issue24_1[T ~int64](db *sql.DB, a T) {
	rows, _ := db.Query("select id from tb") // want "rows.Err must be checked"
	for rows.Next() {
		// ...
	}
}

func issue24_2[T ~int64, DbT interface {
	*sql.DB
	Query(string) (*sql.Rows, error)
}](db DbT, _ T) error {
	rows, _ := db.Query("select id from tb") // want "rows.Err must be checked"
	for rows.Next() {
		// ...
	}
	return nil
}

func issue24_3[T ~int64, DbT interface {
	*sql.DB
	Query(string) (*sql.Rows, error)
}](db DbT, _ T) error {
	rows, _ := db.Query("select 1")
	for rows.Next() {
		// ...
	}
	return rows.Err()
}
