package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func (s *Store) CreatePostgreClient(url string) {

	db, err := sql.Open("postgres", url)
	if err != nil {
		panic("Error in connection Postgres")
	}

	if err := db.Ping(); err != nil {
		panic("Error in connecting Postgres")
	}

	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(5)

	s.Postgre = db
}
