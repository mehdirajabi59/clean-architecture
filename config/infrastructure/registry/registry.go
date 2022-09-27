package registry

import "database/sql"

type Registry struct {
	db *sql.DB
}

func NewRegistry(db *sql.DB) *Registry {
	return &Registry{
		db,
	}
}
