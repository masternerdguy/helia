# Helia
Helia is going to be a harsh, massively multiplayer, single world space game with nonconsensual PVP everywhere.

# Docker
Helia supports local development using docker containers. There are three containers defined:

* `helia-db`: A postgres server that also contains flyway. This provides your local database during development.
* `helia-engine`: A golang environment which allows you to build and run the backend game engine server.
* `helia-frontend`: An angular environment that allows you to build and run the frontend web client.

To start these containers, simply run `docker compose up`. From there, you can connect to and use these containers as you see fit. Note that both your repository and your database files are mapped as volumes within these containers to easily propagate changes and persist the local database.

Be aware that docker on Windows has extremely poor filesystem performance when using volumes - a good workaround is to install docker within WSL instead of directly on Windows. Visual Studio Code provides excellent tooling for working with both docker containers and WSL.

The environment variables shared by these containers are defined in `.env`.

# Restoring the database
This repo contains a sample local development database within `db.sample.7z`. Note that, due to its size, Git LFS is required to retrieve this file. Once extracted, the database can be restored by running `restore-db.sh` within the database container.

The database configuration is in `.env`.

# Flyway
Flyway is used for db migrations, which are stored in `flyway/sql`. Migrations can be applied to the local database by running `flyway-migrate.sh` in the database container. `flyway-info.sh` can be used to get the migration status. Note that SQL files are ignored in this repo, so you will need to use force when attempting to add them using git add.

# Starting the backend (local development)
Helia's backend is written in go (1.18). Since it makes heavy use of goroutines, it should be run in an environment with at least 4 cores - more is better, and core count is far more important than clock speed in determining overall performance.

To start the go backend, run `go run main.go` in the engine container.

# Starting the frontend (local development)
To start the angular frontend, run `npm start` in the frontend container. Note that obfuscation is applied as part of the build, so it may take a few minutes to start and also to apply any changes. This can be temporarily disabled by commenting out the obfuscator in `custom-webpack.config.json`, but under no circumstances should such a change be checked into source control (nor should unobfuscated code ever be exposed to the internet). Obfuscation of the client is an important part of Helia's security!

# Deployment Process (alpha, frontend)
Helia's open alpha is currently hosted on an Azure storage accout as a static website. This should allow hosting for pennies on the dollar compared to a traditional app service. The easiest way to deploy is to

1. Run `npm run build:alpha` in the frontend directory.
2. Replace the files in the `$web` container of the `heliaalphafrontend` storage account with the new files under the `dist` folder - the easiest way is to just use Azure Storage Explorer which takes ~1 minute.

# Deployment Process (alpha, backend)
Deploying the backend is less trivial than the frontend. Currently, the only way to run go code in an Azure app service is to use a docker container hosted in some kind of docker registry. For privacy reasons, the `helia-backend-engine` app service is configured to pull from the `heliaalpharegistry` Azure Container Registry automatically. Whenever a new docker image is pushed, a deployment occurs. Be aware that this will result in a sudden restart, so it is critical to perform a clean shutdown using the `Save and Shutdown` endpoint, allow it to complete, and then manually stop the app service before the deployment. Otherwise, players will lose progress and players don't tend to like that.

Given the above, performing a backend deployment can be done by

0. Perform a clean shutdown, wait for it to complete, and then manually stop the app service.
1. `cd` into the frontend directory - npm is being used as a helper here, at least for now. This isn't ideal as you will see.
2. Run `npm run build-docker:alpha` to build the docker image. To reduce image size significantly (large images are unreliable on Azure) some interesting tricks are done before building the container.
3. Run `cd .. && cd frontend` to get into the restored frontend folder again - one of the interesting tricks done by the build script is deleting that folder and then letting git restore it after the image is built. As a result, you lose your path reference.
4. Run `npm run deploy-docker:alpha` to push the image to Azure.
5. Start the app service.

# Shutting down the server properly (aka not making players very angry)
A server shutdown can be initiated using the `Save and Shutdown` endpoint (see `Useful Links`). This will save key aspects of the current state of the simulation and shut down the server. It is critical to always perform a clean shutdown - whether before a backend deployment or otherwise. If a clean shutdown is not performed, players will lose progress.

Also note that the app service will try to restart automatically a few minutes after the shutdown because it detects an "unhealthy resource" - it is important to monitor the logs table for the `shutdown complete` message and then manually stop the app service. Note that Helia does take ~10 minutes to boot within this particular Azure app service, and it is completely safe to stop the app service during startup. However, if the simulation is allowed to fully start a clean shutdown must be performed again.

# Daily Downtime
Helia is intended to be cleanly shut down and restarted once a day. If this doesn't happen, some gameplay aspects (such as transient jumphole connections) will become stale and won't be regenerated. Automatic restarts are scheduled to occur at noon EDT and are handled by a core goroutine.

# Single Process, Single Server
Helia was designed from the beginning to run as a single process - it will break in lots of fascinating ways if any kind of horizontal scaling is used. Only vertical scaling can be used with Helia! This is very intentional, as ultimately Helia is planned to run on a self-hosted, ARM-based, server where go can take full advantage of an enormous core count.

Helia is designed for all players to play within the same world and server - any kind of "realms" or similar instancing are philosophically incompatible with Helia. Obviously test environments are fine, but production should always be a single world and a single server.

# Useful links
* Register: http://localhost:4200/#/auth/signup
* Login: http://localhost:4200/#/auth/signin
* Save and Shutdown: http://localhost:8000/api/shutdown?key=shutdownToken

The shutdown token is in `.env`. Note that ports and protocols (http vs https) will differ between environments.

# Devhax
Various cheats are provided that can be used to more easily test helia or just generally show off to plebs. These are executed using the system chat window. Only a user who's record in the users table has `isdev` set to `true` can run these commands. Note that these are HACKS, and things may not respond totally as expected in all cases!

## teleport
syntax:  /devhax teleport [systemname]
example: /devhax teleport C-5

Transports you to the system who's name is provided. You must be undocked to use this command. This may lead to you being in an unexpected place upon a system restart.

## cargo
syntax:  /devhax cargo [quantity]~~[itemtypename]
example: /devhax cargo 200~~Basic Armor Plate

Creates a new stack of a given size of a given item type in your cargo bay. Completely ignores volume constraints!

## remax
syntax:  /devhax remax all
example: /devhax remax all

Restores your ship's shield, fuel, etc, to their maximum values.

## yankall
syntax:  /devhax yankall [bots|humans|users]
example: /devhax yankall users

Pulls *ALL* undocked ships of a given controller type to your current system. Bots will yank only NPCs, humans will yank only human players, and users will yank everyone. This may lead to users being in unexpected places upon system restart, and is also very rude!

## yankfact
syntax:  /devhax yankfact [ticker]
example: /devhax yankfact TVC

Pulls *ALL* undocked ships of a given faction to your current system. The faction is specified using the short ticker.
This may lead to users being in unexpected places upon system restart, and is also very rude!

## wallet
syntax:  /devhax wallet [integer]
example: /devhax wallet 9000

Sets your current ship's wallet to the provided value.

## ship
syntax:  /devhax ship [itemtypename]
example: /devhax ship Robin

Creates a new ship of a given type at the station you are currently docked at. Ignores limits on maximum parked ships at the same station!

## tonpc
syntax:  /devhax tonpc [behaviourtype]
example: /devhax tonpc 3

Sets your current ship up as an "NPC" following a given behaviour mode number.

* None:         0
* Wander:       1
* Patrol:       2
* PatchTrade:   3
* PatchMine:    4
* PatchSalvage: 5

Sending a negative number will undo this effect.
