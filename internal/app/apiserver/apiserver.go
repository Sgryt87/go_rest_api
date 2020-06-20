package apiserver

import (
	"database/sql"
	"github.com/gorilla/sessions"
	"net/http"
	"rest_api_yt/internal/app/store/sqlstore"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()
	store := sqlstore.New(db)
	sessionsStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(store, sessionsStore)
	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
