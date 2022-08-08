package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Postgres struct {
	db *gorm.DB
}

func (p *Postgres) connect(host, port, user, password, dbname string) error {
	info := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		host, port, user, password, dbname, "disable", "Asia/Seoul")

	config := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "gin.", // schema name
			SingularTable: true,
		},
	}

	var err error
	p.db, err = gorm.Open(postgres.Open(info), &config)

	if err != nil {
		return err
	}

	sqlDB, err := p.db.DB()
	if err != nil {
		return err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(3)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
}

func (p *Postgres) getDB() *gorm.DB {
	return p.db
}
