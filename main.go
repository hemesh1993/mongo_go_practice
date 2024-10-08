package main

import (
	"log"
	"net/http"
	"github.com/hemesh1993/mongo_go_restapi/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	if err := http.ListenAndServe("localhost:9000", r); err != nil {
		log.Fatal(err)
	}

}

func getSession() *mgo.Session {

	s, err := mgo.Dial("mongodb://localhost:27107")
	if err != nil {
		panic(err)
	}
	return s
}
