package sql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	_ "github.com/lib/pq" // driver
)

// Shared database connection for services
var sharedPool *sql.DB = nil
var poolExpires time.Time

// Shared mutex for connection with db
var sharedPoolLock sync.Mutex = sync.Mutex{}

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
	sharedPoolLock.Lock()
	defer sharedPoolLock.Unlock()

	// check if we are already connected
	if sharedPool != nil {
		now := time.Now()
		// check for expiration
		if now.After(poolExpires) {
			// reset shared connection
			sharedPool.Close()
			sharedPool = nil
		} else {
			// return existing connection
			return sharedPool, nil
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
	db.SetMaxIdleConns(50)

	// stash config for reuse across goroutines
	sharedPool = db
	poolExpires = time.Now().Add(time.Minute * 30)

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
