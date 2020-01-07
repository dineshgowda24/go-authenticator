package dbutils

import (
	"fmt"
	"log"
	"os"

	"github.com/dineshgowda24/go-authenticator/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Gorm mysql dialect interface
	"github.com/joho/godotenv"
)

//ConnectDB function: Make database connection
func ConnectDB() *gorm.DB {

	//Load environmenatal variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env config file")
	}

	username := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	databaseName := os.Getenv("DATABASE_NAME")
	databaseHost := os.Getenv("DATABASE_HOST")
	databasePort := os.Getenv("DATABASE_PORT")

	//Define DB connection string
	dbConnString := username + ":" + password + "@(" + databaseHost + ":" + databasePort + ")/" + databaseName + "?charset=utf8&parseTime=True&loc=Local"
	//connect to db URI
	db, err := gorm.Open("mysql", dbConnString)
	if err != nil {
		fmt.Println("error connecting to datastore", err)
		panic(err)
	}
	// close db when not in use
	//	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Session{})
	db.Model(&models.Session{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	fmt.Println("Successfully connected to datastore at!", databaseHost+":"+databasePort)
	return db
}
