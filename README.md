# Helia
Helia is going to be a harsh, massively multiplayer, single world space game with nonconsensual PVP everywhere.

# Restoring the database
Helia ideally would use PostgreSQL 14 as its database, but note that Azure is quite out of date at 11. However, things are close enough that development between 11 and 14 shouldn't be an issue since very basic postgres features are being used. As of the open alpha, a backup of the `helia` database within the `helia-alpha` Azure Postgres Server should be grabbed for local development. There is no database migration strategy yet, but there will be.

The database configuration is in `db-configuration.json`.

# Starting the backend (local development)
Helia's backend is written in go (1.18). Since it makes heavy use of goroutines, it should be run in an environment with at least 4 cores - more is better, and core count is far more important than clock speed in determining overall performance.

To start the go backend, run `go run main.go` in the root of the project.

# Starting the frontend (local development)
To start the angular frontend, run `npm start` in the frontend directory (one level below the root of the project). Note that obfuscation is applied as part of the build, so it may take a few minutes to start and also to apply any changes. This can be temporarily disabled by commenting out the obfuscator in `custom-webpack.config.json`, but under no circumstances should such a change be checked into source control (nor should unobfuscated code ever be exposed to the internet). Obfuscation of the client is an important part of helia's security!

# Deployment Process (alpha, frontend)
Helia's open alpha is currently hosted on an Azure storage accout as a static website. This should allow hosting for pennies on the dollar compared to a traditional app service. The easiest way to deploy is to

1. Run `npm run build:alpha` in the frontend directory.
2. Replace the files in the `$web` container of the `heliaalphafrontend` storage account with the new files under the `dist` folder - the easiest way is to just use Azure Storage Explorer which takes ~1 minute.

# Deployment Process (alpha, backend)
Deploying the backend is less trivial than the frontend. Currently, the only way to run go code in an Azure app service is to use a docker container hosted in some kind of docker registry. For privacy reasons, the `helia-backend-engine` app service is configured to pull from the `heliaalpharegistry` Azure Container Registry automatically. Whenever a new docker image is pushed, a deployment occurs. Be aware that this will result in a sudden restart, so it is critical perform a clean shutdown using the `Save and Shutdown` endpoint, allow it to complete, and then stop the app service before the deployment. Otherwise, players will lose progress and players don't tend to like that.

# Shutting down the server properly (aka not making players very angry)
A server shutdown can be initiated using the `Save and Shutdown` endpoint (see Useful Links). This will save key aspects of the current state of the simulation and shut down the server. It is critical to always perform a clean shutdown - whether before a backend deployment or otherwise. If a clean shutdown is not performed, players will lose progress.

Also note that the app service will try to restart automatically a few minutes after the shutdown because it detects an "unhealthy resource" - it is important to monitor the logs table for the `shutdown complete` message and then manually stop the app service. Note that helia does take ~10 minutes to boot within this particular Azure app service, and it is completely safe to stop the app service during startup. However, if the simulation is allowed to fully start a clean shutdown must be performed again.

# Useful links
* Register: http://localhost:4200/#/auth/signup
* Login: http://localhost:4200/#/auth/signin
* Save and Shutdown: http://localhost:8000/api/shutdown?key=shutdownToken

The shutdown token is in `listener-configuration.json`. Note that ports and protocols (http vs https) will differ between environments.
