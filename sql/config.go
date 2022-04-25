package sql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	_ "github.com/lib/pq" // driver
)

// Shared database connection for services
var sharedConfig *sql.DB = nil

// Shared mutex for connection with db
var sharedDbLock sync.Mutex = sync.Mutex{}

type dbConfig struct {
	DbName  string `json:"dbname"`
	DbHost  string `json:"dbhost"`
	DbUser  string `json:"dbuser"`
	DbPass  string `json:"dbpass"`
	DbPort  int    `json:"dbport"`
	SSLMode string `json:"sslmode"`
}

func connect() (*sql.DB, error) {
	// lock connect to handle excessive clients
	sharedDbLock.Lock()
	defer sharedDbLock.Unlock()

	// check if we are already connected
	if sharedConfig != nil {
		// make sure we aren't over on connections
		if sharedConfig.Stats().OpenConnections+1 >= sharedConfig.Stats().MaxOpenConnections {
			// reset shared connection
			sharedConfig.Close()
			sharedConfig = nil
		} else {
			// return existing connection
			return sharedConfig, nil
		}
	}

	// load config
	config, err := loadConfiguration()

	if err != nil {
		return nil, err
	}

	// create connection string
	conn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		config.DbHost, config.DbPort, config.DbUser, config.DbPass, config.DbName, config.SSLMode)

	// connect to the db
	db, err := sql.Open("postgres", conn)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(2)

	// stash config for reuse across goroutines
	sharedConfig = db

	// return handle to db
	return db, err
}

func loadConfiguration() (dbConfig, error) {
	var config dbConfig

	// try to load under main.go position
	configFile, err := os.Open("db-configuration.json")

	if err != nil {
		// try to load under a child position
		configFile, err = os.Open("../db-configuration.json")
	}

	if err != nil {
		return dbConfig{}, err
	}

	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config, nil
}
