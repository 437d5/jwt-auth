package config

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
)

type Config struct {
	Token JWTToken
	DB    Database
	SRV   Server
}

type JWTToken struct {
	Secret string
	ExpAt  int
}

type Database struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DatabaseName string
}

type Server struct {
	Port string
}

func NewConfig() (*Config, error) {
	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "27017"))
	if err != nil {
		log.Fatal("invalid DB_PORT variable")
		return nil, err
	}

	expAt, err := strconv.Atoi(getEnv("EXP_AT", "1"))
	if err != nil {
		log.Fatal("invalid EXP_AT variable")
		return nil, err
	}

	return &Config{
		Token: JWTToken{
			Secret: getEnv("SECRET_KEY", "secret"),
			ExpAt:  expAt,
		},
		DB: Database{
			Host:         getEnv("HOST", "localhost"),
			Port:         dbPort,
			Username:     getEnv("DB_USERNAME", ""),
			Password:     getEnv("PASSWORD", ""),
			DatabaseName: getEnv("DBNAME", ""),
		},
		SRV: Server{
			Port: getEnv("SRV_PORT", "8080"),
		},
	}, nil
}

func (c Config) CreateConnString() string {
	log.Print(c)
	dbURL := &url.URL{
		Scheme: "mongodb",
		User:   url.UserPassword(c.DB.Username, c.DB.Password),
		Host:   fmt.Sprintf("%s:%d", c.DB.Host, c.DB.Port),
		Path:   c.DB.DatabaseName,
	}
	log.Print(dbURL)
	return dbURL.String()
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
