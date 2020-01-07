package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/dineshgowda24/go-authenticator/dbutils"
	"github.com/dineshgowda24/go-authenticator/models"

	"golang.org/x/crypto/bcrypt"
)

var db = dbutils.ConnectDB()

//CreateUser creates a user and stores in DB
//All the form data is parsed and stored in the database
//The password will hashed and then stored
func CreateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		//check if the session is alive
		t := template.Must(template.ParseFiles("views/signuporlogin.html")) // Parse template file.
		t.Execute(w, nil)
		return
	}

	user := &models.User{}
	if err := r.ParseForm(); err != nil {
		t := template.Must(template.ParseFiles("views/signuporlogin.html")) // Parse template file.
		w.WriteHeader(http.StatusInternalServerError)
		t.Execute(w, err)
	}

	//get the form fields
	user.Email = r.FormValue("user_email")
	user.PhoneNumber = r.FormValue("user_pnumber")
	user.FirstName = r.FormValue("user_fname")
	user.LastName = r.FormValue("user_lname")

	pass, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("user_password")), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	user.Password = string(pass)

	createdUser := db.Create(user)
	t := template.Must(template.ParseFiles("views/signuporlogin.html"))
	if createdUser.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		t.Execute(w, createdUser.Error)
	} else {
		t.Execute(w, "User Created Successfully.")
	}
}

//Login validates all the form parameters passed
//It checks if the user email is present in our database and also generates a hash of the password, then
//compares it with the stored hash in our database
func Login(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/signuporlogin.html")) // Parse template file.

	if r.Method == http.MethodGet {
		//check if any active session exists
		_, err := r.Cookie("session_x_value")
		if err == http.ErrNoCookie {
			t.Execute(w, nil)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		t.Execute(w, nil)
		return
	}

	userRecord := &models.User{}
	if err := db.Where("email = ?", r.FormValue("user_email")).First(userRecord).Error; err != nil {
		t.Execute(w, "Invalid Credentials")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(userRecord.Password), []byte(string(r.FormValue("user_pass"))))
	if err != nil {
		t.Execute(w, "Invalid Credentials")
		return
	}

	//if the login is successful set cookiee value
	sessionrecord := &models.Session{}
	sessionrecord.UserID = userRecord.ID
	sessionencvalue, _ := bcrypt.GenerateFromPassword([]byte(userRecord.Email+strconv.Itoa(int(time.Now().Unix()))), bcrypt.DefaultCost)
	sessionrecord.SessionXValue = string(sessionencvalue)
	sessionrecord.ExpiresAt = time.Now().Add(time.Second * 600)
	session := &http.Cookie{
		Name:     "session_x_value",
		Value:    string(sessionencvalue),
		HttpOnly: true,
		Expires:  sessionrecord.ExpiresAt,
		Path:     "/",
	}
	db.Create(sessionrecord)
	http.SetCookie(w, session)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

//Welcome redirects the user to home page after login
func Welcome(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/signuporlogin.html")) // Parse template file.
	c, err := r.Cookie("session_x_value")
	if err == http.ErrNoCookie {
		t.Execute(w, nil)
		return
	}
	//if the cookie exists check if its a valid one in database
	sessionrec := &models.Session{}
	if err := db.Where("session_x_value = ?", c.Value).First(sessionrec).Error; err != nil {
		t.Execute(w, nil)
		return
	}

	//if its an expired cookie
	if time.Now().After(sessionrec.ExpiresAt) {
		t.Execute(w, nil)
		return
	}
	user := &models.User{}
	db.First(user, sessionrec.UserID)
	t = template.Must(template.ParseFiles("views/welcome.html")) // Parse template file.
	t.Execute(w, user)                                           // merge.
}

//Logout clears all the session cookies
func Logout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/signuporlogin.html")) // Parse template file.
	c, err := r.Cookie("session_x_value")
	if err == http.ErrNoCookie {
		t.Execute(w, nil)
		return
	}

	//expire the session in db
	db.Model(&models.Session{}).Where("session_x_value = ?", c.Value).Update("expires_at", time.Now())
	//Delete the Cookie
	c.MaxAge = -1
	http.SetCookie(w, c)
	t.Execute(w, nil)
}
