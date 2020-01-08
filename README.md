# GO Authenticator

Go Authenticator is http app written in GO. The app rigisters a user into the datastore and allows signing into to the app.
The app also does a session management upon login. The session are valid for 10 minutes. Upon login users are redirected to home page.

## Why GO?

## Prequits 

In order for Go Authenticator to work it needs
  * MYSQL - MYSQL server install and running.
  * GO - GO support packages.
  * dep - dep is depenedency management for GO.
  
## Packages Used

The Go Authenticator uses `dep` for package management. `dep` is similar to `npm`.
The Go Authenticator requires below packages

* [godotenv](https://www.github.com/joho/godotenv) - .env configuration management
* [mux](https://github.com/gorilla/mux) - Http mux for routing
* [gorm](https://github.com/jinzhu/gorm) - ORM for Databases
* [mysql](https://github.com/go-sql-driver/mysql) - Mysql client for ORM

## Configuration

All the database and ports related configuration are present in `.env` file. Make sure the `.env` file is updated before starting the app. Also the database needs to be created before starting the app.

## Database Migrations

Go just needs the database to be created and all the migrations are handled by `gorm` package without needing to run any scripts.

## Installing Dependencies

## 

After cloning the project. `cd project`. Then just run the below command.

```go
go run main.go
```

## License
[MIT](https://choosealicense.com/licenses/mit/)
