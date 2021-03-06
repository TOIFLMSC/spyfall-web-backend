package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/TOIFLMSC/spyfall-web-backend/internal/app/store/sqlstore"
)

// Start func
func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)

	srv := newServer(store)

	go func() {
		http.ListenAndServe(config.BindAddr, srv)
	}()

	select {}

	return nil
}

// newDB func
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
