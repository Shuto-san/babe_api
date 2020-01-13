# babe_api
## docker container run
```
docker-compose up -d
```
## go run
```
go run main.go
```

## directory structure
### commands
 Batch files
### configs
 Common variables setting
### controllers
 Contains handler functions for particular route to be called when an api is called
### db
 Contains migration files and db settings
### middlewares
 Middleware to be used for this project
### models
 Database tables to be used as model struct
### routes
 Url routing when a request comes
### services
 Contains all the core business logic including db connections
### tests
 Testing files
### .env
 Environment variables setting
### main.go
 Starting point
