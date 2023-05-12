package repositories

import "database/sql"

type repository struct {
	db *sql.DB
}