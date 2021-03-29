# Helia
Helia is going to be a harsh, massively multiplayer, single world space game with nonconsensual PVP everywhere.

# Restoring the database
Helia uses PostgreSQL 12 as its database. 

Use the `schema.sql` script to scaffold an empty database.
Use the `testdb.sql` script to scaffold a test database.

The database configuration is in `db-configuration.json`

# Starting the backend
Helia's backend is written in go (1.16).

To start the go backend, run `go run main.go` in the root of the project.

# Starting the frontend
To start the angular frontend, run `npm start` in the frontend directory (one level below the root of the project).

# Shutting down the server
A server shutdown can be initiated using the "Save and Shutdown" endpoint (see Useful Links). This will save key aspects of the current state of the simulation and shutdown the server.

# Useful links
* Register: http://localhost:4200/auth/signup
* Login: http://localhost:4200/auth/signin
* Save and Shutdown: localhost:8080/api/shutdown?key=shutdownToken

The shutdown token is in `listener-configuration.json`
