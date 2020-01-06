package main

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Gorm postgres dialect interface
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
	"github.com/dgrijalva/jwt-go"
)
