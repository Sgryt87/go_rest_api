package sqlstore

import (
	"database/sql"
	"strings"
	"testing"
)

func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			err, _ := db.Exec("TRUNCATE %s CASCADE", strings.Join(tables, ", "))
			if err != nil {
				t.Fatal(err)
			}
		}
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}
}
