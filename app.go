package main

import (
	"github.com/jamesmura/gostack/web"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", web.Index)
	router.GET("/register", web.GetRegister)
	router.POST("/register", web.PostRegister)
	router.ServeFiles("/static/*filepath", http.Dir("public/"))
	log.Fatal(http.ListenAndServe(":8080", router))
}
