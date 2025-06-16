package db

import (
	"database/sql"
	"fmt"

	"github.com/almat-kst10/message-service/configs"
	_ "github.com/lib/pq"
)

func NewPostgresClient(cfg *configs.Configs) (*sql.DB, error) {
	d := `
		host=%s 
		port=%s 
		user=%s 
		dbname=%s 
		password=%s 
		sslmode=%s
		client_encoding=%s
	`
	dsn := fmt.Sprintf(d, cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.Name, cfg.DB.Psw, cfg.DB.SllMode, cfg.DB.Encoding)

	db, err := sql.Open(cfg.DB.Driver, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}