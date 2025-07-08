package storage

import "database/sql"

type Store struct {
	Postgre *sql.DB
}
