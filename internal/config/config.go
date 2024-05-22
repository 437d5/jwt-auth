package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"net/url"
	"os"
)

type Config struct {
	Token JWTToken `yaml:"jwt_tokens"`
	DB    Database `yaml:"database"`
	SRV   Server   `yaml:"server"`
}

type JWTToken struct {
	Secret string `yaml:"secret"`
}

type Database struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Username     string `yaml:"user"`
	Password     string `yaml:"pwd"`
	DatabaseName string `yaml:"dbname"`
}

type Server struct {
	Port string `yaml:"port"`
}

func NewConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c Config) CreateConnString() string {
	dbURL := &url.URL{
		Scheme: "mongodb",
		User:   url.UserPassword(c.DB.Username, c.DB.Password),
		Host:   fmt.Sprintf("%s:%d", c.DB.Host, c.DB.Port),
		Path:   c.DB.DatabaseName,
	}
	log.Print(dbURL)
	return dbURL.String()
}
