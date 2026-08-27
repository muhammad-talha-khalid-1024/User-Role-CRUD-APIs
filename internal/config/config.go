package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBCharset  string
	ServerPort string
}

func LoadConfig() *Config {
	port, _ := strconv.Atoi(getEnv("DB_PORT", "3306"))

	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     port,
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "mustaqel_on_go"),
		DBCharset:  getEnv("DB_CHARSET", "utf8mb4"),
		ServerPort: getEnv("SERVER_PORT", ":8080"),
	}
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
		c.DBCharset,
	)
}

func ConnectDB(cfg *Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.GetDSN())
	if err != nil {
		return nil, err
	}

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * 60) // 5 minutes

	log.Println("MySQL database connected successfully")
	return db, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
