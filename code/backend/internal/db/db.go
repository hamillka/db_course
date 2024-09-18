package db

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // revive comment
	"go.uber.org/zap"
)

type DatabaseConfig struct {
	DBHost string `envconfig:"HOST"`
	DBPort string `envconfig:"PORT"`
	DBName string `envconfig:"NAME"`
	DBUser string `envconfig:"USER"`
	DBPass string `envconfig:"PASS"`
}

type DatabaseInstance struct {
	conn               *sqlx.DB
	logger             *zap.SugaredLogger
	config             DatabaseConfig
	maxConnectAttempts int64
}

func NewConn(
	config *DatabaseConfig,
	maxAttempts int64,
	logger *zap.SugaredLogger,
) *DatabaseInstance {
	instance := DatabaseInstance{
		conn:               nil,
		config:             *config,
		logger:             logger,
		maxConnectAttempts: maxAttempts,
	}

	return &instance
}

func (db *DatabaseInstance) GetConn() *sqlx.DB {
	var err error

	if db.conn == nil {
		if db.conn, err = db.reconnect(); err != nil {
			log.Fatalf("%s", err)
		}
	}

	if err = db.conn.Ping(); err != nil {
		var attempt int64
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			if attempt >= db.maxConnectAttempts {
				log.Printf("connection failed after %d attempt\n", attempt)
			}
			attempt++

			log.Println("reconnecting...")

			db.conn, err = db.reconnect()
			if err == nil {
				return db.conn
			}

			log.Printf("connection was lost. Error: %s. Waiting for 5 sec...\n", err)
		}
	}

	return db.conn
}

func (db *DatabaseInstance) reconnect() (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		db.config.DBHost, db.config.DBUser, db.config.DBPass, db.config.DBName, db.config.DBPort)

	conn, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
