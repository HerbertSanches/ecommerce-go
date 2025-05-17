package db

import (
	"database/sql"
	"fmt"

	"github.com/HerbertSanches/ecommerce-go/config"
	_ "github.com/lib/pq"
)

// type PostgresConfig struct {
// 	Host     string
// 	Port     string
// 	User     string
// 	Password string
// 	DBName   string
// 	SSLMode  string
// }


func NewPostgreStorage(cfg config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar com o banco: %w", err)
	}

	return db, nil
}