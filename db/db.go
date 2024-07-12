package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kviatkovsky/auth_service/internal/config"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(cfg *config.Config) (*Database, error) {
	db, err := sql.Open("mysql", dsn(&cfg.MySQL))
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	if err := d.db.Close(); err != nil {
		panic(err)
	}
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}

func dsn(cfg *config.StorageConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
}
