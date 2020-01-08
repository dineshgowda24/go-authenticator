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

* Setting up GO Environment

```bash
wget https://dl.google.com/go/go1.13.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.13.5.linux-amd64.tar.gz
export GOROOT=/usr/local/go
mkdir -p ~/go
mkdir -p ~/go/src
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```
You can add these exports in `~/.profile` so that they get set for every session. Run `go version` to check if GO is installed.

* Installing dep

```bash
sudo apt-get install go-dep
```
Run `dep version` to check if dep is installed.

* Cloning the app

```bash
cd ~/go/src
git clone <go-authenticator-clone-path>
```

Make sure the project is cloned only inside `~/go/src` as this is the workspace for GO.

* Installing dependent packages

```bash
cd ~/go/src/go-authenticator
dep ensure
```
`dep ensure` installs all the dependent packages for the app.

## Running the application

```go
cd ~/go/src/go-authenticator
go run main.go
```
After this you should get an output like this

```bash
Successfully connected to datastore at! localhost:3306
<nil>
2020/01/08 10:53:41 Server started and is listening on port '8080'
```

## License
[MIT](https://choosealicense.com/licenses/mit/)
