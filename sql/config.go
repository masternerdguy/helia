package sql

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	_ "github.com/lib/pq" // driver
)

// Shared database connection for services
var sharedPool *sql.DB = nil
var poolExpires time.Time

// Shared mutex for connection with db
var sharedPoolLock sync.Mutex = sync.Mutex{}

// Structure representing the configuration needed to connect to the database
type dbConfig struct {
	DbName  string
	DbHost  string
	DbUser  string
	DbPass  string
	DbPort  int
	SSLMode string
}

// Establishes a connection to the database or reuses an existing one
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

// Reads the application database configuration from environment variables
func loadConfiguration() (dbConfig, error) {
	// read environment variables
	dbHost := os.Getenv("dbhost")
	dbUser := os.Getenv("dbuser")
	dbPass := os.Getenv("dbpass")
	dbPort := os.Getenv("dbport")
	dbName := os.Getenv("dbname")
	sslMode := os.Getenv("sslmode")

	// parse port id
	dbPortInt, err := strconv.ParseInt(dbPort, 10, 32)

	// return configuration
	return dbConfig{
		DbHost:  dbHost,
		DbUser:  dbUser,
		DbPass:  dbPass,
		DbPort:  int(dbPortInt),
		DbName:  dbName,
		SSLMode: sslMode,
	}, err
}
