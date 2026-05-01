package database

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func getDbDSN(config DatabaseConfig) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.Username, config.Password, config.Database)
}

var db *gorm.DB

func ConnectDatabase() *gorm.DB {
	if db != nil {
		return db
	}
	dbConfig := DatabaseConfig{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "postgres",
		Database: "register_service",
	}
	db, err := gorm.Open(postgres.Open(getDbDSN(dbConfig)), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Msg("failed to connect database")
	}
	log.Info().Msg("Database connected successfully")
	return db
}

func Client() *gorm.DB {
	if db != nil {
		return db
	}
	panic("Database is not connected!")
}

func CloseConnection() {
	log.Info().Msg("Closing database connection")

	conn, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("failed to get database connection")
	}
	if err = conn.Close(); err != nil {
		log.Error().Err(err).Msg("failed to close database connection")
	}
}
