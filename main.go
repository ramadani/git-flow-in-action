package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "User Login Route\n")
}

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Register Route\n")
}

func main() {
	router := httprouter.New()
	router.POST("/login", Login)
	router.POST("/register", Register)

	log.Fatal(http.ListenAndServe(":8080", router))
}
