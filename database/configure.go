package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //importing mysql driver
)

// DBConfig struct for keeping database configuration
type DBConfig struct {
	Alias    string
	Driver   string
	host     string
	name     string
	user     string
	password string
	port     string
}

// DB a structure containing the actual db connection
var DB *gorm.DB

// getDBConfig returns the database configuration
// stored in the .env file
func getDBConfig() *DBConfig {

	alias := os.Getenv("DB_ALIAS")
	if alias == "" {
		alias = "default"
	}

	return &DBConfig{
		Alias:    alias,
		Driver:   "mysql",
		host:     os.Getenv("DB_HOST"),
		name:     os.Getenv("DB_NAME"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASS"),
		port:     os.Getenv("DB_PORT"),
	}
}

// getConnectionString returns the url connection string
// for the configured databases
func getConnectionString(config *DBConfig) string {
	// Connection string format: <username>:<password>@<network-type>(<host>:<port>)/<dbname>
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", config.user, config.password, config.host, config.port, config.name)
}

// InitializeDBConnection return pointer for executing queries
func InitializeDBConnection() *gorm.DB {

	config := getDBConfig()

	var err error
	DB, err = gorm.Open(config.Driver, getConnectionString(config))
	if err != nil {
		panic(err)
	}

	return DB
}
