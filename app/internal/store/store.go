package store

import (
	"database/sql"

	_ "github.com/lib/pq" // ...
)

// Store ...
type Store struct {
	config *Config
	db     *sql.DB
}

// New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// OpenConnection ...
func (s *Store) OpenConnection() error {
	db, err := sql.Open(s.config.DatabaseName, s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// CloseConnection ...
func (s *Store) CloseConnection() {
	s.db.Close()
}
