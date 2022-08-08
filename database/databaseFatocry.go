package database

import (
	"fmt"

	"gorm.io/gorm"
)

type IDatabaseFactory interface {
	connect(host, port, user, password, dbname string) error
	getDB() *gorm.DB
}

func getDatabaseFactory(db string) (IDatabaseFactory, error) {
	switch db {
	case "postgres":
		return &Postgres{}, nil
	}

	return nil, fmt.Errorf(("wrong database type passed"))
}
