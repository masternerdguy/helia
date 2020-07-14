package sql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/lib/pq" //driver
)

//sharedConfig Shared database connection for services
var sharedConfig *sql.DB = nil

type dbConfig struct {
	DbName string `json:"dbname"`
	DbHost string `json:"dbhost"`
	DbUser string `json:"dbuser"`
	DbPass string `json:"dbpass"`
	DbPort int    `json:"dbport"`
}

func connect() (*sql.DB, error) {
	//check if we are already connected
	if sharedConfig != nil {
		//return existing connection
		return sharedConfig, nil
	}

	//load config
	config, err := loadConfiguration()

	if err != nil {
		return nil, err
	}

	//create connection string
	conn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DbHost, config.DbPort, config.DbUser, config.DbPass, config.DbName)

	//connect to the db
	db, err := sql.Open("postgres", conn)

	//stash config for reuse across goroutines
	sharedConfig = db

	//return handle to db
	return db, err
}

func loadConfiguration() (dbConfig, error) {
	var config dbConfig

	configFile, err := os.Open("db-configuration.json")
	defer configFile.Close()

	if err != nil {
		return dbConfig{}, err
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config, nil
}
