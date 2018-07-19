package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Login Route\n")
}

func main() {
	router := httprouter.New()
	router.POST("/login", Login)

	log.Fatal(http.ListenAndServe(":8080", router))
}
