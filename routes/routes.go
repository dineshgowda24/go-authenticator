//Package routes defines all the routes in our webapp
package routes

import (
	"go-authenticator/controllers"

	"github.com/gorilla/mux"
)

//Handlers is basically the list of all the availabe routes in our web
//It returns a pointer to mux.Router which can be send to http.Handle which takes a handle
func Handlers() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", controllers.Welcome).Methods("GET")
	r.HandleFunc("/accounts/signup", controllers.CreateUser).Methods("GET")
	r.HandleFunc("/accounts/login", controllers.Login).Methods("GET")
	r.HandleFunc("/accounts/signup", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/accounts/login", controllers.Login).Methods("POST")
	r.HandleFunc("/accounts/logout", controllers.Logout).Methods("GET")
	r.HandleFunc("/accounts/logout", controllers.Logout).Methods("POST")
	return r
}
