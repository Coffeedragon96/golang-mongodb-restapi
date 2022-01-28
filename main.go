package main

import (
	"net/http"

	"github.com/Coffeedragon96/golang-mongodb-restapi/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	// Initialise new router and assign it to variable r
	r := httprouter.New()

	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe("localhost:9000", r)
}

// To get new session
func getSession() *mgo.Session {
	// Dial a mongoDB session on s
	s, err := mgo.Dial("mongodb://localhost:27017")
	// Handle the error
	if err != nil {
		panic(err)
	}
	// Return session to NewUserController
	return s
}
